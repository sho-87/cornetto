package main

import (
	"fmt"
)

type Node interface {
	Print()
}

type Scalar struct {
	data         []float64
	shape        []int
	dims         int
	requiresGrad bool
}

func (s *Scalar) Print() {
	fmt.Printf("%+v\n", s.data)
}

type Vector struct {
	data         []float64
	shape        []int
	dims         int
	requiresGrad bool
}

func (v *Vector) Print() {
	fmt.Printf("%+v\n", v.data)
}

type Matrix struct {
	data         []float64
	shape        []int
	dims         int
	requiresGrad bool
}

func (m *Matrix) Print() {
	p := VecToMat(m.data, m.shape)
	maxWidth := findMaxWidth(p)

	for _, row := range p {
		fmt.Printf("%*v\n", maxWidth, row)
	}
}

type Tensor struct {
	data         []float64
	shape        []int
	dims         int
	requiresGrad bool
}

func (t *Tensor) Print() {
	p := VecToTensor(t.data, t.shape)
	maxWidth := findMaxWidth(p)

	for _, mat := range p {
		fmt.Println("{")
		for _, row := range mat {
			fmt.Printf(" %*v\n", maxWidth, row)
		}
		fmt.Println("}")
	}
}
