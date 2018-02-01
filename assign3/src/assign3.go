/*
Package assign3 find the min cut in a graph
*/
package assign3

import (
	"sync"
	"math"
	"math/rand"
)

// GetMinCut calculates the minimum cuts of a graph, which is not guaranteed to be the minimum
func GetMinCut(adjlist [][]int) int {
	// Using concurrent threads, ref: https://stackoverflow.com/questions/24238820/parallel-for-loop
	var wg sync.WaitGroup

	repeats := int(math.Log(float64(len(adjlist)))*math.Pow(float64(len(adjlist)), 2))
	trials := make([]int, repeats)
	wg.Add(repeats)

	for i := 0; i < repeats; i++ {
		go func(i int) {
			trials[i] = copyAndCut(adjlist)
			defer wg.Done()
		}(i)
	}
	wg.Wait()

	// get the minimum of trials
	ans := int(1E9)
	for _, trial := range trials {
		if trial < ans {
			ans = trial
		}
	}
	return ans
}

func getCut(adjlist [][]int) int {
	for validV := getValidVertices(adjlist); len(validV) > 2; validV = getValidVertices(adjlist) {
		// choose a pair of vertices
		v0 := validV[rand.Intn(len(validV))] + 1
		e0 := adjlist[v0-1]
		v1 := e0[rand.Intn(len(e0))]
		e1 := adjlist[v1-1]

		// merge v1's edges (e1) to v0
		adjlist[v0-1] = append(e0, e1...)

		// delete the vertex v1, which is stored as the row v1-1 in the list
		adjlist[v1-1] = nil
		// scan all edges and update the old edges
		for _, edges := range adjlist {
			for i := range edges {
				if edges[i] == v1 {
					edges[i] = v0
				}
			}
		}

		// remove the self loops
		var tmp []int
		for _, vertex := range adjlist[v0-1] {
			if vertex != v0 {
				tmp = append(tmp, vertex)
			}
		}
		adjlist[v0-1] = tmp
	}
	validV := getValidVertices(adjlist)
	return len(adjlist[validV[0]])
}

func getValidVertices(adjlist [][]int) []int {
	var validv []int
	for i, edges := range adjlist {
		if edges != nil {
			validv = append(validv, i)
		}
	}
	return validv
}

func copyAndCut(adjlist [][]int) int {
	tmp := make([][]int, len(adjlist))
	for j := range adjlist {
		tmpSlice := make([]int, len(adjlist[j]))
		copy(tmpSlice, adjlist[j])
		tmp[j] = tmpSlice
	}
	return getCut(tmp)
}