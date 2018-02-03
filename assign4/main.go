package main

import (
	"fmt"

	"../utils"
	"./src"
)

func main() {
	arr := utils.ReadEdgeListFromText("data/SCC.txt")
	result := assign4.KosarajuAlgo(arr)
	firstN := 5
	if len(result) < firstN {
		firstN = len(result)
	}
	fmt.Println(result[:firstN])
}
