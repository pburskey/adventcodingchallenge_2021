package main

import (
	"adventcodingchallenge_2021/utility"
	"strings"
)

type Part1 struct {
}

type Field struct {
	coordinates []*utility.Coordinate
}

func Parse(data []string) []*utility.CoordinateRange {

	coordinateRanges := make([]*utility.CoordinateRange, 0)

	for _, row := range data {
		aSlice := strings.Split(row, " -> ")
		startString := aSlice[0]
		endString := aSlice[1]

		aSlice = strings.Split(startString, ",")
		aStartCoordinate := &utility.Coordinate{
			X: utility.StringToInt(aSlice[0]),
			Y: utility.StringToInt(aSlice[1]),
		}
		aSlice = strings.Split(endString, ",")
		anEndCoordinate := &utility.Coordinate{
			X: utility.StringToInt(aSlice[0]),
			Y: utility.StringToInt(aSlice[1]),
		}

		aCoordinateRange := &utility.CoordinateRange{
			Start: aStartCoordinate,
			End:   anEndCoordinate,
		}
		coordinateRanges = append(coordinateRanges, aCoordinateRange)
	}
	return coordinateRanges
}

func (alg *Part1) Process(data []string, diag bool) (error, interface{}) {

	ranges := Parse(data)
	field := &Field{
		coordinates: make([]*utility.Coordinate, 0),
	}
	//coordinates := make([]*Coordinate, 0)
	for _, aRange := range ranges {
		aSetOfCoordinates := aRange.DetermineCoordinatesInRange(diag)
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
		tempField[aCoordinate.Y][aCoordinate.X]++
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
