package advent

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Two struct{}

func (d Two) PartOne() error {
	file, err := os.Open("inputs/two.txt")
	if err != nil {
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	sum := 0
	for scanner.Scan() {
		line := scanner.Text()
		lineParts := strings.Split(line, ":")

		gameID := strings.Split(lineParts[0], " ")[1]
		sets := strings.Split(lineParts[1], ";")
		if d.areSetsValid(sets) {
			num, _ := strconv.Atoi(gameID)
			sum += num
		}
	}

	fmt.Println("Sum: ", sum)

	return nil
}

func (d Two) areSetsValid(sets []string) bool {
	limits := map[string]int{
		"red":   12,
		"green": 13,
		"blue":  14,
	}

	for _, set := range sets {
		pairs := strings.Split(set, ",")
		for _, pair := range pairs {
			pairParts := strings.Split(strings.TrimSpace(pair), " ")
			num, _ := strconv.Atoi(pairParts[0])
			color := pairParts[1]
			if num > limits[color] {
				return false
			}
		}
	}

	return true
}

func (d Two) PartTwo() error {
	file, err := os.Open("inputs/two.txt")
	if err != nil {
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	sum := 0
	for scanner.Scan() {
		line := scanner.Text()
		lineParts := strings.Split(line, ":")
		sets := strings.Split(lineParts[1], ";")
		sum += d.minCubePower(sets)
	}

	fmt.Println("Sum: ", sum)
	return nil
}

func (d Two) minCubePower(sets []string) int {
	mins := map[string]int{
		"red":   0,
		"green": 0,
		"blue":  0,
	}

	for _, set := range sets {
		pairs := strings.Split(set, ",")
		for _, pair := range pairs {
			pairParts := strings.Split(strings.TrimSpace(pair), " ")
			num, _ := strconv.Atoi(pairParts[0])
			color := pairParts[1]
			if num > mins[color] {
				mins[color] = num
			}
		}
	}

	power := 1
	for _, min := range mins {
		power *= min
	}

	return power
}
