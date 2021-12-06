package main

type Part2 struct {
	numberOfDays int
}

func (alg *Part2) Process(data []string) (error, interface{}) {
	part1 := &Part1{
		numberOfDays: alg.numberOfDays,
	}
	part1.Process(data)
	return nil, nil
}
