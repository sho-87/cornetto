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
