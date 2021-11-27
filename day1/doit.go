package main

import (
	"adventcodingchallenge_2021/utility"
	"fmt"
	"path/filepath"
	"strconv"
)

const day = "1"

func main() {
	data, err := utility.ParseInputFileIntoRows(filepath.Join(utility.AssembleFilePathToDay(day), "data.txt"))
	if err != nil {
		panic(err)
	}

	solution := Solution(data)
	fmt.Println(fmt.Sprintf("Result: %d", solution))
}

func Solution(data []string) interface{} {
	sum := 0
	for i, aRow := range data {
		fmt.Println(fmt.Sprintf("Row: %d ...... %s", i, aRow))
		aNumber, _ := strconv.Atoi(aRow)
		sum += aNumber
	}
	return sum
}
