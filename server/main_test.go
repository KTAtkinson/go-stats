package main

import (
	"flag"
	"testing"
)

func TestMain(t *testing.T) {
	var exitCode int
	exitFn = func(code int) { exitCode = code }
	cases := []struct {
		port      string
		flushAddr string
		exitCode  int
	}{
		{"6000", "https://www.stats.server", 2},
		{"6000", "/stats/", 0},
	}

	for _, c := range cases {
		flag.Set("flush-to-addr", c.flushAddr)
		flag.Set("port", c.port)
		main()
		if exitCode != c.exitCode {
			t.Errorf("Expected error code %d when flush address is %s, found %d.", c.exitCode, c.flushAddr, exitCode)
		}
		exitCode = 0
	}
}
