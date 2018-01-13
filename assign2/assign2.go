/*
Package assign2 sort the number by random sort
Input: [3, 2, 5, 6, 20]
Output: [2, 3, 5, 6, 20] */
package assign2

import (
	"math/rand"
)

// QuickSort sort the arr randomly
func QuickSort(arr []int) {
	if len(arr) <= 1 {
		return
	}
	p := choosePivot(arr)
	z := partition(arr, p)
	QuickSort(arr[:z])
	QuickSort(arr[z+1:])
}

// choosePivot choose the pivot in a array randomly
func choosePivot(arr []int) int {
	return rand.Intn(len(arr))
}

// partition reorder the array according to the pivot
func partition(arr []int, p int) int {
	arr[0], arr[p] = arr[p], arr[0]
	i := 1
	for j := 1; j < len(arr); j++ {
		if arr[j] < arr[0] {
			arr[i], arr[j] = arr[j], arr[i]
			i++
		}
	}
	arr[0], arr[i-1] = arr[i-1], arr[0]
	return i - 1
}
