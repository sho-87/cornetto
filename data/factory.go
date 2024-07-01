package data

import (
	"errors"
)

// Create a new data container. Purpose is to provide a singular entry point to container creation.
// The type of the data is inferred from the dimensions of the input data.
// The data can be a scalar (float64), vector ([]float64), matrix ([][]float64) or tensor ([][][]float64).
// Dimension order is [depth][row][column]float64.
func Create(data interface{}) (Container, error) {
	switch data := data.(type) {
	case float64:
		return &Scalar{data: []float64{data}, dimension: 0, shape: []int{1}}, nil
	case []float64:
		return &Vector{data: data, dimension: 1, shape: []int{len(data)}}, nil
	case [][]float64:
		// Assert that sub dimensions are equal
		row0 := data[0]
		for _, row := range data {
			if len(row) != len(row0) {
				panic("Column dimensions must be equal")
			}
		}
		m := Matrix{dimension: 2, shape: []int{len(data), len(data[0])}}
		m.data = m.Flatten(data)
		return &m, nil
	case [][][]float64:
		// Assert that sub dimensions are equal
		row0 := data[0]
		col0 := row0[0]
		for _, row := range data {
			if len(row) != len(row0) {
				panic("Row dimensions must be equal")
			}

			for _, col := range row {
				if len(col) != len(col0) {
					panic("Column dimensions must be equal")
				}
			}
		}
		t := Tensor{dimension: 3, shape: []int{len(data), len(data[0]), len(data[0][0])}}
		t.data = t.Flatten(data)
		return &t, nil
	default:
		return nil, errors.New("unknown container type")
	}
}

func Zeros(shape []int) Tensor {
	return Tensor{shape: shape, dimension: len(shape)}
}

func Ones(shape []int) Tensor {
	return Tensor{shape: shape, dimension: len(shape)}
}

func Rand(shape []int) Tensor {
	return Tensor{shape: shape, dimension: len(shape)}
}
