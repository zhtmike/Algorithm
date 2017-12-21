package main

import (
	"testing"
)

func TestBruteForceCount(t *testing.T) {
	seq := []int{3, 2, 5, 6, 20}
	v := BruteForceCount(seq)
	if v != 1 {
		t.Error("Expected 1, got ", v)
	}
}

func TestInverseCount(t *testing.T) {
	seq := []int{3, 2, 5, 6, 20}
	v := InverseCount(seq)
	if v != 1 {
		t.Error("Expected 1, got ", v)
	}
}
