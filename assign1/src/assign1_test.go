package assign1

import (
	"math/rand"
	"testing"
)

var seq = []int{3, 2, 5, 6, 20}

func TestBruteForceCount(t *testing.T) {
	v := BruteForceCount(seq)
	if v != 1 {
		t.Error("Expected 1, got ", v)
	}
}

func TestInverseCount(t *testing.T) {
	v := InverseCount(seq)
	if v != 1 {
		t.Error("Expected 1, got ", v)
	}
}

func TestInverseCountEqual(t *testing.T) {
	longSeq := make([]int, 50)
	longSeqBak := make([]int, 50)
	for i := range longSeq {
		longSeq[i] = rand.Intn(100)
	}
	copy(longSeqBak, longSeq)
	v1 := InverseCount(longSeq)
	v2 := BruteForceCount(longSeqBak)
	if v1 != v2 {
		t.Error("Expected true, got ", v1, v2)
	}
}
