package com

import "golang.org/x/exp/constraints"

type Number interface {
	int
	int64
}

// PermutationsFormula	n!
func PermutationsFormula[T constraints.Integer](n, k T) float64 {
	return Factorial(n)
}

// ArrangementsFormula	n! / (n-k)!
func ArrangementsFormula[T constraints.Integer](n, k T) float64 {
	return PartialFactorial(n-k+1, n)
}

// CombinationsFormula	n! / k!*(n-k)!
func CombinationsFormula[T constraints.Integer](n, k T) float64 {
	if k > n-k {
		return PartialFactorial(n-k+1, n) / Factorial(k)
	}
	return PartialFactorial(k, n) / Factorial(n-k+1)
}

func Factorial[T constraints.Integer](n T) float64 {
	result := 1.0

	if n <= 1 {
		return result
	}

	for i := T(2); i <= n; i++ {
		result *= float64(i)
	}

	return result
}

func PartialFactorial[T constraints.Integer](start, end T) float64 {
	if end < start {
		panic("start cant be less than end")
	}

	if end < 0 {
		panic("end cant be less than 0")
	}

	result := 1.0
	if end <= 1 {
		return result
	}

	for i := start; i <= end; i++ {
		result *= float64(i)
	}

	return result
}
