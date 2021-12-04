package main

import (
	"adventcodingchallenge_2021/utility"
)

type Part1 struct {
}

type BitPositionTracker struct {
	value      byte
	occurances int
}

type Metrics struct {
	parsedInput [][]byte
	trackers    []map[byte]*BitPositionTracker
}

func (m *Metrics) gamma() int {
	byteArray := make([]byte, 0)

	for _, aMap := range m.trackers {
		var interestedTracker *BitPositionTracker

		for _, aTracker := range aMap {
			if interestedTracker == nil {
				interestedTracker = aTracker
			} else if aTracker.occurances > interestedTracker.occurances {
				interestedTracker = aTracker
			}
		}
		byteArray = append(byteArray, interestedTracker.value)
	}

	anInt := utility.BytesToInt(byteArray)

	return anInt
}

func (m *Metrics) oxygen() int {

	workingMetric := &Metrics{
		parsedInput: m.parsedInput,
		trackers:    m.trackers,
	}
	currentIndex := 0

	for len(workingMetric.parsedInput) > 1 {

		for i, aMap := range workingMetric.trackers {
			if i == currentIndex {
				var interestedTracker *BitPositionTracker

				interestingByte := byte(49)
				for _, aTracker := range aMap {
					if interestedTracker == nil {
						interestedTracker = aTracker
					} else if aTracker.occurances > interestedTracker.occurances {
						interestedTracker = aTracker
					} else if aTracker.occurances == interestedTracker.occurances && aTracker.value == interestingByte {
						interestedTracker = aTracker
					}
				}

				workingMetric.parsedInput = gatherFromHavingValueInPosition(workingMetric.parsedInput, interestedTracker.value, i)
				workingMetric.BuildTrackers()
				currentIndex++
				break

			}

		}
	}

	valueInBytes := workingMetric.parsedInput[0]
	anInt := utility.BytesToInt(valueInBytes)
	return anInt

}

func gatherFromHavingValueInPosition(bytes [][]byte, aByte byte, position int) [][]byte {
	filteredBytes := make([][]byte, 0)
	for _, thisBytes := range bytes {
		if thisBytes[position] == aByte {
			filteredBytes = append(filteredBytes, thisBytes)
		}
	}
	return filteredBytes
}

func (m *Metrics) epsilon() int {
	byteArray := make([]byte, 0)
	for _, aMap := range m.trackers {
		var interestedTracker *BitPositionTracker

		for _, aTracker := range aMap {
			if interestedTracker == nil {
				interestedTracker = aTracker
			} else if aTracker.occurances < interestedTracker.occurances {
				interestedTracker = aTracker
			}
		}
		byteArray = append(byteArray, interestedTracker.value)
	}

	anInt := utility.BytesToInt(byteArray)

	return anInt
}

func (m *Metrics) c02Scrubber() int {
	workingMetric := &Metrics{
		parsedInput: m.parsedInput,
		trackers:    m.trackers,
	}
	currentIndex := 0

	for len(workingMetric.parsedInput) > 1 {

		for i, aMap := range workingMetric.trackers {
			if i == currentIndex {
				var interestedTracker *BitPositionTracker

				interestingByte := byte(48)
				for _, aTracker := range aMap {
					if interestedTracker == nil {
						interestedTracker = aTracker
					} else if aTracker.occurances < interestedTracker.occurances {
						interestedTracker = aTracker
					} else if aTracker.occurances == interestedTracker.occurances && aTracker.value == interestingByte {
						interestedTracker = aTracker
					}
				}

				workingMetric.parsedInput = gatherFromHavingValueInPosition(workingMetric.parsedInput, interestedTracker.value, i)
				workingMetric.BuildTrackers()
				currentIndex++
				break

			}

		}
	}

	valueInBytes := workingMetric.parsedInput[0]
	anInt := utility.BytesToInt(valueInBytes)
	return anInt
}

func (m *Metrics) BuildTrackers() {
	m.trackers = make([]map[byte]*BitPositionTracker, 0)
	for _, aParsedDataRow := range m.parsedInput {

		if aParsedDataRow != nil {

			for i, aByte := range aParsedDataRow {

				var trackerMap map[byte]*BitPositionTracker
				if i < len(m.trackers) {
					trackerMap = m.trackers[i]
				}

				if trackerMap == nil {
					trackerMap = make(map[byte]*BitPositionTracker)
					m.trackers = append(m.trackers, trackerMap)
				}

				aTracker, found := trackerMap[aByte]
				if !found {
					aTracker = &BitPositionTracker{
						value:      aByte,
						occurances: 0,
					}
					trackerMap[aByte] = aTracker
				}
				aTracker.occurances++
			}
		}
	}
}

func ParseBytes(data [][]byte) *Metrics {

	metrics := &Metrics{
		parsedInput: data,
	}
	metrics.BuildTrackers()

	return metrics
}

func Parse(data []string) *Metrics {

	parsedData := make([][]byte, len(data))

	for i, row := range data {
		bytes := []byte(row)
		parsedData[i] = bytes
	}
	return ParseBytes(parsedData)
}

func (alg *Part1) Process(data []string) (error, interface{}) {

	metrics := Parse(data)

	gamma := metrics.gamma()
	epsilon := metrics.epsilon()

	return nil, (gamma * epsilon)
}
