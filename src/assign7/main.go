package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	if len(os.Args) < 3 {
		panic("Missing input file.")
	}

	arr, err := ReadJobListFromText(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}

	// Question 1
	ans := Greedy(arr, true)
	fmt.Println(ans)

	// Question 2
	ans = Greedy(arr, false)
	fmt.Println(ans)

	// Question 3
	arr2, err := ReadWEdgeListFromText(os.Args[2])
	if err != nil {
		log.Fatal(err)
	}
	ans = Prism(arr2)
	fmt.Println(ans)
}
