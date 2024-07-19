package com

import (
	"testing"
)

func TestFactorial(t *testing.T) {
	type args struct {
		n int
	}
	type testCase struct {
		name string
		args args
		want float64
	}
	tests := []testCase{
		{"", args{n: 0}, 1},
		{"", args{n: 1}, 1},
		{"", args{n: 2}, 2},
		{"", args{n: 10}, 3628800},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Factorial(tt.args.n); got != tt.want {
				t.Errorf("Factorial() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPermutationsFormula(t *testing.T) {
	type args struct {
		n int
	}
	type testCase struct {
		name string
		args args
		want float64
	}
	tests := []testCase{
		{"", args{n: 5}, 120},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Factorial(tt.args.n); got != tt.want {
				t.Errorf("Factorial() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestArrangementsFormula(t *testing.T) {
	type args struct {
		n int
		k int
	}
	type testCase struct {
		name string
		args args
		want float64
	}
	tests := []testCase{
		{"", args{n: 5, k: 2}, 20},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ArrangementsFormula(tt.args.n, tt.args.k); got != tt.want {
				t.Errorf("Factorial() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCombinationsFormula(t *testing.T) {
	type args struct {
		n int
		k int
	}
	type testCase struct {
		name string
		args args
		want float64
	}
	tests := []testCase{
		{"", args{n: 6, k: 4}, 15},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CombinationsFormula(tt.args.n, tt.args.k); got != tt.want {
				t.Errorf("Factorial() = %v, want %v", got, tt.want)
			}
		})
	}
}
