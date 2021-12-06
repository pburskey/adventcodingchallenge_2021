package main

import (
	"adventcodingchallenge_2021/utility"
	"strings"
)

type Part1 struct {
}

type Coordinate struct {
	x int
	y int
}

type CoordinateRange struct {
	start *Coordinate
	end   *Coordinate
}

type Field struct {
	coordinates []*Coordinate
}

func (c *CoordinateRange) determineCoordinatesInRange() []*Coordinate {
	coordinates := make([]*Coordinate, 0)
	if c.start.x == c.end.x {
		x := c.start.x
		//diff := int(math.Abs(float64(c.start.y - c.end.y)))
		numbers := []int{c.start.y, c.end.y}
		least, max := utility.LeastAndMax(numbers)
		for y := least; y <= max; y++ {
			aCoordinate := &Coordinate{
				x: x,
				y: y,
			}
			coordinates = append(coordinates, aCoordinate)
		}

	} else if c.start.y == c.end.y {

		y := c.start.y
		//diff := int(math.Abs(float64(c.start.x - c.end.x)))
		numbers := []int{c.start.x, c.end.x}
		least, max := utility.LeastAndMax(numbers)
		for x := least; x <= max; x++ {
			aCoordinate := &Coordinate{
				x: x,
				y: y,
			}
			coordinates = append(coordinates, aCoordinate)
		}

	} else {
		xNumbers := utility.NumbersBetween(c.start.x, c.end.x)
		yNumbers := utility.NumbersBetween(c.start.y, c.end.y)

		for j := 0; j < len(xNumbers); j++ {
			aCoordinate := &Coordinate{
				x: xNumbers[j],
				y: yNumbers[j],
			}
			coordinates = append(coordinates, aCoordinate)
		}
	}

	return coordinates
}

func Parse(data []string) []*CoordinateRange {

	coordinateRanges := make([]*CoordinateRange, 0)

	for _, row := range data {
		aSlice := strings.Split(row, " -> ")
		startString := aSlice[0]
		endString := aSlice[1]

		aSlice = strings.Split(startString, ",")
		aStartCoordinate := &Coordinate{
			x: utility.StringToInt(aSlice[0]),
			y: utility.StringToInt(aSlice[1]),
		}
		aSlice = strings.Split(endString, ",")
		anEndCoordinate := &Coordinate{
			x: utility.StringToInt(aSlice[0]),
			y: utility.StringToInt(aSlice[1]),
		}

		aCoordinateRange := &CoordinateRange{
			start: aStartCoordinate,
			end:   anEndCoordinate,
		}
		coordinateRanges = append(coordinateRanges, aCoordinateRange)
	}
	return coordinateRanges
}

func (alg *Part1) Process(data []string) (error, interface{}) {

	ranges := Parse(data)
	field := &Field{
		coordinates: make([]*Coordinate, 0),
	}
	//coordinates := make([]*Coordinate, 0)
	for _, aRange := range ranges {
		aSetOfCoordinates := aRange.determineCoordinatesInRange()
		if aSetOfCoordinates != nil {
			for _, aCoordinate := range aSetOfCoordinates {
				field.coordinates = append(field.coordinates, aCoordinate)
			}
		}
	}

	tempField := make([][]int, 999)
	for i := range tempField {
		tempField[i] = make([]int, 999)
	}
	for _, aCoordinate := range field.coordinates {
		tempField[aCoordinate.y][aCoordinate.x]++
	}

	overlaps := 0
	for _, row := range tempField {
		for _, column := range row {
			if column >= 2 {
				overlaps++
			}
		}
	}

	return nil, overlaps
}
