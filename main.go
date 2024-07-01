package main

import (
	"fmt"

	"github.com/sho-87/cornetto/data"
)

func main() {
	a, _ := data.Create(5.0)
	fmt.Printf("%+v\n", a)
	a.Print()

	b, _ := data.Create([]float64{5.0, 6.0})
	fmt.Printf("%+v\n", b)
	b.Print()

	c, _ := data.Create([][]float64{{5.0, 6.0, 6.0}, {7.0, 8.0, 8.0}})
	fmt.Printf("%+v\n", c)
	c.Print()

	d, _ := data.Create(
		[][][]float64{
			{
				{5.0, 6.0, 6.0, 7.0},
				{7.0, 8.0, 8.0, 7.0},
				{7.0, 8.0, 8.0, 7.0},
			},
			{
				{9.0, 10.0, 10.0, 7.0},
				{11.0, 12.0, 12.0, 7.0},
				{11.0, 12.0, 12.0, 7.0},
			},
		},
	)
	fmt.Printf("%+v\n", d)
	d.Print()
}
