package main

import (
	"reflect"
	"testing"
)

var adjlist = make([][]int, 9)

func TestGetMinCut(t *testing.T) {
	adjlist[0] = []int{6}
	adjlist[1] = []int{4}
	adjlist[2] = []int{8}
	adjlist[3] = []int{0}
	adjlist[4] = []int{7}
	adjlist[5] = []int{2, 7}
	adjlist[6] = []int{3, 8}
	adjlist[7] = []int{1}
	adjlist[8] = []int{5}
	v := KosarajuAlgo(adjlist)
	ans := []int{3, 3, 3}
	if !reflect.DeepEqual(v, ans) {
		t.Error("Expected ", ans, "got ", v)
	}
}
