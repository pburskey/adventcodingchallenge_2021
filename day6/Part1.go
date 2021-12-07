package main

import (
	"fmt"
	"strconv"
	"strings"
)

type Part1 struct {
	numberOfDays int
}

type Lanternfish struct {
	internalTimer int
	count         int
}

func Parse(data []string) []*Lanternfish {

	fish := make([]*Lanternfish, 0)

	for _, row := range data {
		strings := strings.Split(row, ",")
		for _, aString := range strings {
			aNumber, _ := strconv.Atoi(aString)
			aFish := &Lanternfish{internalTimer: aNumber, count: 1}
			fish = append(fish, aFish)
		}
	}
	return fish
}

func (l *Lanternfish) startNewDay() *Lanternfish {
	var newFish *Lanternfish
	if l.internalTimer == 0 {
		l.internalTimer = 6
		newFish = &Lanternfish{internalTimer: 8, count: l.count}
	} else {
		l.internalTimer--
	}
	return newFish
}

func howManyFish(aMap map[int]*Lanternfish) int {
	count := 0
	for _, aFish := range aMap {
		count += aFish.count
	}
	return count
}

func (alg *Part1) Process(data []string) (error, interface{}) {
	fish := Parse(data)
	fishMap := make(map[int]*Lanternfish, 0)

	for _, aFish := range fish {
		updateFishInMap(fishMap, aFish)
	}
	counter := 0
	for counter = 0; counter < alg.numberOfDays; counter++ {
		fmt.Println("Day: " + fmt.Sprint(counter) + " fish count: " + fmt.Sprint(howManyFish(fishMap)))
		updatedMap := make(map[int]*Lanternfish, 0)
		for _, aFish := range fishMap {
			if aNewFish := aFish.startNewDay(); aNewFish != nil {
				updateFishInMap(updatedMap, aNewFish)
			}
			updateFishInMap(updatedMap, aFish)
		}
		fishMap = updatedMap
	}
	fmt.Println("Day: " + fmt.Sprint(counter) + " fish count: " + fmt.Sprint(howManyFish(fishMap)))
	return nil, howManyFish(fishMap)
}

func updateFishInMap(aMap map[int]*Lanternfish, aFish *Lanternfish) {
	if existingFish, ok := aMap[aFish.internalTimer]; ok {
		existingFish.count += aFish.count
	} else {
		aMap[aFish.internalTimer] = aFish
	}

}
