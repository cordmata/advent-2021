package main

import (
	"fmt"
	"log"
	"strings"
)

const p1Example = 1656
const p2Example = 195

type octopi [10][10]int

func (o *octopi) step() step {
	step := step{octopi: o}
	step.execute()
	return step
}

type step struct {
	octopi     *octopi
	flashCount int
	hasFlashed [10][10]bool
}

func (s *step) execute() {
	for r, row := range s.octopi {
		for c := range row {
			s.increment(r, c)
		}
	}
}

func (s *step) increment(r, c int) {
	if r < 0 || r > 9 || c < 0 || c > 9 || s.hasFlashed[r][c] {
		return
	}
	s.octopi[r][c]++
	if s.octopi[r][c] > 9 {
		s.flashCount++
		s.hasFlashed[r][c] = true
		s.octopi[r][c] = 0
		s.incrementNeighbors(r, c)
	}
}

func (s *step) incrementNeighbors(r, c int) {
	neighbors := []struct{ row, col int }{
		{r - 1, c},
		{r - 1, c - 1},
		{r - 1, c + 1},
		{r, c - 1},
		{r, c + 1},
		{r + 1, c},
		{r + 1, c - 1},
		{r + 1, c + 1},
	}
	for _, n := range neighbors {
		s.increment(n.row, n.col)
	}
}

func part1(o octopi) int {
	var totalFlashes int
	for i := 0; i < 100; i++ {
		totalFlashes += o.step().flashCount
	}
	return totalFlashes
}

func part2(o octopi) int {
	var stepCount int
	for {
		s := o.step()
		stepCount++
		if s.flashCount == 100 {
			break
		}
	}
	return stepCount
}

func processInput(s string) octopi {
	var octopi octopi
	for r, row := range strings.Split(s, "\n") {
		for c, val := range row {
			octopi[r][c] = int(val - '0')
		}
	}
	return octopi
}

func main() {
	exampleInput := processInput(exampleInput)
	actualInput := processInput(actualInput)

	if r := part1(exampleInput); r != p1Example {
		log.Fatalf("[PART 1]: expected %d, got %d", p1Example, r)
	}
	fmt.Println("The answer to part 1 is:", part1(actualInput))

	if r := part2(exampleInput); r != p2Example {
		log.Fatalf("[PART 2]: expected %d, got %d", p2Example, r)
	}
	fmt.Println("The answer to part 2 is:", part2(actualInput))
}

var exampleInput = `5483143223
2745854711
5264556173
6141336146
6357385478
4167524645
2176841721
6882881134
4846848554
5283751526`

var actualInput = `4472562264
8631517827
7232144146
2447163824
1235272671
5133527146
6511372417
3841841614
8621368782
3246336677`
