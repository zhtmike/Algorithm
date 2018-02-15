package main

import (
	"container/heap"
	"fmt"
	"sync"
)

// TwoSum compute the number of elements which their sum equals to x, where -10000 <= x <= 10000
func TwoSum(arr []int, threads int) int {
	const start = -10000
	const end = 10000

	var wg sync.WaitGroup
	wg.Add(threads)
	mutex := &sync.Mutex{}

	// convert arr to hash table
	hashTable := make(map[int]int)
	for i, val := range arr {
		hashTable[val] = i
	}

	sum := 0
	n := (end - start) / float32(threads)
	for k := 0; k < threads; k++ {
		go func(k int) {
			defer wg.Done()
			for i := int(n*float32(k)) + start; i <= int(n*float32(k+1))+start; i++ {
				for _, num := range arr {
					j, ok := hashTable[i-num]
					if ok && j != i {
						mutex.Lock()
						sum++
						mutex.Unlock()
					}
				}
				fmt.Println(i)
			}
		}(k)
	}
	wg.Wait()
	return sum
}

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
