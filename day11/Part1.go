package main

import (
	"adventcodingchallenge_2021/utility"
	"fmt"
)

type Part1 struct {
	steps int
	data  [][]int
}

func Parse(data []string) [][]int {

	return utility.ParseRowsToInts(data)
}

func (alg *Part1) Process(data []string) (error, interface{}) {
	alg.data = Parse(data)

	numberOfFlashes := 0
	for x := 0; x <= alg.steps; x++ {

		fmt.Printf("Step %d\n", int(x))
		utility.PrettyPrintInts(alg.data)
		numberOfFlashes += alg.ProcessStep()
		fmt.Printf("\n\n\n\n")
	}

	return nil, numberOfFlashes
}

func (alg *Part1) ProcessStep() int {
	numberOfFlashes := 0

	/*
		phase 1
	*/
	for y, _ := range alg.data {
		for x, _ := range alg.data[y] {
			alg.data[y][x]++
		}
	}

	for true {

		ok, flashesThisIteration := alg.ProcessFlashes()
		numberOfFlashes += flashesThisIteration

		if !ok {
			break
		}

	}

	return numberOfFlashes
}

func (alg *Part1) ProcessFlashes() (bool, int) {
	numberOfFlashes := 0

	coordinatesEligibleToFlash := alg.FindReadyToFlash()
	if coordinatesEligibleToFlash == nil || len(coordinatesEligibleToFlash) == 0 {
		return false, numberOfFlashes
	}

	for _, aCoordinate := range coordinatesEligibleToFlash {
		/*
			coordinate energy is set to 0
		*/
		alg.data[aCoordinate.Y][aCoordinate.X] = 0
		numberOfFlashes++
	}

	/*
		for each coordinate, increase energy levels in all directions; distance 1,
	*/
	for _, aCoordinate := range coordinatesEligibleToFlash {
		relativeCoordinates := utility.PositionsRelativeTo(aCoordinate.Y, aCoordinate.X, len(alg.data), len(alg.data[0]))
		for _, aRelativeCoordinate := range relativeCoordinates {
			currentEnergyLevel := alg.data[aRelativeCoordinate.Y][aRelativeCoordinate.X]
			if currentEnergyLevel > 0 {
				alg.data[aRelativeCoordinate.Y][aRelativeCoordinate.X]++
			}

		}
	}

	return true, numberOfFlashes
}

func (alg *Part1) FindReadyToFlash() []*utility.Coordinate {

	flashEligible := make([]*utility.Coordinate, 0)
	for y, _ := range alg.data {
		for x, _ := range alg.data[y] {
			energyLevel := alg.data[y][x]
			if energyLevel != 0 && energyLevel > 9 {

				flashEligible = append(flashEligible, &utility.Coordinate{
					X: x,
					Y: y,
				})
			}

		}
	}

	return flashEligible

}
