package advent

import (
	"bufio"
	"math"
	"os"
	"regexp"
	"strconv"
)

type Five struct{}

type MapEntry struct {
	sourceStart      int
	destinationStart int
	length           int
}

func NewMapEntry(line string) MapEntry {
	numberPattern := regexp.MustCompile(`\d+`)
	numbers := numberPattern.FindAllString(line, -1)
	sourceStart, _ := strconv.Atoi(numbers[1])
	destinationStart, _ := strconv.Atoi(numbers[0])
	length, _ := strconv.Atoi(numbers[2])

	return MapEntry{
		sourceStart:      sourceStart,
		destinationStart: destinationStart,
		length:           length,
	}
}

func (m MapEntry) getMapped(value int) (int, bool) {
	if m.sourceStart <= value && value <= m.sourceStart+m.length {
		spread := value - m.sourceStart
		return m.destinationStart + spread, true
	}

	return 0, false
}

func (d Five) PartOne() error {
	file, err := os.Open("inputs/five.txt")
	if err != nil {
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	numberPattern := regexp.MustCompile(`\d+`)

	seedsLine := scanner.Text()
	seeds := numberPattern.FindAllString(seedsLine, -1)

	lowestLocation := math.MaxInt
	mappedNextLevel := false
	for _, seed := range seeds {
		value, _ := strconv.Atoi(seed)
		for scanner.Scan() {
			line := scanner.Text()

			// Skip empty lines
			if len(line) == 0 {
				continue
			}

			if line[0] >= '0' && line[0] <= '9' {
				if !mappedNextLevel {
					mapEntry := NewMapEntry(line)
					nextValue, ok := mapEntry.getMapped(value)
					if ok {
						value = nextValue
						mappedNextLevel = true
					}
				}
			} else {
				mappedNextLevel = false
			}
		}

		if value < lowestLocation {
			lowestLocation = value
		}

		// reset the scanner to the first line in file
		file.Seek(0, 0)
		scanner = bufio.NewScanner(file)
	}

	println("Lowest Location: ", lowestLocation)

	return nil
}

func (d Five) PartTwo() error {
	file, err := os.Open("inputs/five.txt")
	if err != nil {
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	numberPattern := regexp.MustCompile(`\d+`)

	seedsLine := scanner.Text()
	seeds := numberPattern.FindAllString(seedsLine, -1)

	seedSet := make(map[int]bool)

	// Put all the seeds in a set
	for i := 0; i < len(seeds); {
		seed, _ := strconv.Atoi(seeds[i])
		i++
		seedRange, _ := strconv.Atoi(seeds[i])
		i++
		for j := seed; j <= seed+seedRange; j++ {
			seedSet[j] = true
		}
	}

	levels := make([][]MapEntry, 7)
	mapLevel := -1
	for scanner.Scan() {
		line := scanner.Text()

		// Skip empty lines
		if len(line) == 0 {
			continue
		}

		if line[0] >= '0' && line[0] <= '9' {
			levels[mapLevel] = append(levels[mapLevel], NewMapEntry(line))
		} else {
			mapLevel++
		}
	}

	lowestLocation := math.MaxInt
	for value := range seedSet {
		for i := 0; i < len(levels); i++ {
			for _, level := range levels[i] {
				nextValue, ok := level.getMapped(value)
				if ok {
					value = nextValue
					break
				}

			}
		}

		if value < lowestLocation {
			lowestLocation = value
		}
	}

	println("Lowest Location: ", lowestLocation)

	return nil
}
