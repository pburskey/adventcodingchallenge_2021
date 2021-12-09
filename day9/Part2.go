package main

import (
	"adventcodingchallenge_2021/utility"
	"math"
	"sort"
)

type Part2 struct {
}

func (alg *Part2) Process(data []string) (error, interface{}) {
	input := Parse(data)
	measuredLowPoints := FindLowPoints(input)

	basins := make([][]*utility.Coordinate, 0)
	for _, aLowPointCoordinate := range measuredLowPoints {

		search := &SearchForBasins{
			input:            input,
			basinCoordinates: make([]*utility.Coordinate, 0),
			visited:          make([]*utility.Coordinate, 0),
		}

		search.search(aLowPointCoordinate)
		basins = append(basins, search.basinCoordinates)

	}
	sort.Sort(ByLength(basins))
	count := 0
	result := 1
	for x := len(basins) - 1; count < 3 && x >= 0; x-- {
		count++
		length := len(basins[x])
		result *= length

	}

	return nil, result
}

type ByLength [][]*utility.Coordinate

func (a ByLength) Len() int {
	return len(a)
}

func (a ByLength) Less(i, j int) bool {
	return len(a[i]) < len(a[j])
}

func (a ByLength) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

type SearchForBasins struct {
	input            [][]int
	basinCoordinates []*utility.Coordinate
	visited          []*utility.Coordinate
}

func (s *SearchForBasins) search(aLowPointCoordinate *utility.Coordinate) {

	if exists(aLowPointCoordinate.X, aLowPointCoordinate.Y, s.visited) {
		return
	}
	s.visited = append(s.visited, aLowPointCoordinate)

	aLowPointValue := s.input[aLowPointCoordinate.Y][aLowPointCoordinate.X]
	if aLowPointValue != 9 {

		if !exists(aLowPointCoordinate.X, aLowPointCoordinate.Y, s.basinCoordinates) {
			s.basinCoordinates = append(s.basinCoordinates, aLowPointCoordinate)
		}

		lengthOfY := len(s.input)
		lengthOfX := len(s.input[0])

		adjacentLocations := SimplePositionsRelativeTo(aLowPointCoordinate.X, aLowPointCoordinate.Y, lengthOfY, lengthOfX)

		for _, aLocationCoordinate := range adjacentLocations {
			aCandidateHeight := s.input[aLocationCoordinate.Y][aLocationCoordinate.X]
			if aCandidateHeight < 9 {
				if !exists(aLocationCoordinate.X, aLocationCoordinate.Y, s.basinCoordinates) {
					s.basinCoordinates = append(s.basinCoordinates, &utility.Coordinate{
						X: aLocationCoordinate.X,
						Y: aLocationCoordinate.Y,
					})
				}
			}
		}

		if s.basinCoordinates != nil && len(s.basinCoordinates) > 0 {
			for _, aBasinCoordinate := range adjacentLocations {
				s.search(aBasinCoordinate)

			}
		}

	}

}

func exists(x, y int, array []*utility.Coordinate) bool {
	found := false
	for i := 0; !found && i < len(array); i++ {
		aCoordinate := array[i]
		found = (aCoordinate.X == x && aCoordinate.Y == y)
	}

	return found
}

func FindLowPoints(input [][]int) []*utility.Coordinate {

	measuredLowPoints := make([]*utility.Coordinate, 0)
	for y, _ := range input {
		for x, _ := range input[y] {
			height := input[y][x]
			//fmt.Printf("Y: %d X: %d Height: %d\n", y, x, height)
			adjacentLocations := SimplePositionsRelativeTo(x, y, len(input), len(input[y]))

			lowPoint := math.MaxInt32
			for _, aLocationCoordinate := range adjacentLocations {
				aCandidateHeight := input[aLocationCoordinate.Y][aLocationCoordinate.X]
				if aCandidateHeight <= lowPoint {
					lowPoint = aCandidateHeight
				}
			}
			if height < lowPoint {
				measuredLowPoints = append(measuredLowPoints, &utility.Coordinate{
					X: x,
					Y: y,
				})
			}

		}
	}

	return measuredLowPoints

}
