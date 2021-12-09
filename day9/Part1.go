package main

import (
	"adventcodingchallenge_2021/utility"
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

//
//func PositionsRelativeTo(xstart, ystart int) []*utility.Coordinate {
//	coordinates := make([]*utility.Coordinate, 0)
//	for x := (xstart - 1); x < (xstart + 1); x++ {
//		for y := (ystart - 1); y < (ystart + 1); y++ {
//
//			if x >= 0 && y >= 0 && (x != xstart && y != ystart) {
//				coordinate := &utility.Coordinate{
//					X: x,
//					Y: y,
//				}
//				coordinates = append(coordinates, coordinate)
//			}
//
//		}
//	}
//	return coordinates
//}

func SimplePositionsRelativeTo(xstart, ystart int, endOfY int, endOfX int) []*utility.Coordinate {
	coordinates := make([]*utility.Coordinate, 0)
	for x := (xstart - 1); x <= (xstart + 1); x++ {

		if x != xstart {
			if x >= 0 && x < endOfX {
				coordinate := &utility.Coordinate{
					X: x,
					Y: ystart,
				}
				coordinates = append(coordinates, coordinate)
			}
		}

	}

	for y := (ystart - 1); y <= (ystart + 1); y++ {

		if y != ystart {
			if y >= 0 && y < endOfY {
				coordinate := &utility.Coordinate{
					X: xstart,
					Y: y,
				}
				coordinates = append(coordinates, coordinate)
			}
		}

	}

	return coordinates
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
