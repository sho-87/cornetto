package main

func Product[T float64 | int](arr []T) T {
	if len(arr) == 0 {
		return 0
	}

	result := arr[0]
	for _, v := range arr[1:] {
		result *= v
	}
	return result
}

func Sum[T float64 | int](arr []T) T {
	var result T
	for _, v := range arr {
		result += v
	}
	return result
}
