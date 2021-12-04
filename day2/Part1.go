package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"
)

type Part1 struct {
	increase int
	decrease int
	previous int
	match    int
}

type Direction int

const (
	Undefined Direction = iota
	Forward
	Backward
	Up
	Down
)

type Command struct {
	Dir   Direction
	Value int
}

type position struct {
	Horizontal int
	Depth      int
}

func Parse(data []string) []Command {
	commands := make([]Command, 0)
	for _, row := range data {
		aCommand := Command{}
		aCommand.parse(row)
		commands = append(commands, aCommand)
	}
	return commands
}

func (c *Command) parse(data string) {
	anArray := strings.Split(data, " ")

	c.Value, _ = strconv.Atoi(anArray[1])
	aDirectionAsString := anArray[0]
	if strings.EqualFold(aDirectionAsString, "forward") {
		c.Dir = Forward
	} else if strings.EqualFold(aDirectionAsString, "up") {
		c.Dir = Up
	} else if strings.EqualFold(aDirectionAsString, "down") {
		c.Dir = Down
	} else {
		log.Fatalln(fmt.Printf("Direction: %s is not understood", aDirectionAsString))
	}
}

func (alg *Part1) Process(data []string) (error, interface{}) {

	commands := Parse(data)
	currentPosition := position{
		Horizontal: 0,
		Depth:      0,
	}

	for _, aCommand := range commands {

		if aCommand.Dir == Up {
			currentPosition.Depth -= aCommand.Value

		} else if aCommand.Dir == Down {
			currentPosition.Depth += aCommand.Value

		} else if aCommand.Dir == Forward {
			currentPosition.Horizontal += aCommand.Value

		}

	}
	return nil, (currentPosition.Horizontal * currentPosition.Depth)
}
