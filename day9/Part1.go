package main

import (
	"strconv"
	"strings"
)

type Part1 struct {
}

func Parse(data []string) [][]int {

	array := make([][]int, len(data))

	for y, row := range data {
		anArrayOfNumbers := make([]int, len(row))
		strings := strings.Split(row, "")

		for x, aString := range strings {
			aNumber, _ := strconv.Atoi(aString)
			anArrayOfNumbers[x] = aNumber
		}

		array[y] = anArrayOfNumbers
		//fmt.Printf("%s\n", row)
	}
	return array
}

func (alg *Part1) Process(data []string) (error, interface{}) {
	input := Parse(data)
	measuredLowPoints := FindLowPoints(input)
	count := 0

	for _, aCoordinate := range measuredLowPoints {
		aValue := input[aCoordinate.Y][aCoordinate.X]
		count += (1 + aValue)

	}

	return nil, count
}
