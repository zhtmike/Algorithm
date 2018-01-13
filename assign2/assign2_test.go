package assign2

import (
	"math/rand"
	"reflect"
	"sort"
	"testing"
)

var seq = []int{20, 0, 2, 5, -2, 999, 22, 6, 3}

func TestQuickSort(t *testing.T) {
	seqBak := make([]int, len(seq))
	copy(seqBak, seq)
	QuickSort(seq)
	sort.Ints(seqBak)
	if !reflect.DeepEqual(seq, seqBak) {
		t.Error("Expected ", seqBak, ", got ", seq)
	}
}

func TestQuickSortLong(t *testing.T) {
	longSeq := make([]int, 50)
	longSeqBak := make([]int, 50)
	for i := range longSeq {
		longSeq[i] = rand.Intn(100)
	}
	copy(longSeqBak, longSeq)
	QuickSort(longSeq)
	sort.Ints(longSeqBak)
	if !reflect.DeepEqual(longSeq, longSeqBak) {
		t.Error("Expected ", longSeqBak, ", got ", longSeq)
	}
}
