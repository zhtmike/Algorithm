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
	arr := fileio.ReadWeightedAdjListFromText(os.Args[1])

	result := DijkstraAlgo(arr)
	// Get the shortest path of specific nodes
	// Index is changed from 1-based indexing to 0-based indexing
	fmt.Println(result[6], result[36], result[58],
		result[81], result[98], result[114],
		result[132], result[164], result[187], result[196],
	)
}
