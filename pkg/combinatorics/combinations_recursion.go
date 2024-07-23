package collections

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

		end := len(arr) + len(result) + 1 - size
		for j := start; j < end; j++ {
			result = append(result, arr[j])
			backtracking(j + 1)
			result = result[:len(result)-1]
		}
	}

	backtracking(0)
	return results
}
