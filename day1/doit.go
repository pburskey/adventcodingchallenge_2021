package main

import (
	"adventcodingchallenge_2021/utility"
	"fmt"
	"path/filepath"
)

const day = "1"

func main() {
	data, err := utility.ParseInputFileIntoRows(filepath.Join(utility.AssembleFilePathToDay(day), "data.txt"))
	if err != nil {
		panic(err)
	}

	_, solution_a := solution_part_a(data)
	fmt.Println(fmt.Sprintf("Result: %d", solution_a))

	_, solution_b := solution_part_b(data)
	fmt.Println(fmt.Sprintf("Result: %d", solution_b))

}

func solution_part_a(data []string) (error, interface{}) {
	algorithm := &Part1{}
	_, solution := algorithm.Process(data)
	return nil, solution

}

func solution_part_b(data []string) (error, interface{}) {
	algorithm := &Part2{}
	_, solution := algorithm.Process(data)
	return nil, solution

}
