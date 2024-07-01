package data

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreate(t *testing.T) {
	t.Run("Scalar", func(t *testing.T) {
		s, err := Create(5.0)
		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}

		if s, ok := s.(*Scalar); !ok {
			t.Errorf("Expected Scalar, got %T", s)
		}
	})

	t.Run("Vector", func(t *testing.T) {
		s, err := Create([]float64{5.0, 6.0})
		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}

		if s, ok := s.(*Vector); !ok {
			t.Errorf("Expected vector, got %T", s)
		}
	})

	t.Run("Matrix", func(t *testing.T) {
		s, err := Create([][]float64{
			{5.0, 6.0},
			{7.0, 8.0},
		})
		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}

		if s, ok := s.(*Matrix); !ok {
			t.Errorf("Expected matrix, got %T", s)
		}
	})

	t.Run("Tensor", func(t *testing.T) {
		s, err := Create([][][]float64{
			{
				{5.0, 6.0},
				{7.0, 8.0},
			},
			{
				{9.0, 10.0},
				{11.0, 12.0},
			},
		})
		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}

		if s, ok := s.(*Tensor); !ok {
			t.Errorf("Expected tensor, got %T", s)
		}
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
