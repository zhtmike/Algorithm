/*
Package assign1 count the number of inverions in a given sequence
Input: [3, 2, 5, 6, 20]
Output: 1 */
package assign1

// BruteForceCount counts the number of inversion by brute force
func BruteForceCount(seq []int) int {
	counter := 0
	for i, outerVal := range seq {
		for _, innerval := range seq[i:] {
			if outerVal > innerval {
				counter++
			}
		}
	}
	return counter
}

// InverseCount counts the number of inversion by divide and conquer
func InverseCount(seq []int) int {
	if len(seq) == 1 {
		return 0
	}
	middle := len(seq) / 2
	leftSeq := seq[:middle]
	rightSeq := seq[middle:]
	x := InverseCount(leftSeq)
	y := InverseCount(rightSeq)
	// calculate number of inversion, assume leftSeq and rightSeq is sorted
	z := countInvSplit(leftSeq, rightSeq, seq)
	return x + y + z
}

// countInvSplit
func countInvSplit(lseq, rseq, seq []int) int {
	i, j, k := 0, 0, 0
	inversion := 0
	tmp := make([]int, len(seq))

	// merge two sequences
	for i < len(lseq) && j < len(rseq) {
		if lseq[i] <= rseq[j] {
			tmp[k] = lseq[i]
			i++
		} else {
			tmp[k] = rseq[j]
			j++
			inversion += len(lseq) - i
		}
		k++
	}

	// merge the remainder
	for i < len(lseq) || j < len(rseq) {
		if i < len(lseq) {
			tmp[k] = lseq[i]
			i++
		} else {
			tmp[k] = rseq[j]
			j++
		}
		k++
	}

	copy(seq, tmp)
	return inversion
}
