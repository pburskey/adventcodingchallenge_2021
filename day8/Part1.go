package main

import (
	"fmt"
	"sort"
	"strings"
)

type Part1 struct {
}

type entry struct {
	signalPatterns []string
	outputValues   []string
}

type displaySegment struct {
	cypher  string
	value   int
	decoded bool
}

type cypherKeys struct {
	public  []string
	private map[int]string
}

func (ck *cypherKeys) loadKeys() {
	ck.private = make(map[int]string, 0)

	/*
		get the easy ones
	*/
	for _, aKey := range ck.public {
		length := len(aKey)
		switch length {
		case 2:
			ck.private[1] = aKey
		case 3:
			ck.private[7] = aKey
		case 4:
			ck.private[4] = aKey
		case 7:
			ck.private[8] = aKey
		}
	}

	/*
		go after 6s
	*/
	for _, aKey := range ck.public {
		if len(aKey) == 6 {
			/*
				a 0 fully contains a 7
			*/
			if ck.notMapped(9) && contains(aKey, ck.private[4]) {
				ck.private[9] = aKey
			} else if ck.notMapped(0) && contains(aKey, ck.private[7]) {
				ck.private[0] = aKey
			} else {
				ck.private[6] = aKey
			}
		}
	}

	/*
		go after 5s
	*/
	for _, aKey := range ck.public {
		if len(aKey) == 5 {
			/*
				a 0 fully contains a 7
			*/
			if ck.notMapped(3) && contains(aKey, ck.private[7]) {
				ck.private[3] = aKey
			} else if ck.notMapped(5) && contains(ck.private[6], aKey) {
				ck.private[5] = aKey
			} else {
				ck.private[2] = aKey
			}
		}
	}

	ck.prettyPrintKeys()
}

func (ck *cypherKeys) prettyPrintKeys() {
	for key, value := range ck.private {
		fmt.Printf("Value: %d ... Cypher Key: %s\n", key, value)
	}
}

func (ck *cypherKeys) notMapped(aKey int) bool {
	_, ok := ck.private[aKey]
	return !ok
}

func contains(source string, target string) bool {

	//sourceAsArray := strings.Split(source, "")
	//sort.Strings(sourceAsArray)
	targetAsArray := strings.Split(target, "")
	//sort.Strings(targetAsArray)

	found := true
	for x := 0; found && x < len(targetAsArray); x++ {
		found = strings.Contains(source, targetAsArray[x])
	}

	return found

}

func (s *displaySegment) closelyMatches(aString string) bool {
	letters := strings.Split(aString, "")
	sort.Strings(letters)

	cypherAsArrayOfStrings := strings.Split(s.cypher, "")
	sort.Strings(cypherAsArrayOfStrings)

	found := true
	for x := 0; found && x < len(letters) && x < len(cypherAsArrayOfStrings); x++ {
		found = (strings.EqualFold(letters[x], cypherAsArrayOfStrings[x]))
	}

	itdoes := found
	return itdoes
}

func (s *displaySegment) cypherContainsOnlyLetters(letters []string) bool {
	itDoes := false
	cypherAsArrayOfStrings := strings.Split(s.cypher, "")
	//sort.Strings(cypherAsArrayOfStrings)
	sort.Strings(letters)

	found := true
	for x := 0; found && x < len(letters) && x < len(cypherAsArrayOfStrings); x++ {
		found = (strings.EqualFold(letters[x], cypherAsArrayOfStrings[x]))
	}
	itDoes = found
	return itDoes
}

func (s *displaySegment) decode(cypherKeys *cypherKeys) bool {

	cypherAsArrayOfStrings := strings.Split(s.cypher, "")
	sort.Strings(cypherAsArrayOfStrings)
	s.cypher = strings.Join(cypherAsArrayOfStrings, "")

	cypherKeys
	for aKey, aPrivateValue := range cypherKeys.private {
		if s.closestMatch(cypherKeys) {
			s.decoded = true
			s.value = aKey
			break
		}
	}

	return s.decoded
}

func Parse(data []string) []*entry {

	array := make([]*entry, 0)

	for _, row := range data {
		entry := &entry{
			signalPatterns: make([]string, 0),
			outputValues:   make([]string, 0),
		}
		array = append(array, entry)
		prePipe := true
		strings := strings.Split(row, " ")

		for _, aString := range strings {

			if aString == "|" {
				prePipe = false
			} else {
				if prePipe {
					entry.signalPatterns = append(entry.signalPatterns, aString)
				} else {
					entry.outputValues = append(entry.outputValues, aString)
				}
			}

		}

	}
	return array
}

func (alg *Part1) Process(data []string) (error, interface{}) {
	input := Parse(data)

	count := 0
	for _, anEntry := range input {

		for _, aKey := range anEntry.outputValues {
			length := len(aKey)
			switch length {
			case 2:
				count++
			case 3:
				count++
			case 4:
				count++
			case 7:
				count++
			}
		}

	}

	return nil, count
}
