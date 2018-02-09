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
	arr := fileio.ReadEdgeListFromText(os.Args[1])
	result := KosarajuAlgo(arr)
	firstN := 5
	if len(result) < firstN {
		firstN = len(result)
	}
	fmt.Println(result[:firstN])
}
