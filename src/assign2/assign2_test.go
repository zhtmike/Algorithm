package main

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
	num := 0
	QuickSort(seq, &num, ChooseRandomPivot)
	sort.Ints(seqBak)
	if !reflect.DeepEqual(seq, seqBak) {
		t.Error("Expected ", seqBak, ", got ", seq)
	}
}

func TestQuickSortLong(t *testing.T) {
	longSeq := make([]int, 50)
	longSeqBak := make([]int, 50)
	num := 0
	for i := range longSeq {
		longSeq[i] = rand.Intn(100)
	}
	copy(longSeqBak, longSeq)
	QuickSort(longSeq, &num, ChooseRandomPivot)
	sort.Ints(longSeqBak)
	if !reflect.DeepEqual(longSeq, longSeqBak) {
		t.Error("Expected ", longSeqBak, ", got ", longSeq)
	}
}

func TestQuickSortLongFirst(t *testing.T) {
	longSeq := make([]int, 50)
	longSeqBak := make([]int, 50)
	num := 0
	for i := range longSeq {
		longSeq[i] = rand.Intn(100)
	}
	copy(longSeqBak, longSeq)
	QuickSort(longSeq, &num, ChooseFirstPivot)
	sort.Ints(longSeqBak)
	if !reflect.DeepEqual(longSeq, longSeqBak) {
		t.Error("Expected ", longSeqBak, ", got ", longSeq)
	}
}

func TestQuickSortLongLast(t *testing.T) {
	longSeq := make([]int, 50)
	longSeqBak := make([]int, 50)
	num := 0
	for i := range longSeq {
		longSeq[i] = rand.Intn(100)
	}
	copy(longSeqBak, longSeq)
	QuickSort(longSeq, &num, ChooseLastPivot)
	sort.Ints(longSeqBak)
	if !reflect.DeepEqual(longSeq, longSeqBak) {
		t.Error("Expected ", longSeqBak, ", got ", longSeq)
	}
}

func TestQuickSortLongMedian(t *testing.T) {
	longSeq := make([]int, 50)
	longSeqBak := make([]int, 50)
	num := 0
	for i := range longSeq {
		longSeq[i] = rand.Intn(100)
	}
	copy(longSeqBak, longSeq)
	QuickSort(longSeq, &num, ChooseMedianPivot)
	sort.Ints(longSeqBak)
	if !reflect.DeepEqual(longSeq, longSeqBak) {
		t.Error("Expected ", longSeqBak, ", got ", longSeq)
	}
}
