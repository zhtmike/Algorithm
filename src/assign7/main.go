package main

import (
	"fmt"
	"os"
	"utils/fileio"
)

func main() {
	if len(os.Args) < 3 {
		panic("Missing input file.")
	}

	arr := fileio.ReadJobListFromText(os.Args[1])

	// Question 1
	ans := Greedy(arr, true)
	fmt.Println(ans)

	// Question 2
	ans = Greedy(arr, false)
	fmt.Println(ans)

	arr2 := fileio.ReadWEdgeListFromText(os.Args[2])
	ans = Prism(arr2)
	fmt.Println(ans)
}
