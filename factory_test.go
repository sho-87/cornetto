package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreate(t *testing.T) {
	t.Run("Scalar", func(t *testing.T) {
		s, err := Create(5.0)
		assert.NoError(t, err, "Expected no error")

		_, ok := s.(*Scalar)
		assert.True(t, ok, "Expected type *Scalar, got %T", s)
	})

	t.Run("Vector", func(t *testing.T) {
		v, err := Create([]float64{5.0, 6.0})
		assert.NoError(t, err, "Expected no error")

		_, ok := v.(*Vector)
		assert.True(t, ok, "Expected type *Vector, got %T", v)
	})

	t.Run("Matrix", func(t *testing.T) {
		m, err := Create([][]float64{
			{5.0, 6.0},
			{7.0, 8.0},
		})
		assert.NoError(t, err, "Expected no error")

		_, ok := m.(*Matrix)
		assert.True(t, ok, "Expected type *Matrix, got %T", m)
	})

	t.Run("Tensor", func(t *testing.T) {
		a, err := Create([][][]float64{
			{
				{5.0, 6.0},
				{7.0, 8.0},
			},
			{
				{9.0, 10.0},
				{11.0, 12.0},
			},
		})
		assert.NoError(t, err, "Expected no error")

		_, ok := a.(*Tensor)
		assert.True(t, ok, "Expected type *Tensor, got %T", a)
	})
}

func TestCreateWrongDimensions(t *testing.T) {
	t.Run("Matrix", func(t *testing.T) {
		assert.Panics(t, func() {
			_, _ = Create([][]float64{
				{5.0, 6.0},
				{7.0, 8.0, 8.0},
			})
		}, "Expected panic for wrong dimensions")
	})

	t.Run("Tensor", func(t *testing.T) {
		assert.Panics(t, func() {
			_, _ = Create([][][]float64{
				{
					{5.0, 6.0},
					{7.0, 8.0, 8.0},
				},
				{
					{9.0, 10.0},
					{11.0, 12.0},
				},
			})
		}, "Expected panic for wrong dimensions")
	})
}

func TestZeros(t *testing.T) {
	t.Run("1D Vector", func(t *testing.T) {
		shape := []int{5}
		node := Zeros(shape)
		vector, ok := node.(*Vector)
		assert.True(t, ok)
		assert.Equal(t, shape, vector.shape)
		assert.Equal(t, 1, vector.dims)
		assert.Equal(t, make([]float64, 5), vector.data)
	})

	t.Run("2D Matrix", func(t *testing.T) {
		shape := []int{3, 4}
		node := Zeros(shape)
		matrix, ok := node.(*Matrix)
		assert.True(t, ok)
		assert.Equal(t, shape, matrix.shape)
		assert.Equal(t, 2, matrix.dims)
		assert.Equal(t, make([]float64, 12), matrix.data)
	})

	t.Run("3D Tensor", func(t *testing.T) {
		shape := []int{2, 3, 4}
		node := Zeros(shape)
		tensor, ok := node.(*Tensor)
		assert.True(t, ok)
		assert.Equal(t, shape, tensor.shape)
		assert.Equal(t, 3, tensor.dims)
		assert.Equal(t, make([]float64, 24), tensor.data)
	})

	t.Run("Invalid Shape", func(t *testing.T) {
		shape := []int{}
		assert.Panics(t, func() { Zeros(shape) }, "The code did not panic")
	})
}

func TestOnes(t *testing.T) {
	t.Run("1D Vector", func(t *testing.T) {
		shape := []int{5}
		node := Ones(shape)
		vector, ok := node.(*Vector)
		assert.True(t, ok)
		assert.Equal(t, shape, vector.shape)
		assert.Equal(t, 1, vector.dims)
		for _, v := range vector.data {
			assert.Equal(t, 1.0, v)
		}
	})

	t.Run("2D Matrix", func(t *testing.T) {
		shape := []int{3, 4}
		node := Ones(shape)
		matrix, ok := node.(*Matrix)
		assert.True(t, ok)
		assert.Equal(t, shape, matrix.shape)
		assert.Equal(t, 2, matrix.dims)
		for _, v := range matrix.data {
			assert.Equal(t, 1.0, v)
		}
	})

	t.Run("3D Tensor", func(t *testing.T) {
		shape := []int{2, 3, 4}
		node := Ones(shape)
		tensor, ok := node.(*Tensor)
		assert.True(t, ok)
		assert.Equal(t, shape, tensor.shape)
		assert.Equal(t, 3, tensor.dims)
		for _, v := range tensor.data {
			assert.Equal(t, 1.0, v)
		}
	})

	t.Run("Invalid Shape", func(t *testing.T) {
		shape := []int{}
		assert.Panics(t, func() { Ones(shape) }, "The code did not panic")
	})
}

