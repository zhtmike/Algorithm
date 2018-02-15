package main

import (
	"container/heap"
)

// MedianMaintenance computes the all medians from index 0 to index i.
// Where i is from 1 to len(arr)
func MedianMaintenance(arr []int) []int {
	// Declare two heaps for storing data
	var minhp IntHeap
	var maxhp IntHeap
	maxhp.ismax = true

	medians := make([]int, len(arr))

	// Init heaps
	heap.Init(&minhp)
	heap.Init(&maxhp)

	for i, val := range arr {
		if maxhp.hp != nil {
			if val > maxhp.hp[0] {
				heap.Push(&minhp, val)
			} else {
				heap.Push(&maxhp, val)
			}
		} else {
			heap.Push(&maxhp, val)
		}

		if len(minhp.hp)-len(maxhp.hp) > 0 {
			heap.Push(&maxhp, heap.Pop(&minhp))
		} else if len(maxhp.hp)-len(minhp.hp) > 1 {
			heap.Push(&minhp, heap.Pop(&maxhp))
		}

		medians[i] = maxhp.hp[0]
	}
	return medians
}
