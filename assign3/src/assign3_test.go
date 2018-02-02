package assign3

import "testing"

var adjlist = make([][]int, 4)

func TestGetMinCut(t *testing.T) {
	adjlist[0] = []int{1, 2}
	adjlist[1] = []int{0, 2, 3}
	adjlist[2] = []int{0, 1, 3}
	adjlist[3] = []int{1, 2}
	v := GetMinCut(adjlist)
	if v != 2 {
		t.Error("Expected 2, got ", v)
	}
}
