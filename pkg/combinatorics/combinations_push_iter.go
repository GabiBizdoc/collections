package com

type CombinatorPushIter[T ~[]E, E any] struct {
	arr T
	k   int
}

func NewCombinatorPushIter[T ~[]E, E any](arr T, k int) CombinatorPushIter[T, E] {
	return CombinatorPushIter[T, E]{arr: arr, k: k}
}

// ForEach returns a list containing all possible combinations
// of the specified size from the given array.
func (c CombinatorPushIter[T, E]) ForEach(yield func(T)) {
	arr := c.arr
	size := c.k

	result := make(T, 0)

	var backtracking func(int)
	backtracking = func(start int) {
		if len(result) == size {
			yield(result)
			return
		}

		end := len(arr) + len(result) + 1 - size
		for j := start; j < end; j++ {
			result = append(result, arr[j])
			backtracking(j + 1)
			result = result[:len(result)-1]
		}
	}

	backtracking(0)
}
