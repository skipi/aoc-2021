package main

import (
	"fmt"
	"os"

	"kutryj.io/aoc21/pkg/day1"
	"kutryj.io/aoc21/pkg/day2"
)

func main() {
	fmt.Printf("%v", os.Args)

	if len(os.Args) != 2 {
		fmt.Println("Please provide a day number")
		os.Exit(1)
	}

	switch os.Args[1] {
	case "1":
		day1.Run()
	case "2":
		day2.Run()
	default:
		fmt.Println("Usage: aoc21 <day_number>")
	}

}
