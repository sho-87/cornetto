package main

func main() {
	a, _ := Create(5.0)
	a.Print()

	b, _ := Create([]float64{5.0, 6.0})
	b.Print()

	c, _ := Create([][]float64{{5.0, 6.0, 61.0}, {7.0, 8.0, 8.0}})
	c.Print()

	d, _ := Create(
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
	d.Print()

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
}
