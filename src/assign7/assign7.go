package main

import (
	"container/heap"
	"math/rand"
	"sort"
	"utils/graph"
	"utils/vheap"
)

type jobList [][]int

func (s jobList) Len() int {
	return len(s)
}

func (s jobList) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

// compare the job by weight-length
type jobListDiff struct {
	jobList
}

func (job jobListDiff) Less(i, j int) bool {
	s := job.jobList
	diff := (s[i][0] - s[i][1]) - (s[j][0] - s[j][1])
	if diff == 0 {
		return s[i][0] < s[j][0]
	}
	return diff < 0
}

// compare the job by weight/length
type jobListRatio struct {
	jobList
}

func (job jobListRatio) Less(i, j int) bool {
	s := job.jobList
	return float64(s[i][0])/float64(s[i][1]) < float64(s[j][0])/float64(s[j][1])
}

// Greedy minimize the weighted completion time of a job list
func Greedy(arr [][]int, useDiff bool) int {
	if useDiff {
		sort.Sort(sort.Reverse(jobListDiff{jobList(arr)}))
	} else {
		sort.Sort(sort.Reverse(jobListRatio{jobList(arr)}))
	}

	completeT := make([]int, len(arr))
	for i, v := range arr {
		if i > 1 {
			completeT[i] = v[1] + completeT[i-1]
		} else {
			completeT[i] = v[1]
		}
	}
	totalCost := 0
	for i, v := range completeT {
		totalCost += v * arr[i][0]
	}
	return totalCost
}

// Prism uses minimum spanning tree algorithm and calculate the total cost of the minimum spanning tree
func Prism(g graph.WAdjmap) int {
	checkList := make(map[int]bool)

	// Pick up a random node and set up the heap list
	var pq vheap.PriorityQueue
	pq.Hp = make([]*vheap.Node, len(g))
	pq.Indices = make([]int, len(g))

	initialNode := rand.Intn(len(g))
	checkList[initialNode] = true
	for i := 0; i < len(g); i++ {
		v, ok := g[initialNode][i]
		if !ok {
			v = int(10E9)
		}
		pq.Indices[i] = i
		pq.Hp[i] = &vheap.Node{
			ID:       i,
			Distance: v,
		}
	}
	heap.Init(&pq)

	// iterate the remaining node until all nodes are explored.
	// the last one corresponds to the initial node
	totalCost := 0
	for pq.Len() > 1 {
		node := heap.Pop(&pq).(*vheap.Node)
		id, distance := node.ID, node.Distance
		totalCost += distance
		checkList[id] = true
		for k, v := range g[id] {
			ok := checkList[k]
			if !ok {
				pq.Update(k, min(v, pq.Hp[pq.Indices[k]].Distance))
			}
		}
	}
	return totalCost
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
