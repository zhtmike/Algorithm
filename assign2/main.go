package main

import (
	"fmt"

	"./src"

	"../utils"
)

func main() {
	arr := utils.ReadIntArrayFromText("data/QuickSort.txt")
	arrBak := make([]int, len(arr))
	copy(arrBak, arr)

	// First Question
	num := 0
	assign2.QuickSort(arr, &num, assign2.ChooseFirstPivot)
	fmt.Println(num)

	// Second Question
	num = 0
	copy(arr, arrBak)
	assign2.QuickSort(arr, &num, assign2.ChooseLastPivot)
	fmt.Println(num)

	// third Question
	num = 0
	copy(arr, arrBak)
	assign2.QuickSort(arr, &num, assign2.ChooseMedianPivot)
	fmt.Println(num)
}
