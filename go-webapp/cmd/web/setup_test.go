package main

import (
	"os"
	"testing"
)

var app application

// TestMain is the entry point for all tests.
func TestMain(m *testing.M) {

	os.Exit(m.Run())
}
