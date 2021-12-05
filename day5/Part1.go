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

func (c *CoordinateRange) determineSimpleCoordinatesInRange() []*Coordinate {
	coordinates := make([]*Coordinate, 0)
	if (c.start.x == c.end.x) || (c.start.y == c.end.y) {
		for x := c.start.x; x <= c.end.x; x++ {
			for y := c.start.y; y <= c.end.y; y++ {
				aCoordinate := &Coordinate{
					x: x,
					y: y,
				}
				coordinates = append(coordinates, aCoordinate)
			}
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
		aSetOfCoordinates := aRange.determineSimpleCoordinatesInRange()
		if aSetOfCoordinates != nil {
			for _, aCoordinate := range aSetOfCoordinates {
				field.coordinates = append(field.coordinates, aCoordinate)
			}
		}
	}

	tempField := make([][]int, 10)
	for i := range tempField {
		tempField[i] = make([]int, 10)
	}
	for _, aCoordinate := range field.coordinates {
		tempField[aCoordinate.x][aCoordinate.y]++
	}

	return nil, nil
}
