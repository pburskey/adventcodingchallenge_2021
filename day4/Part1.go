package main

import (
	"strconv"
	"strings"
)

type Part1 struct {
}

type BingoBoard struct {
	squares [][]*BingoSquare
	winner  bool
}

func (bb *BingoBoard) PlayNumber(aNumber int) (*BingoSquare, bool) {
	var foundSquare *BingoSquare
	if bb.winner {
		return nil, false
	}
	for x, _ := range bb.squares {

		for _, square := range bb.squares[x] {
			if square.value == aNumber {
				square.marked = true
				foundSquare = square
				break
			}
		}
		if foundSquare != nil {
			break
		}
	}

	return foundSquare, bb.Validate()
}

func (bb *BingoBoard) GetHorizontalSliceFromAt(y int) []*BingoSquare {
	sliceOfSquares := bb.squares[y]
	return sliceOfSquares

}

func (bb *BingoBoard) GetVerticalSliceFromAt(x int) []*BingoSquare {
	sliceOfSquares := make([]*BingoSquare, 0)

	for y := 0; y < len(bb.squares); y++ {
		square := bb.squares[y][x]
		sliceOfSquares = append(sliceOfSquares, square)
	}

	return sliceOfSquares

}

func (bb *BingoBoard) Validate() bool {
	winner := false

	/*
		check across first
	*/
	for y := 0; y < len(bb.squares); y++ {
		aRow := bb.GetHorizontalSliceFromAt(y)
		found := true
		for z := 0; z < len(aRow) && found; z++ {
			aSquare := aRow[z]
			found = found && aSquare.marked
		}
		if found {
			return true
		}
	}

	/*
		get our depth
	*/
	widthOfX := len(bb.squares[0])

	/*
		check vertical next
	*/
	for x := 0; x < widthOfX; x++ {
		aColumn := bb.GetVerticalSliceFromAt(x)
		found := true
		for z := 0; z < len(aColumn) && found; z++ {
			aSquare := aColumn[z]
			found = found && aSquare.marked
		}
		if found {
			return true
		}

	}

	return winner
}

func (b *BingoBoard) Unmarked() []*BingoSquare {
	sliceOfSquares := make([]*BingoSquare, 0)
	for y, _ := range b.squares {
		for x, _ := range b.squares[y] {
			aSquare := b.squares[y][x]
			if !aSquare.marked {
				sliceOfSquares = append(sliceOfSquares, aSquare)
			}
		}
	}
	return sliceOfSquares
}

type BingoSquare struct {
	x      int
	y      int
	value  int
	marked bool
}

type Bingo struct {
	input          []int
	playedNumbers  []int
	boards         []*BingoBoard
	winners        []*BingoBoard
	winningNumbers []int
}

func (b *Bingo) Play() {

	for _, aNumber := range b.input {
		b.playedNumbers = append(b.playedNumbers, aNumber)
		for _, aBoard := range b.boards {
			foundSquare, winner := aBoard.PlayNumber(aNumber)
			if foundSquare != nil && winner {
				b.winners = append(b.winners, aBoard)
				b.winningNumbers = append(b.winningNumbers, aNumber)
				aBoard.winner = true
			}
		}
	}
}

func Parse(data []string) (input []int, boards []*BingoBoard) {

	input = make([]int, 0)
	boards = make([]*BingoBoard, 0)

	var currentBoard *BingoBoard

	for i, row := range data {

		if i == 0 {
			strings := strings.Split(row, ",")
			for _, aString := range strings {
				anInt, _ := strconv.Atoi(aString)
				input = append(input, anInt)
			}
		} else {
			if row == "" {
				currentBoard = &BingoBoard{}
				boards = append(boards, currentBoard)
				currentBoard.squares = make([][]*BingoSquare, 0)
			} else {
				strings := strings.Split(row, " ")
				sliceOfSquares := make([]*BingoSquare, 0)
				for x, aString := range strings {
					anInt, err := strconv.Atoi(aString)
					if err == nil {
						aBingoSquare := &BingoSquare{
							x:      x,
							y:      0,
							value:  anInt,
							marked: false,
						}
						sliceOfSquares = append(sliceOfSquares, aBingoSquare)
					}

				}
				currentBoard.squares = append(currentBoard.squares, sliceOfSquares)

			}

		}

	}
	return input, boards
}

func (alg *Part1) Process(data []string) (error, interface{}) {

	input, boards := Parse(data)

	bingo := &Bingo{
		input:  input,
		boards: boards,
	}

	bingo.Play()
	winningNumber := bingo.winningNumbers[0]
	winningBoard := bingo.winners[0]
	pieces := winningBoard.Unmarked()
	var sum int
	for _, piece := range pieces {
		sum += piece.value
	}

	return nil, (sum * winningNumber)
}
