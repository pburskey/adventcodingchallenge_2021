package main

import (
	"strconv"
	"strings"
)

type Part1 struct {
}

type BingoBoard struct {
	squares [][]*BingoSquare
}

func (b *BingoBoard) Unmarked() []*BingoSquare {
	return nil
}

type BingoSquare struct {
	x      int
	y      int
	value  int
	marked bool
}

type Bingo struct {
	input  []int
	boards []*BingoBoard
}

func (b *Bingo) Play() (winningNumber int, winningBoard *BingoBoard) {
	return 0, nil
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
			} else {
				strings := strings.Split(row, " ")
				for x, aString := range strings {
					anInt, _ := strconv.Atoi(aString)
					aBingoSquare := &BingoSquare{
						x:      0,
						y:      0,
						value:  anInt,
						marked: false,
					}
					currentBoard.squares = append(currentBoard.squares, aBingoSquare)

				}
			}

		}

	}
	return input, nil
}

func (alg *Part1) Process(data []string) (error, interface{}) {

	input, boards := Parse(data)

	bingo := &Bingo{
		input:  input,
		boards: boards,
	}

	winningNumber, winningBoard := bingo.Play()
	pieces := winningBoard.Unmarked()
	var sum int
	for _, piece := range pieces {
		sum += piece.value
	}

	return nil, (sum * winningNumber)
}
