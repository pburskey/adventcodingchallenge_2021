package main

import (
	"log"
	"sort"
	"strings"
)

type Part1 struct {
}

func (alg *Part1) Decypher(data []string) (input []*entry, countMap map[int]int, decodedMap map[string]*displaySegment) {
	input = Parse(data)

	countMap = make(map[int]int, 0)
	decodedMap = make(map[string]*displaySegment, 0)
	for _, anEntry := range input {
		keys := &cypherKeys{
			public: anEntry.signalPatterns,
		}
		keys.loadKeys()

		for _, anOutputValue := range anEntry.outputValues {

			cypherAsArrayOfStrings := strings.Split(anOutputValue, "")
			sort.Strings(cypherAsArrayOfStrings)
			anOutputValue = strings.Join(cypherAsArrayOfStrings, "")

			aCypher, _ := decodedMap[anOutputValue]
			if aCypher == nil {
				aCypher = &displaySegment{
					cypher: anOutputValue,
				}

				successful := aCypher.decode(keys)
				if successful {
					decodedMap[aCypher.cypher] = aCypher
				} else {
					log.Fatalln("Unable to decode: " + anOutputValue)
				}

			}

			if aCypher != nil && aCypher.decoded {
				if _, ok := countMap[aCypher.value]; !ok {
					countMap[aCypher.value] = 0
				}
				countMap[aCypher.value]++
			}

		}
	}

	return
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

	for aKey, aPrivateValue := range cypherKeys.private {
		if s.cypherContainsOnlyLetters(strings.Split(aPrivateValue, "")) {
			s.decoded = true
			s.value = aKey
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
	_, countMap, _ := alg.Decypher(data)

	count := 0
	for key, value := range countMap {
		if key == 1 || key == 4 || key == 8 || key == 7 {
			count += value
		}
	}

	return nil, count
}
