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

func (ck *cypherKeys) decode(s *displaySegment) bool {

	matches := make([]int, 0)
	for aValue, aPrivateKey := range ck.private {
		privateContainsThisCypher := fullyContains(aPrivateKey, s.cypher)

		if privateContainsThisCypher {
			matches = append(matches, aValue)
		}
	}

	for _, aKey := range matches {
		aPrivate := ck.private[aKey]
		if strings.EqualFold(aPrivate, s.cypher) {
			s.decoded = true
			s.value = aKey
			break
		}
	}

	return s.decoded
}

func (ck *cypherKeys) loadKeys() {
	ck.private = make(map[int]string, 0)

	for x, aString := range ck.public {

		aStringAsArray := strings.Split(aString, "")
		sort.Strings(aStringAsArray)
		ck.public[x] = strings.Join(aStringAsArray, "")
	}

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
			if ck.notMapped(9) && fullyContains(aKey, ck.private[4]) {
				ck.private[9] = aKey
			} else if ck.notMapped(0) && fullyContains(aKey, ck.private[7]) {
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
			if ck.notMapped(3) && fullyContains(aKey, ck.private[7]) {
				ck.private[3] = aKey
			} else if ck.notMapped(5) && fullyContains(ck.private[6], aKey) {
				ck.private[5] = aKey
			} else {
				ck.private[2] = aKey
			}
		}
	}

	//ck.prettyPrintKeys()
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

func fullyContains(source string, target string) bool {

	sourceAsArray := strings.Split(source, "")
	sort.Strings(sourceAsArray)
	targetAsArray := strings.Split(target, "")
	sort.Strings(targetAsArray)

	found := true
	for x := 0; found && x < len(targetAsArray); x++ {
		found = strings.Contains(source, targetAsArray[x])
	}

	return found

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
