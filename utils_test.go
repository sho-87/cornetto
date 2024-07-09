package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindMaxWidth(t *testing.T) {
	tests := []struct {
		input    interface{}
		name     string
		expected int
	}{
		{
			name:     "Vector",
			input:    []float64{1.1, 22.22, 333.333},
			expected: 7,
		},
		{
			name:     "Matrix",
			input:    [][]float64{{1.1, 22.22}, {333.333, 4444.4444}},
			expected: 9,
		},
		{
			name: "Tensor",
			input: [][][]float64{
				{{1.1, 22.22}, {333.333, 4444.4444}},
				{{55555.55555, 666666.666666}, {7777777.7777777, 88888888.88888888}},
			},
			expected: 17,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := findMaxWidth(tt.input)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func BenchmarkFindMaxWidth(b *testing.B) {
	for i := 0; i < b.N; i++ {
		findMaxWidth([][]float64{{1.1, 22.22}, {333.333, 4444.4444}})
	}
}
