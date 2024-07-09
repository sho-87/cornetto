package main

func MatToVec(data [][]float64) []float64 {
	result := make([]float64, 0)
	for _, row := range data {
		result = append(result, row...)
	}
	return result
}

func VecToMat(data []float64, shape []int) [][]float64 {
	if len(shape) != 2 || len(data) != Product(shape) {
		panic("Invalid shape for matrix")
	}

	result := make([][]float64, 0)
	for i := 0; i < len(data); i += shape[1] {
		result = append(result, data[i:i+shape[1]])
	}
	return result
}

func TensorToVec(data [][][]float64) []float64 {
	result := make([]float64, 0)
	for _, row := range data {
		for _, col := range row {
			result = append(result, col...)
		}
	}
	return result
}

func VecToTensor(data []float64, shape []int) [][][]float64 {
	if len(shape) != 3 || len(data) != Product(shape) {
		panic("Invalid shape for tensor")
	}

	result := make([][][]float64, 0)
	for i := 0; i < len(data); i += shape[1] * shape[2] {
		row := make([][]float64, 0)
		for j := 0; j < shape[1]; j++ {
			row = append(row, data[i+j*shape[2]:(i+(j+1)*shape[2])])
		}
		result = append(result, row)
	}
	return result
}
