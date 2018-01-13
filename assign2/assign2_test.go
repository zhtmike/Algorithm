package assign2

import (
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
