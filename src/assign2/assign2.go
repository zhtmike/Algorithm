/*
Sort the number by quick sort
*/
package main

import (
	"math/rand"
	"sort"
)

type fn func(arr []int) int

// QuickSort sort the arr randomly
func QuickSort(arr []int, num *int, choosePivot fn) {
	if len(arr) <= 1 {
		return
	}
	*num += len(arr) - 1
	p := choosePivot(arr)
	z := partition(arr, p)
	QuickSort(arr[:z], num, choosePivot)
	QuickSort(arr[z+1:], num, choosePivot)
}

// ChooseRandomPivot choose the pivot in a array randomly
func ChooseRandomPivot(arr []int) int {
	return rand.Intn(len(arr))
}

// ChooseFirstPivot choose the first element in a array as a pivot
func ChooseFirstPivot(arr []int) int {
	return 0
}

// ChooseLastPivot choose the last element in a array as a pivot
func ChooseLastPivot(arr []int) int {
	return len(arr) - 1
}

// ChooseMedianPivot choose the median of three element in a array as a pivot
func ChooseMedianPivot(arr []int) int {
	x, y, z := arr[0], arr[(len(arr)-1)/2], arr[len(arr)-1]
	temp := []int{x, y, z}
	sort.Ints(temp)
	switch temp[1] {
	case x:
		return 0
	case y:
		return (len(arr) - 1) / 2
	default:
		return len(arr) - 1
	}
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
