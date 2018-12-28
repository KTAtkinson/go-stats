package main

import (
	"net/http"
)

// Record data points in collector.
func CollectPoint(http.ResponseWriter, *http.Request) error {
	return NOT_IMPLEMENTED
}

// Increment a counter.
func CollectCount(http.ResponseWriter, *http.Request) error {
	return NOT_IMPLEMENTED
}

// List recorded points for a given tag.
func GetPoints(http.ResponseWriter, *http.Request) error {
	return NOT_IMPLEMENTED
}

// Retrieve current count for given tag.
func GetCounter(http.ResponseWriter, *http.Request) error {
	return NOT_IMPLEMENTED
}

// Returns health of the stats server.
func healthz(rsp http.ResponseWriter, req *http.Request) {
	rsp.WriteHeader(200)
}
