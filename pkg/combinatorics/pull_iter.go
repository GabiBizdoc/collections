package com

import . "github.com/GabiBizdoc/collections/pkg/stack"

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
