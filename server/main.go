package main

import (
	"errors"
	"flag"
	"log"
	"os"

	"github.com/KTAtkinson/go-stats/collector"
)

type exitFn func(int)

func getCollector(managerAddr string, statsPath string, isManager bool) (CollectorIface, error) {
	var collectorIface CollectorIface
	if managerAddr == "" && !isManager {
		return collectorIface, errors.New("Either the manager address should be provided or the stats server should be the manager.")
	}

	if isManager {
		collectorIface = collector.NewStatsManager(
			statsPath,
		)
	} else {
		collectorIface = collector.NewStatsWorker(
			statsPath,
			managerAddr,
		)
	}

	return collectorIface, nil
}

func run(exit exitFn) {
	port := flag.String("port", "5000", "Port on which the server should run")
	statsPath := flag.String("stats-path", "/var/stats/", "Location on disk where stats will be written.")
	isManager := flag.Bool("is-manager", true, "If the server being created is the manager stats server.")
	managerAddr := flag.String("manager-address", "", "The address to the manager server.")
	flag.Parse()
	statsCollector, err := getCollector(*managerAddr, *statsPath, *isManager)
	if err != nil {
		log.Printf("Error while making stats collector. %s", err)
		exit(3)
	}
	if err := Start(*port, statsCollector); err != nil {
		log.Printf("Server failed. %s", err)
		exit(2)
	}
}

func main() {
	run(os.Exit)
}
