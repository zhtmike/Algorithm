package main

import (
	"fmt"
	"os"

	"utils/fileio"
)

func main() {
	if len(os.Args) < 2 {
		panic("Missing input file.")
	}
	arr := fileio.ReadIntArrayFromText(os.Args[1])
	arrBak := make([]int, len(arr))
	copy(arrBak, arr)

	// First Question
	num := 0
	QuickSort(arr, &num, ChooseFirstPivot)
	fmt.Println(num)

	// Second Question
	num = 0
	copy(arr, arrBak)
	QuickSort(arr, &num, ChooseLastPivot)
	fmt.Println(num)

	// third Question
	num = 0
	copy(arr, arrBak)
	QuickSort(arr, &num, ChooseMedianPivot)
	fmt.Println(num)
}
