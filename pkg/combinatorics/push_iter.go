package com

type CombinatorPushIter[T any] struct {
	arr  []T
	size int
}

func NewCombinatorPushIter[T any](arr []T, size int) CombinatorPushIter[T] {
	return CombinatorPushIter[T]{arr: arr, size: size}
}

// ForEach returns a list containing all possible combinations
// of the specified size from the given array.
func (x CombinatorPushIter[T]) ForEach(yield func([]T)) {
	arr := x.arr
	size := x.size

	result := make([]T, 0)

	var backtracking func(int)
	backtracking = func(start int) {
		if len(result) == size {
			yield(result)
			return
		}

		for j := start; j < len(arr); j++ {
			result = append(result, arr[j])
			backtracking(j + 1)
			result = result[:len(result)-1]
		}
	}

	backtracking(0)
}
