package main

import (
	"flag"
	"fmt"
	stats "github.com/KTAtkinson/go-stats/collector"
	"net/url"
	"os"
)

var exitFn func(int)

func init() {
	exitFn = os.Exit
}

func main() {
	flag.Parse()

	url_, err := url.Parse(flushToAddr)
	if err != nil {
		fmt.Printf("Failed to parse flush address %s due to error: %s\n", flushToAddr, err)
		exitFn(1)
	}

	var collector *stats.Collector
	if url_.IsAbs() {
		fmt.Println("Flushing to a remote stats service not implemented.")
		exitFn(2)
	} else {
		collector = stats.NewOnDiskCollector(url_.String())
	}

	Start(port, collector)
}
