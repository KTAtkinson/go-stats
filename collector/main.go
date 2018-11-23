package collector

import (
	"fmt"
)

type statType int

// An incrementing counter
var Counter statType = 0

// A statitic with distrete data points
var Point statType = 1

type Stat struct {
	type_  statType
	tag    string
	millis int64
	value  int64
}

type Stats struct {
	// The channel which holds stats waiting to be collected.
	collector *map[string]int64
	// The location on disk holding the stats
	statsRoot string
}

// Increments the counter with the given tag.
func (s *Stats) Count(counts []*Stat) error {
	return NOT_IMPLEMENTED
}

// Records a data point with the given value with the given tag.
func (s *Stats) Record([]*Stat) error {
	return NOT_IMPLEMENTED
}

// Collects and combines stats.
func (s *Stats) Collect() error {
	return NOT_IMPLEMENTED
}

// Collects stats indefinately at the given interval seconds.
func (s *Stats) CollectorAlways(interval int64) {
	fmt.Print("Not implemented")
}

// Flushes stats from memory to disk.
func (s *Stats) Flush() error {
	return NOT_IMPLEMENTED
}

// Flushes stats indefinately at interval
func (s *Stats) FlushAlways() {
	fmt.Print("Not implemented")
}

type StatsManager struct {
	*Stats
}

func NewStatsManager(statsRoot string) *StatsManager {
	return &StatsManager{
		Stats: &Stats{
			collector: new(map[string]int64),
			statsRoot: statsRoot,
		},
	}
}

type StatsWorker struct {
	*Stats
	managerAddr string
}

func NewStatsWorker(statsRoot, managerAddr string) *StatsWorker {
	return &StatsWorker{
		Stats: &Stats{
			statsRoot: statsRoot,
			collector: new(map[string]int64),
		},
		managerAddr: managerAddr,
	}
}

// Flushes stats saved on local disk to stats manager.
func (s *StatsWorker) FlushiToRemote(remoteStats string) error {
	return NOT_IMPLEMENTED
}

// Flushes stats indefinately at the given interval.
func (s *StatsWorker) FlushToRemoteAlways(interval int64) {
	fmt.Print("Not implemented")
}
