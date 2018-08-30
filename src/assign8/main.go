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

	totalNode, edgeList := fileio.ReadWEdgeList(os.Args[1])
	ans := KCluster(totalNode, edgeList, 4)
	fmt.Println(ans)
}
