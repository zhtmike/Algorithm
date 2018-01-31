package main

import (
	"fmt"

	"../utils"
	"./src"
)

func main() {
	arr := utils.ReadAdjListFromText("data/kargerMinCut.txt")
	ans := assign3.GetMinCut(arr)
	fmt.Println(ans)
}
