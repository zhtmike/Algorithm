package main

import (
	"errors"
	"sort"
	"utils/unionfind"
)

// KCluster implements max-spacing k-clustering.
func KCluster(totalNode int, edges [][]int, clusterNum int) (int, error) {
	// sort the edges by cost
	sort.Slice(edges, func(i, j int) bool {
		if edges[i][2] <= edges[j][2] {
			return true
		}
		return false
	})

	count := totalNode
	union := unionfind.New(totalNode)
	for _, v := range edges {
		i, j, cost := v[0], v[1], v[2]
		if union.Find(i) != union.Find(j) {
			if count == clusterNum {
				return cost, nil
			}
			union.Merge(i, j)
			count--
		}
	}
	return 0, errors.New("failed in finding the K-cluster")
}

// KClusterBit implements max-spacing k-clustering for bit location.
func KClusterBit(length int, bitList []int) int {
	// filter the unique elements in bitList
	bitExitMap := make(map[int]bool)
	for _, v := range bitList {
		bitExitMap[v] = true
	}
	var uniBitList []int
	for k := range bitExitMap {
		uniBitList = append(uniBitList, k)
	}

	// store the list as map, the value is the location
	bitMap := make(map[int]int)
	for i, v := range uniBitList {
		bitMap[v] = i
	}

	// create masks
	masks := append(createOneMask(uint(length)), createTwoMasks(uint(length))...)

	// create union
	union := unionfind.New(len(uniBitList))
	for i, u := range uniBitList {
		for _, mask := range masks {
			v := u ^ mask
			j, ok := bitMap[v]
			if ok && union.Find(i) != union.Find(j) {
				union.Merge(i, j)
			}
		}
	}
	return union.Len()
}

func createTwoMasks(length uint) (masks []int) {
	for i := uint(0); i < length; i++ {
		for j := i + 1; j < length; j++ {
			masks = append(masks, 1<<i|1<<j)
		}
	}
	return
}

func createOneMask(length uint) (masks []int) {
	for i := uint(0); i < length; i++ {
		masks = append(masks, 1<<i)
	}
	return
}
