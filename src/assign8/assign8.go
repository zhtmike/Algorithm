package main

import (
	"sort"
	"utils/unionfind"
)

// KCluster implements max-spacing k-clustering.
func KCluster(totalNode int, edges [][]int, clusterNum int) int {
	sort.Slice(edges, func(i, j int) bool {
		if edges[i][2] <= edges[j][2] {
			return true
		}
		return false
	})

	count := totalNode
	union := unionfind.CreateUnionFind(totalNode)
	for _, v := range edges {
		i, j, cost := v[0], v[1], v[2]
		if union.Find(i) != union.Find(j) {
			if count == clusterNum {
				return cost
			}
			union.Merge(i, j)
			count--
		}
	}
	return -1
}
