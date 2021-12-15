package main

import (
	"adventcodingchallenge_2021/utility"
	"fmt"
)

type Part2 struct {
	steps int
}

func (alg *Part2) Process(data []string) (error, interface{}) {

	part1 := &Part1{data: Parse(data)}
	numberOfFlashesToLookFor := (len(part1.data) * len(part1.data[0]))

	stepToFind := 0

	utility.PrettyPrintInts(part1.data)
	fmt.Printf("\n\n\n\n")
	for x := 1; stepToFind == 0; x++ {

		numberOfFlashes := part1.ProcessStep()
		fmt.Printf("After Step %d\n", int(x))
		utility.PrettyPrintInts(part1.data)
		fmt.Printf("\n\n\n\n")
		if numberOfFlashes == numberOfFlashesToLookFor {
			stepToFind = x
			break
		}
	}

	return nil, stepToFind
}
