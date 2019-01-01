package collector

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"sync"
	"testing"
	"time"
)

func TestFlushToDisk(t *testing.T) {
	saveOutput := false
	tmpDir, err := ioutil.TempDir("", "test_events")
	if err != nil {
		t.Fatalf("Failed to create tempfile. %s", err)
	}
	defer func() {
		if saveOutput {
			log.Printf("Test output was not deleted due to errors. It can be found at %s.\n", tmpDir)
		} else {
			os.RemoveAll(tmpDir)
		}
	}()

	recorder := make(StatsRecorder)
	stat0Name := "stat0"
	stat0Fn := filepath.Join(tmpDir, stat0Name, fmt.Sprintf("%d", time.Now().UTC().Round(time.Hour*24).Unix()))
	if err = os.MkdirAll(filepath.Dir(stat0Fn), 0744); err != nil {
		t.Fatalf("Failed to create fake stat directory. %s", err)
	}
	statFile, err := os.OpenFile(stat0Fn, os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		t.Fatalf("Open example stats file for writing. %s", err)
	}
	defer statFile.Close()

	_, err = statFile.Write(testData0)
	if err != nil {
		t.Fatalf("Failed to write example stats. %s", err)
	}
	newStats := []*Stat{
		&Stat{
			time:  time.Now(),
			value: 4,
		},
	}
	recorder[stat0Name] = &Stats{
		lock:  &sync.Mutex{},
		stats: newStats,
	}

	stat1Name := "stat1"
	newStats = []*Stat{
		&Stat{
			time:  time.Now(),
			value: 0,
		},
	}
	recorder[stat1Name] = &Stats{
		lock:  &sync.Mutex{},
		stats: newStats,
	}

	flusher := &OnDiskFlusher{
		flusherRoot: tmpDir,
	}
	err = flusher.Flush(recorder)
	if err != nil {
		t.Fatalf("Failed to flush new stats. %s", err)
	}

	for stat, _ := range recorder {
		dirExists, err := hasDir(tmpDir, stat)
		if err != nil {
			t.Fatalf("Failed to list temperary directory %s. Error, %s", tmpDir, err)
		}
		if !dirExists {
			t.Errorf("Failed to find directory %s in temperary directory %s.", stat, tmpDir)
		}
	}
	exampleStat, err := os.Open(stat0Fn)
	if err != nil {
		t.Errorf("Failed to open stats file for reading. %s", err)
	}
	defer exampleStat.Close()

	readBytes, err := ioutil.ReadFile(stat0Fn)
	if err != nil {
		t.Errorf("Failed to read from stats file. %s", err)
	} else {
		maybeNewBytes := bytes.TrimPrefix(readBytes, testData0)
		log.Printf("New bytes: %s", string(maybeNewBytes))
	}
}

func hasDir(dir, contains string) (bool, error) {
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		return false, err
	}
	for _, f := range files {
		if f.Name() == contains {
			return true, nil
		}
	}

	return false, nil
}
