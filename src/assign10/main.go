package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	if len(os.Args) < 4 {
		panic("Missing input file.")
	}

	for i := 1; i <= 3; i++ {
		arr, err := ReadWEdgeListFromText(os.Args[i])
		if err != nil {
			log.Fatal(err)
		}
		ans, err := FloydWarshall(arr)
		if err != nil {
			fmt.Println("NULL")
		} else {
			fmt.Println(ans)
		}
	}
}
