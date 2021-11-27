package main

import (
	"fmt"
	"strconv"
)

type Part1 struct {
}

func (alg *Part1) Process(data []string) (error, interface{}) {
	sum := 0
	for i, aRow := range data {
		fmt.Println(fmt.Sprintf("Row: %d ...... %s", i, aRow))
		aNumber, _ := strconv.Atoi(aRow)
		sum += aNumber
	}
	return nil, sum
}
