package main

import (
	"testing"
)

func BenchmarkMatToVec(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = MatToVec([][]float64{
			{1.0, 2.0, 3.0},
			{4.0, 5.0, 6.0},
		})
	}
}

func BenchmarkVecToMat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = VecToMat([]float64{1.0, 2.0, 3.0, 4.0, 5.0, 6.0}, []int{2, 3})
	}
}

func BenchmarkTensorToVec(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = TensorToVec([][][]float64{
			{
				{1.0, 2.0, 3.0},
				{4.0, 5.0, 6.0},
			},
			{
				{7.0, 8.0, 9.0},
				{10.0, 11.0, 12.0},
			},
		})
	}
}

func BenchmarkVecToTensor(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = VecToTensor(
			[]float64{1.0, 2.0, 3.0, 4.0, 5.0, 6.0, 7.0, 8.0, 9.0, 10.0, 11.0, 12.0},
			[]int{2, 3, 2},
		)
	}
}
