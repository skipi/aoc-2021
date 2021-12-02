package day2

import (
	"fmt"
	"strings"

	"kutryj.io/aoc21/pkg/utils"
)

func Run() {
	course := utils.LoadData("./inputs/day2.txt")

	fmt.Printf("\nPart 1: %d\n", Part1(course))

	fmt.Printf("Part 2: %d\n", Part2(course))
}

func Part1(course []string) int {
	x, y := LayCourse(course)

	return x * y
}

func Part2(course []string) int {
	submarine := NewSubmarine()

	for _, operation := range course {
		submarine.Execute(operation)
	}

	return submarine.Position()
}

type Submarine struct {
	depth      int
	horizontal int
	aim        int
}

func NewSubmarine() *Submarine {
	return &Submarine{
		depth:      0,
		horizontal: 0,
		aim:        0,
	}
}

func (s *Submarine) Execute(operation string) {
	op := strings.Split(operation, " ")
	command, units := op[0], utils.ParseInt(op[1])

	switch command {
	case "forward":
		s.horizontal += units
		s.depth += s.aim * units
	case "up":
		s.aim -= units
	case "down":
		s.aim += units
	}
}

func (s *Submarine) Position() int {
	return s.depth * s.horizontal
}

func LayCourse(course []string) (x int, y int) {
	x, y = 0, 0

	for _, operation := range course {
		op := strings.Split(operation, " ")
		command, units := op[0], utils.ParseInt(op[1])
		switch command {
		case "forward":
			x += units
		case "up":
			y -= units
		case "down":
			y += units
		}
	}
	return
}
