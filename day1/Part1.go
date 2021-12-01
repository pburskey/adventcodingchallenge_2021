package main

import (
	"strconv"
)

type Part1 struct {
	increase int
	decrease int
	previous int
	match    int
}

func (alg *Part1) Process(data []string) (error, interface{}) {

	for i, aRow := range data {
		//fmt.Println(fmt.Sprintf("Row: %d ...... %s", i, aRow))
		aNumber, _ := strconv.Atoi(aRow)
		if i > 0 {
			if aNumber > alg.previous {
				alg.increase++
			} else if aNumber < alg.previous {
				alg.decrease++
			} else if aNumber == alg.previous {
				alg.match++
			}
		}
		alg.previous = aNumber

	}
	return nil, alg.increase
}
