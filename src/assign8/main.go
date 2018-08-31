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
	totalNode, edgeList, err := ReadWEdgeList(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	ans, err := KCluster(totalNode, edgeList, 4)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(ans)

	// question 2
	bitlength, bitsList, err := ReadBitList(os.Args[2])
	if err != nil {
		log.Fatal(err)
	}
	ans = KClusterBit(bitlength, bitsList)
	fmt.Println(ans)
}
