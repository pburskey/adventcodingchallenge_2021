package main

import (
	"strconv"
)

type Part2 struct {
}

func (alg *Part2) Process(data []string) (error, interface{}) {
	var processed []string
	length := len(data)
	for i, _ := range data {
		if (i + 2) < length {
			a, _ := strconv.Atoi(data[i])
			b, _ := strconv.Atoi(data[i+1])
			c, _ := strconv.Atoi(data[i+2])
			sum := a + b + c
			processed = append(processed, strconv.Itoa(sum))
		}
	}

	if len(processed) > 0 {
		part1 := &Part1{}
		part1.Process(processed)
		return nil, part1.increase
	}

	return nil, 0
}
