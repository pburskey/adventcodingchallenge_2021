package main

type Part2 struct {
}

func (alg *Part2) Process(data []string) (error, interface{}) {
	metrics := Parse(data)

	oxygen := metrics.oxygen()
	c02 := metrics.c02Scrubber()

	return nil, (oxygen * c02)
}
