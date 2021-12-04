package main

type Part2 struct {
}

type positionWithAim struct {
	Horizontal int
	Depth      int
	Aim        int
}

func (alg *Part2) Process(data []string) (error, interface{}) {

	commands := Parse(data)
	currentPosition := positionWithAim{
		Aim:        0,
		Horizontal: 0,
		Depth:      0,
	}

	for _, aCommand := range commands {

		if aCommand.Dir == Up {
			currentPosition.Aim -= aCommand.Value

		} else if aCommand.Dir == Down {
			currentPosition.Aim += aCommand.Value

		} else if aCommand.Dir == Forward {
			currentPosition.Horizontal += aCommand.Value
			if currentPosition.Aim > 0 {
				currentPosition.Depth += (aCommand.Value * currentPosition.Aim)
			}
		}
	}
	return nil, (currentPosition.Horizontal * currentPosition.Depth)
}
