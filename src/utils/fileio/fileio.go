package fileio

import (
	"bufio"
	"encoding/csv"
	"os"
	"strconv"
	"strings"
	"utils/graph"
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
// Index will be changed from one-based numbering to zero-based numbering
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
		panic(err)
	}

	arr := make([][]int, len(data))
	for i, row := range data {
		vertex, err := strconv.Atoi(row[0])
		// Change to zero-based numbering
		vertex--
		if err != nil || vertex != i {
			panic(err)
		}
		// ignore '\n' at each row
		row = row[1 : len(row)-1]
		tmp := make([]int, len(row))
		for j, num := range row {
			tmp[j], err = strconv.Atoi(num)
			// Change to zero-based numbering
			tmp[j]--
			if err != nil {
				panic(err)
			}
		}
		arr[i] = tmp
	}
	return arr
}

// ReadEdgeListFromText read the edge list from a txt file
func ReadEdgeListFromText(src string) [][]int {
	f, err := os.Open(src)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	reader := csv.NewReader(f)
	// Using space as delimiter
	reader.Comma = ' '
	// Read all data
	data, err := reader.ReadAll()
	if err != nil {
		panic(err)
	}

	// Using dict to store the edges temporary
	tmpArr := make(map[int][]int)
	for _, row := range data {
		vertex, err := strconv.Atoi(row[0])
		if err != nil {
			panic(err)
		}
		// Change to zero-based numbering
		vertex--
		outvertex, err := strconv.Atoi(row[1])
		if err != nil {
			panic(err)
		}
		// Change to zero-based numbering
		outvertex--
		tmpArr[vertex] = append(tmpArr[vertex], outvertex)
	}

	// Find the maximum edge number
	maximum := 0
	for key := range tmpArr {
		if key > maximum {
			maximum = key
		}
	}

	// Convert map to nested list
	arr := make([][]int, maximum+1)
	for key, value := range tmpArr {
		arr[key] = value
	}
	return arr
}

// ReadWeightedAdjListFromText read the weighted adjacency list from a txt file
// Index will be changed from one-based numbering to zero-based numbering
func ReadWeightedAdjListFromText(src string) graph.WAdjlist {
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
		panic(err)
	}

	arr := make(graph.WAdjlist, len(data))
	for i, row := range data {
		vertex, err := strconv.Atoi(row[0])
		// Change to zero-based numbering
		vertex--
		if err != nil || vertex != i {
			panic(err)
		}
		// ignore '\n' at each row
		row = row[1 : len(row)-1]
		tmp := make([]graph.WEdge, len(row))
		for j, num := range row {
			pairs := strings.Split(num, ",")

			ind, err := strconv.Atoi(pairs[0])
			// Change to zero-based numbering
			ind--
			if err != nil {
				panic(err)
			}

			weight, err := strconv.Atoi(pairs[1])
			if err != nil {
				panic(err)
			}

			tmp[j] = graph.WEdge{
				Vertex: ind,
				Weight: weight}
		}
		arr[i] = tmp
	}
	return arr
}

// ReadJobListFromText read the edge list from a txt file
func ReadJobListFromText(src string) [][]int {
	f, err := os.Open(src)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	line := 0
	var totalNum int
	var jobList [][]int
	for scanner.Scan() {
		text := scanner.Text()
		tmp := make([]int, 2)
		if line == 0 {
			totalNum, err = strconv.Atoi(text)
			if err != nil {
				panic(err)
			}
			jobList = make([][]int, totalNum)
		} else {
			vals := strings.Split(text, " ")

			tmp[0], err = strconv.Atoi(vals[0])
			if err != nil {
				panic(err)
			}
			tmp[1], err = strconv.Atoi(vals[1])
			if err != nil {
				panic(err)
			}

			jobList[line-1] = tmp
		}
		line++
	}
	return jobList
}

// ReadWEdgeListFromText read the edge list from a txt file
func ReadWEdgeListFromText(src string) graph.WAdjmap {
	f, err := os.Open(src)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	line := 0
	var totalNode int
	var adjmap graph.WAdjmap
	row := make([]int, 3)
	for scanner.Scan() {
		text := scanner.Text()

		if line == 0 {
			vals := strings.Split(text, " ")

			totalNode, err = strconv.Atoi(vals[0])
			if err != nil {
				panic(err)
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
					panic(err)
				}
			}
			// change to 0 based indexing
			adjmap[row[0]-1][row[1]-1] = row[2]
			adjmap[row[1]-1][row[0]-1] = row[2]
		}
		line++
	}

	return adjmap
}
