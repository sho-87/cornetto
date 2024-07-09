package main

import "strconv"

// Finds the max element width of a Matrix or Tensor
func findMaxWidth(data interface{}) int {
	maxWidth := 0
	switch v := data.(type) {
	case []float64:
		for _, el := range v {
			width := len(strconv.FormatFloat(el, 'f', -1, 64))
			if width > maxWidth {
				maxWidth = width
			}
		}
	case [][]float64:
		for _, row := range v {
			width := findMaxWidth(row)
			if width > maxWidth {
				maxWidth = width
			}
		}
	case [][][]float64:
		for _, mat := range v {
			width := findMaxWidth(mat)
			if width > maxWidth {
				maxWidth = width
			}
		}
	}
	return maxWidth
}
