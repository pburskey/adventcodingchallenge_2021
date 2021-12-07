package main

import (
	"adventcodingchallenge_2021/utility"
	"fmt"
	"log"
	"math"
)

type Part2 struct {
	numberOfDays int
}

func (alg *Part2) Process(data []string) (error, interface{}) {
	input := Parse(data)

	costs := make(map[int]int, 0)
	median := utility.Median(input)
	costs[int(median)] = costExpensiveOfAdjustingPositionsTo(input, int(median))
	mean := utility.Mean(input)
	costs[mean] = costExpensiveOfAdjustingPositionsTo(input, mean)

	least, max := utility.LeastAndMax([]int{int(median), mean})
	least = int(math.RoundToEven(float64(least) * float64(.5)))
	max = int(math.RoundToEven(float64(max) * float64(1.5)))

	for x := least; x < max; x++ {
		costs[x] = costExpensiveOfAdjustingPositionsTo(input, x)
	}

	log.Println(fmt.Sprintf("Median: %.2f", median))
	log.Println(fmt.Sprintf("Mean: %d", mean))

	//log.Println(fmt.Sprintf("Cost to mean: %d", costOfAdjustingPositionsTo(input, mean)))
	//log.Println(fmt.Sprintf("Cost to mean: %d", costOfAdjustingPositionsTo(input, int(median))))

	lowValue := math.MaxInt32
	for key, value := range costs {
		log.Println(fmt.Sprintf("Cost to position: %d is: %d", key, value))
		if value < lowValue {
			lowValue = value
		}
	}

	return nil, lowValue
}
