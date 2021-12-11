package main

import (
	"log"
	"sort"
	"strings"
)

type Part2 struct {
}

func (alg *Part2) Process(data []string) (error, interface{}) {
	input := Parse(data)

	stacks := make([]*SimpleStringStack, 0)
	for _, commandString := range input {

		aCommand := &command{inputAsArray: strings.Split(commandString, "")}
		if ok, stack := aCommand.parse(); ok {
			stacks = append(stacks, stack)
		}

	}

	closureStacks := make([]*SimpleStringStack, 0)
	for _, aStack := range stacks {
		aStackInProgress := &SimpleStringStack{
			stack: make([]string, 0),
		}
		closureStacks = append(closureStacks, aStackInProgress)

		for aStack.hasMore() {
			aString := aStack.pop()
			if isClosing(aString) {
				log.Panic("Not understood what to do with " + aString)

			} else if isOpening(aString) {
				aSegment := findSegmentHaving(aString)
				aStackInProgress.push(aSegment.closing)
			}
		}

	}

	results := make([]int, 0)
	for _, aStack := range closureStacks {
		runningResult := 0
		characters := aStack.reverse()
		for _, aCharacter := range characters {

			if aCharacter == ")" {
				runningResult *= 5
				runningResult += 1
			} else if aCharacter == "]" {
				runningResult *= 5
				runningResult += 2
			} else if aCharacter == "}" {
				runningResult *= 5
				runningResult += 3
			} else if aCharacter == ">" {
				runningResult *= 5
				runningResult += 4
			}
		}
		results = append(results, runningResult)
	}
	sort.Ints(results)

	result := results[len(results)/2]
	//for _, aCharacter := range failures {
	//	if aCharacter == ")" {
	//		result += 3
	//	} else if aCharacter == "]" {
	//		result += 57
	//	} else if aCharacter == "}" {
	//		result += 1197
	//	} else if aCharacter == ">" {
	//		result += 25137
	//	}
	//}

	return nil, result
}
