package advent

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type One struct {
}

func (d One) PartOne() error {
	file, err := os.Open("inputs/one.txt")
	if err != nil {
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	sum := 0
	for scanner.Scan() {
		line := scanner.Text()
		first := d.firstDigit(line)
		last := d.firstDigit(reverseString(line))
		number, _ := strconv.Atoi(first + last)
		sum += number
	}

	fmt.Println("Sum: ", sum)

	return nil
}

func (d One) firstDigit(s string) string {
	for _, c := range s {
		// Check if the character is a numeric digit (0-9)
		if c >= '0' && c <= '9' {
			return string(c)
		}
	}

	panic("No digit found in string")
}

func reverseString(s string) string {
	// Convert string to a slice of runes
	r := []rune(s)

	// Reverse the slice
	for i, j := 0, len(r)-1; i < j; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}

	// Convert the slice back to a string
	return string(r)
}
