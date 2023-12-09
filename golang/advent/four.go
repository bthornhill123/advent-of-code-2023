package advent

import (
	"bufio"
	"os"
	"regexp"
	"strconv"
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

func (d Four) parseLine(line string) (cardNum int, candidates, winners []int) {
	numberPattern := regexp.MustCompile(`\d+`)

	card := strings.Split(line, ":")[0]
	cardNum, _ = strconv.Atoi(numberPattern.FindAllString(card, -1)[0])

	numbers := strings.Split(line, ":")[1]
	winnersStr := strings.Split(numbers, "|")[0]
	winnerStrings := numberPattern.FindAllString(winnersStr, -1)
	for _, winner := range winnerStrings {
		winnerInt, _ := strconv.Atoi(winner)
		winners = append(winners, winnerInt)
	}

	candidatesStr := strings.Split(numbers, "|")[1]
	candidateStrings := numberPattern.FindAllString(candidatesStr, -1)
	for _, candidate := range candidateStrings {
		candidateInt, _ := strconv.Atoi(candidate)
		candidates = append(candidates, candidateInt)
	}

	return
}

func (d Four) PartTwo() error {
	file, err := os.Open("inputs/four.txt")
	if err != nil {
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	copies := make(map[int]int)

	for scanner.Scan() {
		cardNum, candidates, winners := d.parseLine(scanner.Text())

		// Add a copy for the card number
		copies[cardNum]++

		// Create a set containing the winning numbers
		winnersSet := make(map[int]bool)
		for _, winner := range winners {
			winnersSet[winner] = true
		}

		// Count the number of matches
		matches := 0
		for _, candidate := range candidates {
			if _, ok := winnersSet[candidate]; ok {
				matches++
			}
		}

		// Each match adds a copy for each successive card
		nextCard := cardNum + 1
		for j := 0; j < matches; j++ {
			copies[nextCard] += copies[cardNum]
			nextCard++
		}
	}

	sum := 0
	for _, copies := range copies {
		sum += copies
	}

	println("Sum: ", sum)

	return nil
}
