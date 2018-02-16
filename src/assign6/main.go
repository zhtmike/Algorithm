package main

import (
	"fmt"
	"os"
	"runtime"
	"utils/fileio"
)

func main() {
	// arg 1 is the question 1's path , arg 2 is the question 2's
	if len(os.Args) < 3 {
		panic("Missing input files.")
	}
	runtime.GOMAXPROCS(runtime.NumCPU())

	// Question 1
	arr := fileio.ReadIntArrayFromText(os.Args[1])
	result1 := TwoSum(arr, runtime.NumCPU())
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
