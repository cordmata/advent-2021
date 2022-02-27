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
	fish []*lanternfish
}

func (s *school) elapseDay() {
	for _, f := range s.fish {
		_, baby := f.elapseDay()
		if baby != nil {
			s.fish = append(s.fish, baby)
		}
	}
}

type lanternfish struct {
	daysUntilDue int
}

func (l *lanternfish) elapseDay() (int, *lanternfish) {
	if l.daysUntilDue == 0 {
		l.daysUntilDue = matureLanternFishGestationDays
		return l.daysUntilDue, &lanternfish{babyLanternFishGestationDays}
	} else {
		l.daysUntilDue--
		return l.daysUntilDue, nil
	}
}

func part1(s school) int {
	for i := 0; i < 80; i++ {
		s.elapseDay()
	}
	return len(s.fish)
}

func part2(s school) int {
	return -1
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
	var fish []*lanternfish
	timers := utils.StringSliceToIntSlice(strings.Split(in, ","))
	for _, t := range timers {
		fish = append(fish, &lanternfish{t})
	}
	return school{fish}
}

var exampleInput school = processInput("3,4,3,1,2")

var actualInput school = processInput("1,1,1,2,1,5,1,1,2,1,4,1,4,1,1,1,1,1,1,4,1,1,1,1,4,1,1,5,1,3,1,2,1,1,1,2,1,1,1,4,1,1,3,1,5,1,1,1,1,3,5,5,2,1,1,1,2,1,1,1,1,1,1,1,1,5,4,1,1,1,1,1,3,1,1,2,4,4,1,1,1,1,1,1,3,1,1,1,1,5,1,3,1,5,1,2,1,1,5,1,1,1,5,3,3,1,4,1,3,1,3,1,1,1,1,3,1,4,1,1,1,1,1,2,1,1,1,4,2,1,1,5,1,1,1,2,1,1,1,1,1,1,1,1,2,1,1,1,1,1,5,1,1,1,1,3,1,1,1,1,1,3,4,1,2,1,3,2,1,1,2,1,1,1,1,4,1,1,1,1,4,1,1,1,1,1,2,1,1,4,1,1,1,5,3,2,2,1,1,3,1,5,1,5,1,1,1,1,1,5,1,4,1,2,1,1,1,1,2,1,3,1,1,1,1,1,1,2,1,1,1,3,1,4,3,1,4,1,3,2,1,1,1,1,1,3,1,1,1,1,1,1,1,1,1,1,2,1,5,1,1,1,1,2,1,1,1,3,5,1,1,1,1,5,1,1,2,1,2,4,2,2,1,1,1,5,2,1,1,5,1,1,1,1,5,1,1,1,2,1")
