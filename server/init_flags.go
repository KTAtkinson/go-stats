package main

import (
	"flag"
)

var flushToAddr string
var flushIntervalSecs int
var port int
var healthzPort int

func init() {
	flag.StringVar(&flushToAddr, "flush-to-addr", "/stats/", "A URL where to flush statistics from memory.")
	flag.IntVar(&flushIntervalSecs, "flush-interval-seconds", 60, "The interval in seconds in which stats will be sent to the flush address.")
	flag.IntVar(&port, "port", 1119, "The port on which the server shoulr run.")
	flag.IntVar(&healthzPort, "healthz-port", 1120, "The port on which the healthcheck server runs.")
}
