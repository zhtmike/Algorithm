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

	arr := fileio.ReadJobListFromText(os.Args[1])

	// Question 1
	ans := Greedy(arr, true)
	fmt.Println(ans)

	// Question 2
	ans = Greedy(arr, false)
	fmt.Println(ans)
}
