package utils

import (
	"encoding/csv"
	"os"
	"strconv"
)

// ReadIntArrayFromText read the array from a txt file
func ReadIntArrayFromText(src string) []int {
	var arr []int

	f, err := os.Open(src)
	if err != nil {
		panic(err)
	}
	defer f.Close()

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
	return arr
}
