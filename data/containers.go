package data

import "fmt"

type Container interface {
	Print()
}

type Scalar struct {
	shape        []int
	data         []float64
	dimension    int
	requiresGrad bool
}

func (s *Scalar) Set(value float64) {
	s.data[0] = value
}

func (s *Scalar) Print() {
	fmt.Printf("%+v\n", s.data)
}

type Vector struct {
	shape        []int
	data         []float64
	dimension    int
	requiresGrad bool
}

func (v *Vector) Set(index int, value float64) {
	v.data[index] = value
}

func (v *Vector) Print() {
	fmt.Printf("%+v\n", v.data)
}

type Matrix struct {
	shape        []int
	data         []float64
	dimension    int
	requiresGrad bool
}

func (m *Matrix) Print() {
	unflattened := m.Unflatten(m.data)
	for _, row := range unflattened {
		fmt.Printf("%+v\n", row)
	}
}

func (m *Matrix) Flatten(data [][]float64) []float64 {
	result := make([]float64, 0)
	for _, row := range data {
		result = append(result, row...)
	}
	return result
}

func (m *Matrix) Unflatten(data []float64) [][]float64 {
	result := make([][]float64, 0)
	for i := 0; i < len(data); i += m.shape[1] {
		result = append(result, data[i:i+m.shape[1]])
	}
	return result
}

type Tensor struct {
	shape        []int
	data         []float64
	dimension    int
	requiresGrad bool
}

func (t *Tensor) Print() {
	unflattened := t.Unflatten(t.data)
	for _, row := range unflattened {
		fmt.Println("{")
		for _, col := range row {
			fmt.Printf(" %+v\n", col)
		}
		fmt.Println("},")
	}
}

func (t *Tensor) Flatten(data [][][]float64) []float64 {
	result := make([]float64, 0)
	for _, row := range data {
		for _, col := range row {
			result = append(result, col...)
		}
	}
	return result
}

func (t *Tensor) Unflatten(data []float64) [][][]float64 {
	result := make([][][]float64, 0)
	for i := 0; i < len(data); i += t.shape[1] * t.shape[2] {
		row := make([][]float64, 0)
		for j := 0; j < t.shape[1]; j++ {
			row = append(row, data[i+j*t.shape[2]:(i+(j+1)*t.shape[2])])
		}
		result = append(result, row)
	}
	return result
}
