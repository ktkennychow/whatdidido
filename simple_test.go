package main

import (
	"testing"
)

func TestDummy(t *testing.T) {
	// Dummy test to check if testing framework works
	if 1+1 != 2 {
		t.Error("Math is broken")
	}
}
