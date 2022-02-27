package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/cordmata/advent-2021/utils"
)

const p1Example = 5934
const p2Example = 26984457539
const matureLanternFishGestationDays = 6
const babyLanternFishGestationDays = 8

type school struct {
	fish map[int]int // counts of fish grouped by due date
}

func (s *school) elapseDay() {
	prevDay := s.fish
	fish := make(map[int]int)
	for daysUntilDue, count := range prevDay {
		if daysUntilDue == 0 {
			fish[babyLanternFishGestationDays] += count
			fish[matureLanternFishGestationDays] += count
		} else {
			fish[daysUntilDue-1] += count
		}
	}
	s.fish = fish
}

func (s school) fishCount() int {
	var numFishies int
	for _, count := range s.fish {
		numFishies += count
	}
	return numFishies
}

func simulate(s school, days int) int {
	for i := 0; i < days; i++ {
		s.elapseDay()
	}
	return s.fishCount()
}

func part1(s school) int {
	return simulate(s, 80)
}

func part2(s school) int {
	return simulate(s, 256)
}

func main() {
	if p1 := part1(exampleInput); p1 != p1Example {
		log.Fatalln("[ERROR]", p1, "!=", p1Example)
	}
	fmt.Println("[PART 1 ANSWER]", part1(actualInput))
	if p2 := part2(exampleInput); p2 != p2Example {
		log.Fatalln("[ERROR]", p2, "!=", p2Example)
	}
	fmt.Println("[PART 2 ANSWER]", part2(actualInput))
}

func processInput(in string) school {
	fish := make(map[int]int)
	timers := utils.StringSliceToIntSlice(strings.Split(in, ","))
	for _, t := range timers {
		fish[t]++
	}
	return school{fish}
}

var exampleInput school = processInput("3,4,3,1,2")

var actualInput school = processInput("1,1,1,2,1,5,1,1,2,1,4,1,4,1,1,1,1,1,1,4,1,1,1,1,4,1,1,5,1,3,1,2,1,1,1,2,1,1,1,4,1,1,3,1,5,1,1,1,1,3,5,5,2,1,1,1,2,1,1,1,1,1,1,1,1,5,4,1,1,1,1,1,3,1,1,2,4,4,1,1,1,1,1,1,3,1,1,1,1,5,1,3,1,5,1,2,1,1,5,1,1,1,5,3,3,1,4,1,3,1,3,1,1,1,1,3,1,4,1,1,1,1,1,2,1,1,1,4,2,1,1,5,1,1,1,2,1,1,1,1,1,1,1,1,2,1,1,1,1,1,5,1,1,1,1,3,1,1,1,1,1,3,4,1,2,1,3,2,1,1,2,1,1,1,1,4,1,1,1,1,4,1,1,1,1,1,2,1,1,4,1,1,1,5,3,2,2,1,1,3,1,5,1,5,1,1,1,1,1,5,1,4,1,2,1,1,1,1,2,1,3,1,1,1,1,1,1,2,1,1,1,3,1,4,3,1,4,1,3,2,1,1,1,1,1,3,1,1,1,1,1,1,1,1,1,1,2,1,5,1,1,1,1,2,1,1,1,3,5,1,1,1,1,5,1,1,2,1,2,4,2,2,1,1,1,5,2,1,1,5,1,1,1,1,5,1,1,1,2,1")
