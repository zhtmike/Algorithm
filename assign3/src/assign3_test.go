package assign3

import "testing"

var adjlist = make([][]int, 4)

func TestGetMinCut(t *testing.T) {
	adjlist[0] = []int{2, 3}
	adjlist[1] = []int{1, 3, 4}
	adjlist[2] = []int{1, 2, 4}
	adjlist[3] = []int{2, 3}
	v := GetMinCut(adjlist)
	if v != 2 {
		t.Error("Expected 2, got ", v)
	}
}
