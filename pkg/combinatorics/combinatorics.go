package com

import . "github.com/GabiBizdoc/collections/pkg/stack"

// GenerateAllCombinations returns a list containing all possible combinations
// of the specified size from the given array.
func GenerateAllCombinations[T any](arr []T, size int) [][]T {
	var results [][]T
	result := make([]T, 0)

	var backtracking func(int)
	backtracking = func(start int) {
		if len(result) == size {
			tmp := make([]T, len(result))
			copy(tmp, result)
			results = append(results, tmp)
			return
		}

		for j := start; j < len(arr); j++ {
			result = append(result, arr[j])
			backtracking(j + 1)
			result = result[:len(result)-1]
		}
	}

	backtracking(0)
	return results
}

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

type CombinatorPullIter[T any] struct {
	size   int
	data   []T
	result Stack[T]
	stack  Stack[int]
}

func NewCombinatorPullIter[T any](data []T, size int) *CombinatorPullIter[T] {
	result := make(Stack[T], size)
	stack := make(Stack[int], 0, len(data))

	return &CombinatorPullIter[T]{
		size:   size,
		data:   data,
		result: result,
		stack:  stack,
	}
}

func (c *CombinatorPullIter[T]) Next() bool {
	if c.stack.IsEmpty() {
		for i := 0; i < c.size; i++ {
			c.stack.Push(i)
		}
		return true
	}

	for !c.stack.IsEmpty() {
		last := c.stack.Pop()

		limit := len(c.data) + c.stack.Size() - c.size
		if last < limit {
			c.stack.Push(last + 1)

			for c.stack.Size() < c.size {
				c.stack.Push(c.stack.Peek() + 1)
			}
			return true
		}
	}
	return false
}

func (c *CombinatorPullIter[T]) Value() []T {
	for i, j := range c.stack {
		c.result[i] = c.data[j]
	}
	return c.result
}

func (c *CombinatorPullIter[T]) CopyValue() []T {
	s := make([]T, len(c.stack))
	for i, j := range c.stack {
		s[i] = c.data[j]
	}
	return s
}
