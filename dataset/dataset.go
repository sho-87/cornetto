package dataset

import (
	"encoding/csv"
	"os"
	"strconv"
)

type Dataset struct {
	Path    string
	Headers []string
	Data    [][]float64
}

func ReadCSV(path string, includeHeaders bool) Dataset {
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}

	r := csv.NewReader(file)
	records, err := r.ReadAll()
	if err != nil {
		panic(err)
	}

	var data [][]float64
	var headers []string

	startRow := 0
	if includeHeaders {
		headers = records[0]
		startRow = 1
	}
	removeColumns := []int{}

	for _, record := range records[startRow:] {
		var row []float64
		for colNum, field := range record {
			value, err := strconv.ParseFloat(field, 64)
			if err != nil {
				if !InSlice(colNum, removeColumns) {
					removeColumns = append(removeColumns, colNum)
				}
			}
			row = append(row, value)
		}
		row = DropIndex(row, removeColumns)
		data = append(data, row)
	}

	dataset := Dataset{
		Path:    path,
		Data:    data,
		Headers: []string{},
	}

	if includeHeaders {
		headers = DropIndex(headers, removeColumns)
		dataset.Headers = headers
	}
	return dataset
}

func InSlice[T comparable](needle T, haystack []T) bool {
	for _, v := range haystack {
		if v == needle {
			return true
		}
	}
	return false
}

func DropIndex[T any](arr []T, index []int) []T {
	var result []T
	for i, v := range arr {
		if !InSlice(i, index) {
			result = append(result, v)
		}
	}
	return result
}
