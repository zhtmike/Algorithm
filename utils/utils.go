package utils

import (
	"encoding/csv"
	"io"
	"os"
	"strconv"
)

// ReadIntArrayFromText read the array from a txt file
func ReadIntArrayFromText(src string) []int {
	f, err := os.Open(src)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	data, err := csv.NewReader(f).ReadAll()
	if err != nil {
		panic(err)
	}

	arr := make([]int, len(data))
	for i, v := range data {
		x, err := strconv.Atoi(v[0])
		if err != nil {
			panic(err)
		}
		arr[i] = x
	}
	return arr
}

// ReadAdjListFromText read the adjacency list from a txt file
func ReadAdjListFromText(src string) [][]int {
	f, err := os.Open(src)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	reader := csv.NewReader(f)
	// Using tab as delimiter
	reader.Comma = '\t'
	// Ignore Unequal number of elements per each row
	reader.FieldsPerRecord = -1
	// Read all data
	data, err := reader.ReadAll()
	if err != nil {
		if err != io.EOF {
			panic(err)
		}
	}

	arr := make([][]int, len(data))
	for i, row := range data {
		vertex, err := strconv.Atoi(row[0])
		if err != nil || vertex != i+1 {
			panic(err)
		}
		// ignore '\n' at each row
		row = row[1 : len(row)-1]
		tmp := make([]int, len(row))
		for j, num := range row {
			tmp[j], err = strconv.Atoi(num)
			if err != nil {
				panic(err)
			}
		}
		arr[i] = tmp
	}
	return arr
}
