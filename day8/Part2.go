package main

import (
	"adventcodingchallenge_2021/utility"
	"log"
	"sort"
	"strings"
)

type Part2 struct {
}

func (alg *Part2) Process(data []string) (error, interface{}) {
	input := Parse(data)

	signals := make([]int, 0)

	for _, anEntry := range input {
		keys := &cypherKeys{
			public: anEntry.signalPatterns,
		}
		keys.loadKeys()
		values := make([]int, 0)
		for _, anOutputValue := range anEntry.outputValues {

			cypherAsArrayOfStrings := strings.Split(anOutputValue, "")
			sort.Strings(cypherAsArrayOfStrings)
			anOutputValue = strings.Join(cypherAsArrayOfStrings, "")

			aCypher := &displaySegment{
				cypher: anOutputValue,
			}
			keys.decode(aCypher)

			if !aCypher.decoded {
				log.Fatalln("Unable to decode: " + anOutputValue)
			}
			//fmt.Printf("Cypher: %s Value: %d\n", anOutputValue, aCypher.value)
			values = append(values, aCypher.value)

		}
		signals = append(signals, utility.SliceToInt(values))
	}

	count := 0
	for _, aValue := range signals {
		count += aValue
	}

	return nil, count

}
