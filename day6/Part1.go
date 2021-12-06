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
}

func Parse(data []string) []*Lanternfish {

	fish := make([]*Lanternfish, 0)

	for _, row := range data {
		strings := strings.Split(row, ",")
		for _, aString := range strings {
			aNumber, _ := strconv.Atoi(aString)
			aFish := &Lanternfish{internalTimer: aNumber}
			fish = append(fish, aFish)
		}
	}
	return fish
}

func (l *Lanternfish) startNewDay() *Lanternfish {
	var newFish *Lanternfish
	if l.internalTimer == 0 {
		l.internalTimer = 6
		newFish = &Lanternfish{internalTimer: 8}
	} else {
		l.internalTimer--
	}
	return newFish
}
func (alg *Part1) Process(data []string) (error, interface{}) {
	fish := Parse(data)

	for counter := 0; counter < alg.numberOfDays; counter++ {
		fmt.Println("Day: " + fmt.Sprint(counter) + " fish count: " + fmt.Sprint(len(fish)))
		for _, aFish := range fish {

			if aNewFish := aFish.startNewDay(); aNewFish != nil {
				fish = append(fish, aNewFish)
			}
		}
	}

	return nil, len(fish)
}
