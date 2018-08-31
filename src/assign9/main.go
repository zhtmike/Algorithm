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
	// question 1
	size, arr, err := ReadValueWeight(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	ans := Knapsack(size, arr)
	fmt.Println(ans)

	// question 2
	size, arr, err = ReadValueWeight(os.Args[2])
	if err != nil {
		log.Fatal(err)
	}
	ans = Knapsack(size, arr)
	fmt.Println(ans)
}
