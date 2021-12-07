package main

import (
	"adventcodingchallenge_2021/utility"
	"fmt"
	"log"
	"math"
	"sort"
	"strconv"
	"strings"
)

type Part1 struct {
	numberOfDays int
}

func Parse(data []string) []int {

	array := make([]int, 0)

	for _, row := range data {
		numbersAsStrings := strings.Split(row, ",")
		for _, aNumberAsString := range numbersAsStrings {
			aNumber, _ := strconv.Atoi(aNumberAsString)
			array = append(array, aNumber)
		}

	}
	sort.Sort(sort.IntSlice(array))
	return array
}

func (alg *Part1) Process(data []string) (error, interface{}) {
	input := Parse(data)

	costs := make(map[int]int, 0)
	median := utility.Median(input)

	costs[int(median)] = costOfAdjustingPositionsTo(input, int(median))

	mean := utility.Mean(input)
	costs[mean] = costOfAdjustingPositionsTo(input, mean)

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

func costOfAdjustingPositionsTo(numbers []int, targetNumber int) int {
	cost := 0
	for _, aNumber := range numbers {
		diff := targetNumber - aNumber
		diff = int(math.Abs(float64(diff)))
		cost += diff
	}
	return cost
}

func costExpensiveOfAdjustingPositionsTo(numbers []int, targetNumber int) int {
	cost := 0
	for _, aNumber := range numbers {
		diff := targetNumber - aNumber
		diff = int(math.Abs(float64(diff)))

		for x := 0; x < diff; x++ {
			cost += (x + 1)
		}

	}
	return cost
}
