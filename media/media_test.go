package main

import (
	"testing"
)

func TestScanDirectory(t *testing.T) {
	items := ScanDir("tests/media_repos")

	if len(items) < 2 {
		t.Fatalf("ScanDir should discover at least two media items")
	}
}
