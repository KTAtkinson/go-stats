package main

import (
  "net/http"

//  "github.com/KTAtkinson/go-stats/collector"
)

type CollectorIface interface {
   Count(string) error
   Record( string, int64) error
}

type StatsServer struct {
    Collector CollectorIface
    Server *http.Server
}

// Create a server to record and report metrics. The following endoints are included:
// * /collect/points
// * /collect/counts
// + /list/points/{name}
// * /list/counter/{name}
func Start(port string, collector CollectorIface) error {
    return NOT_IMPLEMENTED
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
