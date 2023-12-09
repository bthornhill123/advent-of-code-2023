package advent

import (
	"bufio"
	"os"
	"regexp"
	"strings"
)

type Four struct{}

func (d Four) PartOne() error {
	file, err := os.Open("inputs/four.txt")
	if err != nil {
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	numberPattern := regexp.MustCompile(`\d+`)

	sum := 0
	for scanner.Scan() {

		card := strings.Split(scanner.Text(), ":")[1]
		winnersStr := strings.Split(card, "|")[0]
		candidatesStr := strings.Split(card, "|")[1]

		winners := numberPattern.FindAllString(winnersStr, -1)
		winnersSet := make(map[string]bool)
		for _, winner := range winners {
			winnersSet[winner] = true
		}

		cardSum := 0
		candidates := numberPattern.FindAllString(candidatesStr, -1)
		for _, candidate := range candidates {
			if _, ok := winnersSet[candidate]; ok {
				if cardSum == 0 {
					cardSum = 1
				} else {
					cardSum *= 2
				}
			}
		}

		sum += cardSum
	}

	println("Sum: ", sum)

	return nil
}

func (d Four) PartTwo() error {
	file, err := os.Open("inputs/four.txt")
	if err != nil {
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		_ = scanner.Text()
	}

	sum := 0
	println("Sum: ", sum)

	return nil
}
