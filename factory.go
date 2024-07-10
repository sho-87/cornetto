package main

import (
	"errors"
	"math/rand/v2"
)

// Create a new data node. Purpose is to provide a singular entry point to node creation.
// The type of the data is inferred from the dimensions of the input data.
// Data can be a Scalar (float64), Vector ([]float64), Matrix ([][]float64) or Tensor ([][][]float64).
// Dimension order is [depth?][row][column]float64.
func Create(data interface{}) (Node, error) {
	switch data := data.(type) {
	case float64:
		return &Scalar{data: []float64{data}, dims: 0, shape: []int{1}}, nil
	case []float64:
		return &Vector{data: data, dims: 1, shape: []int{len(data)}}, nil
	case [][]float64:
		// Assert that sub dimensions are equal
		row0 := data[0]
		for _, row := range data {
			if len(row) != len(row0) {
				panic("Column dimensions must be equal")
			}
		}
		m := &Matrix{data: MatToVec(data), dims: 2, shape: []int{len(data), len(data[0])}}
		return m, nil
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
		t := &Tensor{
			data:  TensorToVec(data),
			dims:  3,
			shape: []int{len(data), len(data[0]), len(data[0][0])},
		}
		return t, nil
	default:
		return nil, errors.New("unknown node type")
	}
}

func CreateNodeWithShape(shape []int, data []float64) Node {
	switch len(shape) {
	case 1:
		return &Vector{shape: shape, dims: 1, data: data}
	case 2:
		return &Matrix{shape: shape, dims: 2, data: data}
	case 3:
		return &Tensor{shape: shape, dims: 3, data: data}
	}
	panic("Invalid shape for node")
}

func Zeros(shape []int) Node {
	data := make([]float64, Product(shape))
	return CreateNodeWithShape(shape, data)
}

func Ones(shape []int) Node {
	data := make([]float64, Product(shape))
	for i := 0; i < Product(shape); i++ {
		data[i] = 1
	}
	return CreateNodeWithShape(shape, data)
}

func Rand(shape []int) Node {
	data := make([]float64, Product(shape))
	for i := 0; i < Product(shape); i++ {
		data[i] = rand.Float64()
	}
	return CreateNodeWithShape(shape, data)
}

func RandNormal(shape []int) Node {
	data := make([]float64, Product(shape))
	for i := 0; i < Product(shape); i++ {
		data[i] = rand.NormFloat64()
	}
	return CreateNodeWithShape(shape, data)
}

func Eye(shape int) Node {
	data := make([]float64, shape*shape)
	for i := 0; i < shape*shape; i++ {
		data[i] = 0
	}
	for i := 0; i < shape; i++ {
		data[i*shape+i] = 1
	}
	return &Matrix{shape: []int{shape, shape}, dims: 2, data: data}
}
