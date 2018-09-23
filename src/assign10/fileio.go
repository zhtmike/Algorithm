package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
	"utils/graph"
)

// ReadWEdgeListFromText read the edge list from a txt file
func ReadWEdgeListFromText(src string) (adjmap graph.WAdjmap, err error) {
	f, err := os.Open(src)
	if err != nil {
		return
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	line := 0
	var totalNode int
	row := make([]int, 3)
	for scanner.Scan() {
		text := scanner.Text()

		if line == 0 {
			vals := strings.Split(text, " ")

			totalNode, err = strconv.Atoi(vals[0])
			if err != nil {
				return
			}

			adjmap = make(graph.WAdjmap, totalNode)
			for i := range adjmap {
				adjmap[i] = make(map[int]int)
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
			adjmap[row[0]-1][row[1]-1] = row[2]
		}
		line++
	}
	return
}
