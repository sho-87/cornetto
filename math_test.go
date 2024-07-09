package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestProduct(t *testing.T) {
	tests := []struct {
		input    interface{}
		expected interface{}
	}{
		{input: []int{}, expected: 0},
		{input: []int{1, 2, 3, 4}, expected: 24},
		{input: []int{4, 3, 2, 1}, expected: 24},
		{input: []float64{1.1, 2.2, 3.3}, expected: 7.986},
	}

	for _, test := range tests {
		switch v := test.input.(type) {
		case []int:
			result := Product(v)
			if result != test.expected.(int) {
				t.Errorf("Product(%v) = %v; want %v", v, result, test.expected)
			}
		case []float64:
			result := Product(v)
			assert.InEpsilon(t, test.expected.(float64), result, 0.01)
		}
	}
}

func TestSum(t *testing.T) {
	tests := []struct {
		input    interface{}
		expected interface{}
	}{
		{input: []int{1, 2, 3, 4}, expected: 10},
		{input: []float64{1.1, 2.2, 3.3}, expected: 6.6},
		{input: []int{}, expected: 0},
	}

	for _, test := range tests {
		switch v := test.input.(type) {
		case []int:
			result := Sum(v)
			if result != test.expected.(int) {
				t.Errorf("Sum(%v) = %v; want %v", v, result, test.expected)
			}
		case []float64:
			result := Sum(v)
			assert.InEpsilon(t, test.expected.(float64), result, 0.01)
		}
	}
}

func BenchmarkProduct(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Product([]float64{333.333, 4444.4444, 10.2, 1.8})
	}
}

func BenchmarkSum(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Sum([]float64{333.333, 4444.4444, 10.2, 1.8})
	}
}
