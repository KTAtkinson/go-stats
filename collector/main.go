package collector

import (
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"sort"
	"sync"
	"time"
)

type Stat struct {
	time  time.Time
	value int64
}

type FlusherIface interface {
	Flush(StatsRecorder) error
}

type Stats struct {
	lock  *sync.Mutex
	stats []*Stat
}

func (s *Stats) Record(stats []*Stat) error {
	return NOT_IMPLEMENTED
}

func (s *Stats) Reset() ([]*Stat, error) {
	s.lock.Lock()

	value := s.stats
	s.stats = make([]*Stat, 0)

	s.lock.Unlock()
	return value, nil
}

type StatsRecorder map[string]*Stats

type Collector struct {
	// The channel which holds stats waiting to be collected.
	stats StatsRecorder
	// Flusher interface to flush stats
	flusher FlusherIface
}

func NewCollector(flusher FlusherIface) *Collector {
	var recorder StatsRecorder
	return &Collector{
		stats:   recorder,
		flusher: flusher,
	}
}

// Records a data point with the given value with the given tag.
func (s *Collector) Record(name string, stats []*Stat) error {
	return NOT_IMPLEMENTED
}

// Flushes stats from memory to disk.
func (s *Collector) Flush() error {
	return s.flusher.Flush(s.stats)
}

// Flushes stats indefinately at interval
func (s *Collector) FlushAlways(interval time.Duration, errs chan<- error) {
	for range time.Tick(interval) {
		errs <- s.Flush()
	}
}

type OnDiskFlusher struct {
	flusherRoot string
}

// Flushes values to disk
func (d *OnDiskFlusher) Flush(stats StatsRecorder) (err error) {
	for name, stat := range stats {
		entries, err := stat.Reset()
		if err != nil {
			log.Printf("Error while resetting entry for %s.\n", name)
			continue
		}
		sort.Slice(entries, func(i int, i1 int) bool { return entries[i].time.Before(entries[i1].time) })

		var roundedTime time.Time
		var writer io.WriteCloser
		for _, evt := range entries {
			roundedEventTime := evt.time.UTC().Round(time.Hour * 24)
			if roundedEventTime != roundedTime {
				if writer != nil {
					writer.Close()
				}
				fullPath := filepath.Join(d.flusherRoot, name, fmt.Sprintf("%d", roundedEventTime.Unix()))
				fileMode := os.O_WRONLY | os.O_APPEND | os.O_CREATE
				writer, err = os.OpenFile(fullPath, fileMode, 0644)
				// This may mean the directory doesn't exist, so create it and see if the error goes away.
				if os.IsNotExist(err) {
					err = os.MkdirAll(filepath.Dir(fullPath), 0744)
					if err != nil {
						log.Printf("Failed to created new directory to record stats for %s.\n", name)
						continue
					}
					writer, err = os.OpenFile(fullPath, fileMode, 0644)
				}
				if err != nil {
					log.Printf("Not able to open file %s to record stats for %s. %s", fullPath, name, err)
					continue
				}
			}
			writer.Write([]byte(fmt.Sprintf("%d %d", evt.time.Unix(), evt.value)))
		}
	}
	return err
}

func NewOnDiskCollector(fp string) *Collector {
	flusher := OnDiskFlusher{
		flusherRoot: fp,
	}
	return NewCollector(&flusher)
}
