package main

import "sort"

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
