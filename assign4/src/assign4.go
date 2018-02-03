/*
Package assign4 find the strong connected component by Kosaraju's algorithm
*/
package assign4

import "sort"

// KosarajuAlgo returns the SCCs' sizes in sorted order
func KosarajuAlgo(adjlist [][]int) []int {
	// declare the global variables
	finishTime := 0
	exploredNodes := make([]bool, len(adjlist))
	finishTimes := make([]int, len(adjlist))

	// first iteration on the reversed graph, calculate the finish time for each node
	reversedAdj := reverse(adjlist)
	for i := len(reversedAdj) - 1; i >= 0; i-- {
		if !exploredNodes[i] {
			dfs(reversedAdj, i, exploredNodes, finishTimes, &finishTime)
		}
	}

	// Replace the graph node by finish times
	finishTAdj := replace(adjlist, finishTimes)

	// second iteration on the graph, calculate the number of element in each SCC
	var counts []int
	exploredNodes = make([]bool, len(adjlist))
	for i := len(finishTAdj) - 1; i >= 0; i-- {
		if !exploredNodes[i] {
			count := 0
			dfs(finishTAdj, i, exploredNodes, finishTimes, &count)
			counts = append(counts, count)
		}
	}

	// Sort counts in reversed order
	sort.Sort(sort.Reverse(sort.IntSlice(counts)))
	return counts
}

func dfs(g [][]int, node int, expNodes []bool, fsTimes []int, t *int) {
	expNodes[node] = true
	for _, v := range g[node] {
		if !expNodes[v] {
			dfs(g, v, expNodes, fsTimes, t)
		}
	}
	fsTimes[node] = *t
	*t++
}

// Reverse the graph
func reverse(adjlist [][]int) [][]int {
	newAdjlist := make([][]int, len(adjlist))
	for i, vertices := range adjlist {
		for _, vertice := range vertices {
			newAdjlist[vertice] = append(newAdjlist[vertice], i)
		}
	}
	return newAdjlist
}

// Replace the graph by the finish times
func replace(adjlist [][]int, fsTimes []int) [][]int {
	newAdjlist := make([][]int, len(adjlist))
	for i, vertices := range adjlist {
		for _, vertice := range vertices {
			newAdjlist[fsTimes[i]] = append(newAdjlist[fsTimes[i]], fsTimes[vertice])
		}
	}
	return newAdjlist
}
