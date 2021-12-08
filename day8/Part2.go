package main

import (
	"adventcodingchallenge_2021/utility"
	"sort"
	"strings"
)

type Part2 struct {
}

func (alg *Part2) Process(data []string) (error, interface{}) {
	part1 := &Part1{}

	input, _, crypto := part1.Decypher(data)

	signals := make([]int, 0)
	for _, anEntry := range input {
		values := make([]int, 0)
		for _, anOutputValue := range anEntry.outputValues {

			cypherAsArrayOfStrings := strings.Split(anOutputValue, "")
			sort.Strings(cypherAsArrayOfStrings)
			anOutputValue = strings.Join(cypherAsArrayOfStrings, "")

			displaySegment := crypto[anOutputValue]
			if displaySegment != nil {
				values = append(values, displaySegment.value)
			}
		}

		signals = append(signals, utility.SliceToInt(values))
	}

	count := 0
	for _, aValue := range signals {
		count += aValue
	}

	return nil, count

}
