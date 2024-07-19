package com_test

import (
	com "github.com/GabiBizdoc/collections/pkg/combinatorics"
	"reflect"
	"slices"
	"testing"
)

func TestCombinations(t *testing.T) {
	type testCase struct {
		name  string
		data  []int
		sizes []int
	}
	tests := []testCase{
		{
			name:  "case 1",
			data:  []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 32, 43, 54, 45},
			sizes: []int{1, 2, 3, 4, 5, 6, 7, 10},
		},
		{
			name:  "case 2",
			data:  []int{1, 2, 3, 4, 5},
			sizes: []int{5, 4, 3, 2, 1},
		},
		{
			name:  "case 3",
			data:  []int{1, 2, 3, 4, 23, 4, 5, 6, 5, 6, 7, 8, 9, 6, 3, 54, 45, 2, 3, 4, 5, 4},
			sizes: []int{7, 14, 2},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			data := tt.data
			for _, size := range tt.sizes {
				results1 := com.GenerateAllCombinations(data, size)

				results2 := make([][]int, 0, len(results1))
				pushIter := com.NewCombinatorPushIter(data, size)
				pushIter.ForEach(func(value []int) {
					results2 = append(results2, slices.Clone(value))
				})

				results3 := make([][]int, 0, len(results1))
				pullIter := com.NewCombinatorPullIter(data, size)
				for pullIter.Next() {
					results3 = append(results3, pullIter.CopyValue())
				}

				ok := reflect.DeepEqual(results1, results2) && reflect.DeepEqual(results2, results3)
				if !ok {

					t.Log(results1)
					t.Log(results2)
					t.Log(results3)
					t.Errorf("Combinations failed for size: %d", size)
				}
			}
		})
	}
}

func BenchmarkCombinatorics(b *testing.B) {
	//arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 6, 3, 2, 3, 4, 5, 10, 11, 12, 32, 43, 54, 45, 2, 3, 4, 5, 4}
	//size := 20

	data := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 6, 3, 2, 3, 4, 5, 10, 11, 12, 32, 43, 54, 45, 2, 3, 4, 5, 4}
	//arr := make([]int, 31)
	//for i := range arr {
	//	arr[i] = i + 1
	//}
	size := 10

	b.Run("GenerateAllCombinations", func(b *testing.B) {
		b.ResetTimer()
		for range b.N {
			com.GenerateAllCombinations(data, size)
		}
	})

	b.Run("yield (readonly)", func(b *testing.B) {
		b.ResetTimer()
		for range b.N {
			iter := com.NewCombinatorPushIter(data, size)
			iter.ForEach(func(result []int) {})
		}
	})

	b.Run("yield and clone", func(b *testing.B) {
		b.ResetTimer()
		for range b.N {
			iter := com.NewCombinatorPushIter(data, size)
			iter.ForEach(func(result []int) {
				slices.Clone(result)
			})
		}
	})

	b.Run("pull based (readonly)", func(b *testing.B) {
		b.ResetTimer()
		for range b.N {
			iter := com.NewCombinatorPullIter(data, size)
			for iter.Next() {
				iter.Value()
			}
		}
	})

	b.Run("pull based (read+clone)", func(b *testing.B) {
		b.ResetTimer()
		for range b.N {
			iter := com.NewCombinatorPullIter(data, size)
			for iter.Next() {
				slices.Clone(iter.Value())
			}
		}
	})

	b.Run("pull based (copy)", func(b *testing.B) {
		b.ResetTimer()
		for range b.N {
			iter := com.NewCombinatorPullIter(data, size)
			for iter.Next() {
				iter.CopyValue()
			}
		}
	})
}
