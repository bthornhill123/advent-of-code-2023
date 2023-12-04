package main

import (
	"fmt"

	"github.com/bthornhill123/advent-of-code-2023/advent"
)

func main() {
	day := advent.One{}
	err := day.PartTwo()
	if err != nil {
		fmt.Println("An error occurred: ", err)
	}
}
