package main

import (
	"flag"
	"testing"
)

func TestRun(t *testing.T) {
	cases := []struct{
		port string
		statsPath string
		isManager string
		managerAddr string
		expectError bool
	}{
		{"8080", "", "false", "managerAddr", false},
		{"8080", "", "false", "", true},
	}

	for _, case_ := range cases {
		flag.Set("port", case_.port)
		flag.Set("statsPath", case_.statsPath)

		if case_.isManager != "" {
			flag.Set("isManager", case_.isManager)
		}

		if case_.managerAddr != "" {
			flag.Set("managerAddr", case_.managerAddr)
		}

		run(func(int){})
	}
}
