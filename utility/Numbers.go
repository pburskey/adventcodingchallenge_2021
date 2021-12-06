package utility

import (
	"log"
	"sort"
	"strconv"
)

func LeastAndMax(numbers []int) (int, int) {

	least := numbers[0]
	max := numbers[0]
	for _, aNumber := range numbers {
		if aNumber < least {
			least = aNumber
		}
		if aNumber > max {
			max = aNumber
		}

	}
	return least, max
}

func OrderNumbersStartingWithAndEndingWith(numbers []int, start int, end int) []int {

	if numbers[0] == start && numbers[len(numbers)-1] == end {
		return numbers
	} else {

		sort.Sort(sort.Reverse(sort.IntSlice(numbers)))
		return numbers
	}
	return numbers
}

func NumbersBetween(start int, end int) []int {
	least, max := LeastAndMax([]int{start, end})

	numbers := make([]int, 0)
	for x := least; x <= max; x++ {
		numbers = append(numbers, x)
	}
	return numbers
}

func StringToInt(aString string) int {
	aNumber, err := strconv.Atoi(aString)
	if err != nil {
		panic("Expected a number and got: " + aString)
	}
	return aNumber
}

func BytesToInt(byteArray []byte) int {
	/*
		lets convert to a string
	*/
	bytesAsString := string(byteArray)
	//for len(bytesAsString) < 8{
	//	bytesAsString = "0" + bytesAsString
	//}

	anInt, err := strconv.ParseInt(bytesAsString, 2, 64)

	if err != nil {
		log.Fatalln(err)
	}
	return int(anInt)
}
