package com

import (
	"reflect"
	"slices"
	"testing"
)

func TestPermutations(t *testing.T) {
	tests := []struct {
		name string
		data []int
	}{
		{"Integers", []int{1, 2, 3, -1, -2, -3}},
		{"Single Element", []int{1}},
		{"Empty Slice", []int{}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			data := tt.data
			v1 := make([][]int, 0)
			v2 := make([][]int, 0)
			v3 := make([][]int, 0)

			NewPermutationsRecursivePushIter(data).ForEach(func(result []int) {
				v1 = append(v1, slices.Clone(result))
			})

			NewPermutationsIterativePushIter(data).ForEach(func(result []int) {
				v2 = append(v2, slices.Clone(result))
			})

			iter := NewPermutationsIterativePullIter(data)
			for iter.Next() {
				v3 = append(v3, iter.CopyValue())
			}

			if !reflect.DeepEqual(v1, v2) {
				t.Errorf("Recursive and Iterative Push results not equal for %v", data)
			}

			if !reflect.DeepEqual(v1, v3) {
				t.Errorf("Recursive Push and Iterative Pull results not equal for %v", data)
			}
		})
	}
}

func BenchmarkPermutations(b *testing.B) {
	size := 10
	data := make([]int, size)

	for i := range data {
		data[i] = i + 1
	}

	b.Run("PermutationsRecursivePushIter", func(b *testing.B) {
		for range b.N {
			iter := NewPermutationsRecursivePushIter(data)
			iter.ForEach(func(result []int) {
				_ = result
			})
		}
	})

	b.Run("PermutationsIterativePushIter", func(b *testing.B) {
		for range b.N {
			iter := NewPermutationsIterativePushIter(data)
			iter.ForEach(func(result []int) {
				_ = result
			})
		}
	})

	b.Run("PermutationsIterativePullIter", func(b *testing.B) {
		for range b.N {
			iter := NewPermutationsIterativePullIter(data)
			for iter.Next() {
				_ = iter.Value()
			}
		}
	})
}
