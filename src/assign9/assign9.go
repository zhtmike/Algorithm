package main

// Knapsack calculates optimal where total weight is less than knapsack size
func Knapsack(size int, items [][]int) int {
	numOfItems := len(items)
	mat := make([][]int, 2)
	for i := range mat {
		mat[i] = make([]int, size+1)
	}
	for i := 0; i < numOfItems; i++ {
		for x := 0; x <= size; x++ {
			val, weight := items[i][0], items[i][1]
			if weight > x {
				mat[(i+1)%2][x] = mat[i%2][x]
			} else {
				mat[(i+1)%2][x] = max(mat[i%2][x], mat[i%2][x-weight]+val)
			}
		}
	}
	return mat[numOfItems%2][size]
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}
