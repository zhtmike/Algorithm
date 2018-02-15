package main

import (
	"fmt"
	"os"
	"utils/fileio"
)

func main() {
	if len(os.Args) < 3 {
		panic("Missing input files.")
	}
	// Question 1
	arr := fileio.ReadIntArrayFromText(os.Args[1])
	result1 := TwoSum(arr)
	fmt.Println(result1)

	// Question 2
	arr = fileio.ReadIntArrayFromText(os.Args[2])
	result2 := MedianMaintenance(arr)
	sum := 0
	for _, val := range result2 {
		sum += val
	}
	fmt.Println(sum % len(result2))
}
