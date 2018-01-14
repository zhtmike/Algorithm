package main

import (
	"fmt"

	"./src"

	"../utils"
)

func main() {
	arr := utils.ReadIntArrayFromText("data/IntegerArray.txt")
	fmt.Println(assign1.InverseCount(arr))
}