func TestRand(t *testing.T) {
	t.Run("1D Vector", func(t *testing.T) {
		shape := []int{5}
		node := Rand(shape)
		vector, ok := node.(*Vector)
		assert.True(t, ok)
		assert.Equal(t, shape, vector.shape)
		assert.Equal(t, 1, vector.dims)
		assert.Equal(t, 5, len(vector.data))
		for _, v := range vector.data {
			assert.True(t, v >= 0.0 && v < 1.0)
		}
	})

	t.Run("2D Matrix", func(t *testing.T) {
		shape := []int{3, 4}
		node := Rand(shape)
		matrix, ok := node.(*Matrix)
		assert.True(t, ok)
		assert.Equal(t, shape, matrix.shape)
		assert.Equal(t, 2, matrix.dims)
		assert.Equal(t, 12, len(matrix.data))
		for _, v := range matrix.data {
			assert.True(t, v >= 0.0 && v < 1.0)
		}
	})

	t.Run("3D Tensor", func(t *testing.T) {
		shape := []int{2, 3, 4}
		node := Rand(shape)
		tensor, ok := node.(*Tensor)
		assert.True(t, ok)
		assert.Equal(t, shape, tensor.shape)
		assert.Equal(t, 3, tensor.dims)
		assert.Equal(t, 24, len(tensor.data))
		for _, v := range tensor.data {
			assert.True(t, v >= 0.0 && v < 1.0)
		}
	})

	t.Run("Invalid Shape", func(t *testing.T) {
		shape := []int{}
		assert.Panics(t, func() { Rand(shape) }, "The code did not panic")
	})
}

func TestRandNormal(t *testing.T) {
	t.Run("1D Vector", func(t *testing.T) {
		shape := []int{5}
		node := RandNormal(shape)
		vector, ok := node.(*Vector)
		assert.True(t, ok)
		assert.Equal(t, shape, vector.shape)
		assert.Equal(t, 1, vector.dims)
		assert.Equal(t, 5, len(vector.data))
		for _, v := range vector.data {
			assert.True(t, v != 0.0)
		}
	})

	t.Run("2D Matrix", func(t *testing.T) {
		shape := []int{3, 4}
		node := RandNormal(shape)
		matrix, ok := node.(*Matrix)
		assert.True(t, ok)
		assert.Equal(t, shape, matrix.shape)
		assert.Equal(t, 2, matrix.dims)
		assert.Equal(t, 12, len(matrix.data))
		for _, v := range matrix.data {
			assert.True(t, v != 0.0)
		}
	})

	t.Run("3D Tensor", func(t *testing.T) {
		shape := []int{2, 3, 4}
		node := RandNormal(shape)
		tensor, ok := node.(*Tensor)
		assert.True(t, ok)
		assert.Equal(t, shape, tensor.shape)
		assert.Equal(t, 3, tensor.dims)
		assert.Equal(t, 24, len(tensor.data))
		for _, v := range tensor.data {
			assert.True(t, v != 0.0)
		}
	})

	t.Run("Invalid Shape", func(t *testing.T) {
		shape := []int{}
		assert.Panics(t, func() { RandNormal(shape) }, "The code did not panic")
	})
}

func TestEye(t *testing.T) {
	t.Run("2x2 Identity Matrix", func(t *testing.T) {
		shape := 2
		node := Eye(shape)
		matrix, ok := node.(*Matrix)
		assert.True(t, ok)
		assert.Equal(t, []int{shape, shape}, matrix.shape)
		assert.Equal(t, 2, matrix.dims)
		expected := []float64{
			1, 0,
			0, 1,
		}
		assert.Equal(t, expected, matrix.data)
	})

	t.Run("3x3 Identity Matrix", func(t *testing.T) {
		shape := 3
		node := Eye(shape)
		matrix, ok := node.(*Matrix)
		assert.True(t, ok)
		assert.Equal(t, []int{shape, shape}, matrix.shape)
		assert.Equal(t, 2, matrix.dims)
		expected := []float64{
			1, 0, 0,
			0, 1, 0,
			0, 0, 1,
		}
		assert.Equal(t, expected, matrix.data)
	})

	t.Run("4x4 Identity Matrix", func(t *testing.T) {
		shape := 4
		node := Eye(shape)
		matrix, ok := node.(*Matrix)
		assert.True(t, ok)
		assert.Equal(t, []int{shape, shape}, matrix.shape)
		assert.Equal(t, 2, matrix.dims)
		expected := []float64{
			1, 0, 0, 0,
			0, 1, 0, 0,
			0, 0, 1, 0,
			0, 0, 0, 1,
		}
		assert.Equal(t, expected, matrix.data)
	})
}
