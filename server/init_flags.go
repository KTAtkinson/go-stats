package main

import (
	"flag"
	"time"
)

var flushToAddr string
var flushDuration time.Duration
var port int
var healthzPort int

func init() {
	flag.StringVar(&flushToAddr, "flush-to-addr", "/stats/", "A URL where to flush statistics from memory.")
	flag.DurationVar(&flushDuration, "flush-duration", time.Second * 60, "The duration at which to flush stats to storage.")
	flag.IntVar(&port, "port", 1119, "The port on which the server shoulr run.")
	flag.IntVar(&healthzPort, "healthz-port", 1120, "The port on which the healthcheck server runs.")
}
