package main

import (
	"errors"
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

// func Zeros(shape []int) Node {
// 	return Node{shape: shape, dims: len(shape)}
// }
//
// func Ones(shape []int) Node {
// 	return Node{shape: shape, dims: len(shape)}
// }
//
// func Rand(shape []int) Node {
// 	return Node{shape: shape, dims: len(shape)}
// }
