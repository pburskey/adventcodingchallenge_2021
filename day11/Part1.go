package main

import (
	"adventcodingchallenge_2021/utility"
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

	for x := 0; x < alg.steps; x++ {
		alg.ProcessStep()
	}

	result := 0
	return nil, result
}

func (alg *Part1) ProcessStep() {
	/*
		phase 1
	*/
	for y, _ := range alg.data {
		for x, _ := range alg.data[y] {
			alg.data[y][x]++
		}
	}

	coordinatesEligibleToFlash := alg.FindReadyToFlash()
	for coordinatesEligibleToFlash != nil && len(coordinatesEligibleToFlash) > ){
		for _, aCoordinate := range coordinatesEligibleToFlash {
			/*
				coordinate energy is set to 0
			*/
			alg.data[aCoordinate.Y][aCoordinate.X] = 0
		}

		/*
			for each coordinate, increase energy levels in all directions; distance 1,
		*/

	}

}

func (alg *Part1) FindReadyToFlash() []*utility.Coordinate {

	flashEligible := make([]*utility.Coordinate, 0)
	for y, _ := range alg.data {
		for x, _ := range alg.data[y] {
			energyLevel := alg.data[y][x]
			if energyLevel >= 9 {

				flashEligible = append(flashEligible, &utility.Coordinate{
					X: x,
					Y: y,
				})
			}

		}
	}

	return flashEligible

}
