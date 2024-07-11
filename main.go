package main

import (
	"fmt"

	"github.com/sho-87/cornetto/dataset"
)

func main() {
	s, _ := Create(5.0)
	s.Print()

	v, _ := Create([]float64{5.0, 6.0})
	v.Print()

	m, _ := Create([][]float64{{5.0, 6.0, 61.0}, {7.0, 8.0, 8.0}})
	m.Print()

	t, _ := Create(
		[][][]float64{
			{
				{5.0, 6.0, 6.0, 7.0},
				{7.0, 8.0, 8.0, 7.0},
				{7.0, 8.0, 8.0, 7.0},
			},
			{
				{9.0, 10.0, 10.0, 7.0},
				{11.0, 12.0, 12.0, 7.0},
				{11.0, 12.0, 121.0, 7.0},
			},
		},
	)
	t.Print()

	zeros := Zeros([]int{2, 3})
	zeros.Print()

	ones := Ones([]int{2, 3})
	ones.Print()

	rand := Rand([]int{2, 3})
	rand.Print()

	randn := RandNormal([]int{2, 3})
	randn.Print()

	eye := Eye(4)
	eye.Print()

	dataset := dataset.ReadCSV("dataset/walmart.csv", true)
	fmt.Println(dataset.Headers)
	data, err := Create(dataset.Data[:5])
	if err != nil {
		panic(err)
	}
	data.Print()
}
