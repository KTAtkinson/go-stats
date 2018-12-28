package collector

import (
	"fmt"
	"sync"
)

type Stat struct {
	millis int64
	value  int64
}

type FlusherIface interface {
	Flush(StatsRecorder) error
}

type Stats struct {
	lock *sync.Mutex
	stats []*Stat
}

func (s *Stats) Record(stats []*Stat) error {
	return NOT_IMPLEMENTED
}

func (s *Stats) Reset() ([]*Stat, error) {
	return nil, NOT_IMPLEMENTED
}

type StatsRecorder map[string]*Stats

type Collector struct {
	// The channel which holds stats waiting to be collected.
	stats StatsRecorder
	// Flusher interface to flush stats
	flusher FlusherIface
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
func (s *Stats) FlushAlways(interval int) {
	fmt.Print("FlushAlways not implimented. Not flushing. Exiting...")
}

type DiskFlusher struct {
	flushRoot string
}

// Flushes values to disk
func (d *DiskFlusher) Flush(stats StatsRecorder) error {
	return NOT_IMPLEMENTED
}
