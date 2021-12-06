package main

type Part2 struct {
}

func (alg *Part2) Process(data []string, diag bool) (error, interface{}) {

	part1 := &Part1{}
	return part1.Process(data, diag)
}
