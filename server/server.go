package main

import (
	"fmt"
	"net/http"

	"github.com/KTAtkinson/go-stats/collector"
)

type CollectorIface interface {
	Record(string, []*collector.Stat) error
	FlushAlways(int)
}

type StatsServer struct {
	Collector CollectorIface
	Server    *http.Server
}

// Create a server to record and report metrics. The following endoints are included:
// * /collect/points
// * /collect/counts
// + /list/points/{name}
// * /list/counter/{name}
func Start(port int, collector CollectorIface) {
	fmt.Println("No server to run")
}

// Record data points in collector.
func (s *StatsServer) CollectPoint(http.ResponseWriter, *http.Request) error {
	return NOT_IMPLEMENTED
}

// Increment a counter.
func (s *StatsServer) CollectCount(http.ResponseWriter, *http.Request) error {
	return NOT_IMPLEMENTED
}

// List recorded points for a given tag.
func (s *StatsServer) GetPoints(http.ResponseWriter, *http.Request) error {
	return NOT_IMPLEMENTED
}

// Retrieve current count for given tag.
func (s *StatsServer) GetCounter(http.ResponseWriter, *http.Request) error {
	return NOT_IMPLEMENTED
}
