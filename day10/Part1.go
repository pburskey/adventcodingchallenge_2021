package main

import (
	"adventcodingchallenge_2021/utility"
	"log"
	"strings"
)

type Part1 struct {
}

func Parse(data []string) []string {

	array := make([]string, 0)

	for _, row := range data {

		array = append(array, row)

	}
	return array
}

type command struct {
	inputAsArray    []string
	failurePosition int
}

type segment struct {
	opening string
	closing string
}

func findSegmentHaving(aString string) *segment {
	var foundSegment *segment
	for _, aSegment := range validSegments {
		if strings.EqualFold(aSegment.opening, aString) || strings.EqualFold(aSegment.closing, aString) {
			foundSegment = aSegment
			break
		}
	}
	return foundSegment
}

func isOpening(aString string) bool {
	itDoes := false
	var aSegment *segment
	aSegment = findSegmentHaving(aString)
	if aSegment != nil {
		itDoes = (strings.EqualFold(aSegment.opening, aString))
	}
	return itDoes
}

func isClosing(aString string) bool {
	itDoes := false
	var aSegment *segment
	aSegment = findSegmentHaving(aString)
	if aSegment != nil {
		itDoes = (strings.EqualFold(aSegment.closing, aString))
	}
	return itDoes
}

func closes(anInterestingString string, lastOnStack string) bool {
	itDoes := false
	var aSegment *segment
	aSegment = findSegmentHaving(lastOnStack)
	if aSegment != nil {
		itDoes = (strings.EqualFold(aSegment.opening, lastOnStack) && strings.EqualFold(aSegment.closing, anInterestingString))
	}
	return itDoes
}

func (c *command) parse() (bool, *utility.SimpleStringStack) {
	c.failurePosition = -1
	//fmt.Println(strings.Join(c.inputAsArray, ""))

	stack := utility.NewSimpleStringStack()
	for x, aCharacter := range c.inputAsArray {
		if isOpening(aCharacter) {
			stack.Push(aCharacter)
		} else if isClosing(aCharacter) {
			if closes(aCharacter, stack.Peek()) {
				stack.Pop()
			} else {
				c.failurePosition = x
				break
			}
		} else {
			log.Fatalln("What is this: " + aCharacter)
		}
	}

	return c.failurePosition == -1, stack
}

func (c *command) error() string {
	error := ""
	if c.failurePosition > -1 {
		error = c.inputAsArray[c.failurePosition]
	}
	return error
}

var (
	validSegments = []*segment{
		&segment{
			opening: "(",
			closing: ")",
		},
		&segment{
			opening: "[",
			closing: "]",
		},
		&segment{
			opening: "<",
			closing: ">",
		},
		&segment{
			opening: "{",
			closing: "}",
		},
	}
)

func (alg *Part1) Process(data []string) (error, interface{}) {
	input := Parse(data)

	failures := make([]string, 0)
	for _, commandString := range input {

		aCommand := &command{inputAsArray: strings.Split(commandString, "")}
		if ok, _ := aCommand.parse(); !ok {
			failures = append(failures, aCommand.error())
		}

	}
	result := 0
	for _, aCharacter := range failures {
		if aCharacter == ")" {
			result += 3
		} else if aCharacter == "]" {
			result += 57
		} else if aCharacter == "}" {
			result += 1197
		} else if aCharacter == ">" {
			result += 25137
		}
	}

	return nil, result
}
