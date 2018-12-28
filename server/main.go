package main

import (
	"flag"
	"fmt"
	stats "github.com/KTAtkinson/go-stats/collector"
	"log"
	"net/http"
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

	errs := make(chan error)
	go collector.FlushAlways(flushDuration, errs)
	go func() {
		log.Println("Recording errors while flushing")
		for err := range errs {
			if err != nil {
				log.Printf("Error while flushing stats. %s\n", err)
			} else {
				log.Printf("Flushing complete")
			}
		}
	}()

	apiServer := http.NewServeMux()
	go func() {
		log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), apiServer))
	}()
	log.Println("Started api server.")

	healthzServer := http.NewServeMux()
	healthzServer.HandleFunc("/healthz", healthz)
	log.Printf("Reporting health at 127.0.0.1:%d.\n", healthzPort)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", healthzPort), healthzServer))
}
