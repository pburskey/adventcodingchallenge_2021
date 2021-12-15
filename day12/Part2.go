package main

type Part2 struct {
}

func (alg *Part2) Process(data []string) (error, interface{}) {

	part1 := &Part1{}
	part1.Process(data)

	return nil, 0
}
