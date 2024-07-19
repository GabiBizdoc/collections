package com

import (
	. "github.com/GabiBizdoc/collections/pkg/stack"
)

type CombinatorPullIter[T any] struct {
	done   bool
	k      int
	arr    []T
	result Stack[T]
	stack  Stack[int]
}

func NewCombinatorPullIter[T any](arr []T, k int) *CombinatorPullIter[T] {
	result := make([]T, 0, k)
	stack := make(Stack[int], 0, len(arr))

	return &CombinatorPullIter[T]{
		k:      k,
		arr:    arr,
		result: result,
		stack:  stack,
	}
}

func (c *CombinatorPullIter[T]) Next() bool {
	c.checkDone()

	if c.stack.IsEmpty() {
		for i := 0; i < c.k; i++ {
			c.stack.Push(i)
			c.result.Push(c.arr[i])
		}
		return true
	}

	for !c.stack.IsEmpty() {
		last := c.stack.Pop()
		c.result.Pop()

		limit := len(c.arr) + c.stack.Size() - c.k
		if last < limit {
			c.stack.Push(last + 1)
			c.result.Push(c.arr[last+1])

			for c.stack.Size() < c.k {
				c.stack.Push(c.stack.Peek() + 1)
				c.result.Push(c.arr[c.stack.Peek()])
			}
			return true
		}
	}

	c.done = true
	return false
}

func (c *CombinatorPullIter[T]) Value() []T {
	c.checkDone()
	return c.result
}

func (c *CombinatorPullIter[T]) CopyValue() []T {
	c.checkDone()

	s := make([]T, len(c.stack))
	copy(s, c.result)

	return s
}

func (c *CombinatorPullIter[T]) checkDone() {
	if c.done {
		panic("read after done")
	}
}
