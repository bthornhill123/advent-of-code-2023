package advent

import (
	"bufio"
	"os"
	"regexp"
	"strconv"
)

type Six struct{}

func (d Six) PartOne() error {
	file, err := os.Open("inputs/six.txt")
	if err != nil {
		return err
	}
	defer file.Close()

	numberPattern := regexp.MustCompile(`\d+`)

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	line := scanner.Text()
	times := numberPattern.FindAllString(line, -1)

	scanner.Scan()
	line = scanner.Text()
	distances := numberPattern.FindAllString(line, -1)

	multiplier := 1

	for i := 0; i < len(times); i++ {
		winningMoves := 0
		time, _ := strconv.Atoi(times[i])
		distance, _ := strconv.Atoi(distances[i])

		for holdTime := 1; holdTime < time; holdTime++ {
			travelTime := time - holdTime
			speed := holdTime
			distanceTraveled := travelTime * speed
			if distanceTraveled > distance {
				winningMoves++
			}
		}

		multiplier *= winningMoves
	}

	return nil
}

func (d Six) PartTwo() error {
	file, err := os.Open("inputs/six.txt")
	if err != nil {
		return err
	}
	defer file.Close()

	numberPattern := regexp.MustCompile(`\d+`)

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	line := scanner.Text()
	times := numberPattern.FindAllString(line, -1)
	totalTime := ""
	for _, time := range times {
		totalTime += time
	}

	scanner.Scan()
	line = scanner.Text()
	distances := numberPattern.FindAllString(line, -1)
	totalDistance := ""
	for _, distance := range distances {
		totalDistance += distance
	}

	time, _ := strconv.Atoi(totalTime)
	distance, _ := strconv.Atoi(totalDistance)

	winningMoves := 0
	for holdTime := 1; holdTime < time; holdTime++ {
		travelTime := time - holdTime
		speed := holdTime
		distanceTraveled := travelTime * speed
		if distanceTraveled > distance {
			winningMoves++
		}
	}

	return nil
}
