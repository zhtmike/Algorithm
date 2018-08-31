package main

import (
	"bufio"
	"errors"
	"os"
	"strconv"
	"strings"
)

// ReadValueWeight read the value and weight for knapsack problem
func ReadValueWeight(src string) (size int, items [][]int, err error) {
	f, err := os.Open(src)
	if err != nil {
		return
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	numOfItems := 0
	line := 0
	row := make([]int, 2)
	for scanner.Scan() {
		text := scanner.Text()

		if line == 0 {
			vals := strings.Split(text, " ")
			size, err = strconv.Atoi(vals[0])
			if err != nil {
				return
			}
			numOfItems, err = strconv.Atoi(vals[1])
			if err != nil {
				return
			}
		} else {
			vals := strings.Split(text, " ")
			for i, v := range vals {
				row[i], err = strconv.Atoi(v)
				if err != nil {
					return
				}
			}
			// change to 0 based indexing
			items = append(items, make([]int, 2))
			copy(items[len(items)-1], row)
		}
		line++
	}

	if numOfItems != len(items) {
		return 0, nil, errors.New("incorrect of number of items")
	}
	return
}
