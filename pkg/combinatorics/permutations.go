package collections

import (
	"slices"
)

// PermutationsRecursivePushIter is a recursive push based implementation of
// https://en.wikipedia.org/wiki/Heap%27s_algorithm
type PermutationsRecursivePushIter[T ~[]E, E any] struct {
	data T
}

func NewPermutationsRecursivePushIter[T ~[]E, E any](data T) *PermutationsRecursivePushIter[T, E] {
	return &PermutationsRecursivePushIter[T, E]{data: data}
}

func (p *PermutationsRecursivePushIter[T, E]) ForEach(yield func(T)) {
	data := slices.Clone(p.data)

	if len(data) < 2 {
		yield(data)
		return
	}

	var generate func(int)
	generate = func(k int) {
		if k == 1 {
			yield(data)
			return
		}
		generate(k - 1)

		for i := range k - 1 {
			if k&1 == 0 {
				data[i], data[k-1] = data[k-1], data[i]
			} else {
				data[0], data[k-1] = data[k-1], data[0]
			}
			generate(k - 1)
		}
	}

	generate(len(data))
}

// PermutationsIterativePushIter is a non-recursive pull based implementation of PermutationsRecursivePushIter
type PermutationsIterativePushIter[T ~[]E, E any] struct {
	data T
}

func NewPermutationsIterativePushIter[T ~[]E, E any](data T) *PermutationsIterativePushIter[T, E] {
	return &PermutationsIterativePushIter[T, E]{data: data}
}

func (p *PermutationsIterativePushIter[T, E]) ForEach(yield func(T)) {
	data := slices.Clone(p.data)
	c := make([]int, len(p.data))

	yield(data)

	i := 1
	for i < len(data) {
		if c[i] < i {
			if i&1 == 0 {
				data[0], data[i] = data[i], data[0]
			} else {
				data[c[i]], data[i] = data[i], data[c[i]]
			}

			yield(data)

			c[i] += 1
			i = 1
		} else {
			c[i] = 0
			i += 1
		}
	}
}

// PermutationsIterativePullIter is a non-recursive pull based implementation of PermutationsRecursivePushIter
type PermutationsIterativePullIter[T ~[]E, E any] struct {
	started bool
	arr     T
	stack   []int
}

func NewPermutationsIterativePullIter[T ~[]E, E any](arr T) *PermutationsIterativePullIter[T, E] {
	stack := make([]int, len(arr))

	return &PermutationsIterativePullIter[T, E]{
		started: false,
		arr:     slices.Clone(arr),
		stack:   stack,
	}
}

func (p *PermutationsIterativePullIter[T, E]) Next() bool {
	if !p.started {
		p.started = true
		return true
	}

	if len(p.arr) < 2 {
		return false
	}

	data := p.arr
	stack := p.stack

	for i := range len(data) {
		if stack[i] < i {
			if i&1 == 0 {
				data[0], data[i] = data[i], data[0]
			} else {
				data[stack[i]], data[i] = data[i], data[stack[i]]
			}

			stack[i] += 1
			return true
		}

		stack[i] = 0
	}

	return false
}

func (p *PermutationsIterativePullIter[T, E]) Value() T {
	return p.arr
}

func (p *PermutationsIterativePullIter[T, E]) CopyValue() T {
	return slices.Clone(p.Value())
}
