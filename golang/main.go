package main

import (
	"fmt"

	"github.com/bthornhill123/advent-of-code-2023/advent"
)

func main() {
	day := advent.Six{}

	fmt.Print("Part One -> ")
	err := day.PartOne()
	if err != nil {
		fmt.Println("An error occurred: ", err)
	}

	fmt.Print("Part Two -> ")
	err = day.PartTwo()
	if err != nil {
		fmt.Println("An error occurred: ", err)
	}
}
