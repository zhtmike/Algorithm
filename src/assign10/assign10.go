package main

import (
	"errors"
	"utils/graph"
)

const maxint = int(1E9)

// FloydWarshall Algorithm compute the shortest path of graph
func FloydWarshall(g graph.WAdjmap) (int, error) {
	n := len(g)
	mat := createMat(n, g)
	ans := computeMin(n, mat)
	minVal := maxint
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if i == j && ans[i][j] < 0 {
				return 0, errors.New("graph contains negative-cost cycles")
			}
			minVal = min(minVal, ans[i][j])
		}
	}
	return minVal, nil
}

func createMat(n int, g graph.WAdjmap) (mat [][][]int) {
	mat = make([][][]int, n)
	for i := 0; i < n; i++ {
		mat[i] = make([][]int, n)
		for j := 0; j < n; j++ {
			mat[i][j] = make([]int, 2)
			if i != j {
				val, ok := g[i][j]
				if ok {
					mat[i][j][1] = val
				} else {
					mat[i][j][1] = maxint
				}
			}
		}
	}
	return
}

func computeMin(n int, mat [][][]int) [][]int {
	for k := 0; k < n; k++ {
		for i := 0; i < n; i++ {
			for j := 0; j < n; j++ {
				mat[i][j][k%2] = min(mat[i][j][(k+1)%2], mat[i][k][(k+1)%2]+mat[k][j][(k+1)%2])
			}
		}
	}
	ans := make([][]int, n)
	k := (n - 1) % 2
	for i := 0; i < n; i++ {
		ans[i] = make([]int, n)
		for j := 0; j < n; j++ {
			ans[i][j] = mat[i][j][k]
		}
	}
	return ans
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}
