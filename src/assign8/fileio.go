package main

import (
	"bufio"
	"errors"
	"os"
	"strconv"
	"strings"
)

// ReadWEdgeList read the edge list from a txt file
func ReadWEdgeList(src string) (totalNode int, edgeList [][]int, err error) {
	f, err := os.Open(src)
	if err != nil {
		return
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	line := 0
	row := make([]int, 3)
	for scanner.Scan() {
		text := scanner.Text()

		if line == 0 {
			totalNode, err = strconv.Atoi(text)
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
			edgeList = append(edgeList, make([]int, 3))
			row[0]--
			row[1]--
			copy(edgeList[len(edgeList)-1], row)
		}
		line++
	}
	return
}

// ReadBitList read the bit list from a txt file
func ReadBitList(src string) (bitlength int, bits []int, err error) {
	f, err := os.Open(src)
	if err != nil {
		return
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	line := 0
	totalNode := -1
	for scanner.Scan() {
		text := scanner.Text()

		if line == 0 {
			vals := strings.Split(text, " ")
			totalNode, err = strconv.Atoi(vals[0])
			if err != nil {
				return
			}
			bitlength, err = strconv.Atoi(vals[1])
			if err != nil {
				return
			}
		} else {
			vals := strings.Join(strings.Split(text, " "), "")
			bit, err := strconv.ParseInt(vals, 2, 64)
			if err != nil {
				return 0, nil, err
			}
			bits = append(bits, int(bit))
		}
		line++
	}
	if totalNode != len(bits) {
		return 0, nil, errors.New("inconsistent number of nodes")
	}
	return
}
