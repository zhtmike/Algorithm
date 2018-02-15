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

	// Question2 2
	arr := fileio.ReadIntArrayFromText(os.Args[1])
	result := MedianMaintenance(arr)
	sum := 0
	for _, val := range result {
		sum += val
	}
	fmt.Println(sum % len(result))
}
