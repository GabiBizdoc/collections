package collections

import (
	. "github.com/GabiBizdoc/collections/pkg/stack"
	"slices"
)

type CombinatorPullIter[T ~[]E, E any] struct {
	k      int
	arr    T
	result Stack[E]
	stack  Stack[int]
}

func NewCombinatorPullIter[T ~[]E, E any](arr T, k int) *CombinatorPullIter[T, E] {
	result := make([]E, 0, k)
	stack := make(Stack[int], 0, len(arr))

	return &CombinatorPullIter[T, E]{
		k:      k,
		arr:    arr,
		result: result,
		stack:  stack,
	}
}

func (c *CombinatorPullIter[T, E]) Next() bool {
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

	return false
}

func (c *CombinatorPullIter[T, E]) Value() T {
	return T(c.result)
}

func (c *CombinatorPullIter[T, E]) CopyValue() T {
	return slices.Clone(c.Value())
}
