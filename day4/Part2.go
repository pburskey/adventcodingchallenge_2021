package main

type Part2 struct {
}

func (alg *Part2) Process(data []string) (error, interface{}) {

	input, boards := Parse(data)

	bingo := &Bingo{
		input:  input,
		boards: boards,
	}

	bingo.Play()
	winningNumber := bingo.winningNumbers[len(bingo.winningNumbers)-1]
	winningBoard := bingo.winners[len(bingo.winners)-1]
	pieces := winningBoard.Unmarked()
	var sum int
	for _, piece := range pieces {
		sum += piece.value
	}

	return nil, (sum * winningNumber)
}
