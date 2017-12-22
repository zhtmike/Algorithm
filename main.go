package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"

	"./assign1"
)

func main() {
	var src string
	var arr []int

	if len(os.Args) > 1 {
		src = os.Args[1]
	} else {
		return
	}

	f, err := os.Open(src)
	if err != nil {
		panic(err)
	}
	data, err := csv.NewReader(f).ReadAll()
	if err != nil {
		panic(err)
	}

	for _, v := range data {
		x, err := strconv.Atoi(v[0])
		if err != nil {
			panic(err)
		}
		arr = append(arr, x)
	}
	f.Close()

	result := assign1.InverseCount(arr)
	fmt.Println(result)
}
