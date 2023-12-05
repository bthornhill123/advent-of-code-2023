package advent

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type One struct{}

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
		last := d.firstDigit(reverse(line))
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

func reverse(s string) string {
	// Convert string to a slice of runes
	r := []rune(s)

	// Reverse the slice
	for i, j := 0, len(r)-1; i < j; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}

	// Convert the slice back to a string
	return string(r)
}

// PartTwo considers both digit ("1") and word ("one") as valid calibration values
func (d One) PartTwo() error {
	stringToDigitMap := map[string]string{
		"one":   "1",
		"two":   "2",
		"three": "3",
		"four":  "4",
		"five":  "5",
		"six":   "6",
		"seven": "7",
		"eight": "8",
		"nine":  "9",
	}

	trie := NewTrie()
	for key := range stringToDigitMap {
		trie.Add(key)
	}

	file, err := os.Open("inputs/one.txt")
	if err != nil {
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	sum := 0
	for scanner.Scan() {
		line := scanner.Text()

		firstDigit := d.firstDigitString(line, trie, stringToDigitMap)
		lastDigit := d.lastDigitString(line, trie, stringToDigitMap)
		number, _ := strconv.Atoi(firstDigit + lastDigit)
		sum += number
	}

	fmt.Println("Sum: ", sum)

	return nil
}

type TrieNode struct {
	children map[rune]*TrieNode
	isEnd    bool
}

func NewTrieNode() *TrieNode {
	return &TrieNode{
		children: make(map[rune]*TrieNode),
		isEnd:    false,
	}
}

type Trie struct {
	root *TrieNode
}

func NewTrie() *Trie {
	return &Trie{
		root: NewTrieNode(),
	}
}

func (t *Trie) Add(word string) {
	node := t.root
	for _, char := range word {
		if _, ok := node.children[char]; !ok {
			node.children[char] = NewTrieNode()
		}
		node = node.children[char]
	}
	node.isEnd = true
}

func (t *Trie) StartsWithMatch(searchString string) (string, bool) {
	node := t.root
	for i, char := range searchString {
		if _, ok := node.children[char]; !ok {
			return "", false
		}

		node = node.children[char]

		if node.isEnd {
			return searchString[:i+1], true
		}
	}

	return "", false
}

func (d One) firstDigitString(line string, trie *Trie, stringToDigitMap map[string]string) string {
	for i := 0; i < len(line); i++ {
		char := string(line[i])
		if char >= "0" && char <= "9" {
			return char
		}
		substring := line[i:]
		match, ok := trie.StartsWithMatch(substring)
		if ok {
			return stringToDigitMap[match]
		}
	}

	panic("Could not find string digit in line")
}

func (d One) lastDigitString(line string, trie *Trie, stringToDigitMap map[string]string) string {
	lastMatch := ""
	for i := 0; i < len(line); i++ {
		char := string(line[i])
		if char >= "0" && char <= "9" {
			lastMatch = char
			continue
		}

		substring := line[i:]
		match, ok := trie.StartsWithMatch(substring)
		if ok {
			lastMatch = stringToDigitMap[match]
		}
	}

	return lastMatch
}

// ------------
// PartTwoViaReplacement doesn't work due to the fact that number strings overlap
func (d One) PartTwoViaReplacement() error {
	wordToDigitMap := map[string]string{
		"one":   "1",
		"two":   "2",
		"three": "3",
		"four":  "4",
		"five":  "5",
		"six":   "6",
		"seven": "7",
		"eight": "8",
		"nine":  "9",
	}

	file, err := os.Open("inputs/one.txt")
	if err != nil {
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	sum := 0
	for scanner.Scan() {
		line := scanner.Text()

		replacedLeftToRight := d.replaceLeftToRight(line, wordToDigitMap)
		first := d.firstDigit(replacedLeftToRight)

		replacedRightToLeft := d.replaceRightToLeft(line, wordToDigitMap)
		last := d.firstDigit(replacedRightToLeft)
		number, _ := strconv.Atoi(first + last)
		sum += number
	}

	fmt.Println("Sum: ", sum)

	return nil
}

func (d One) replaceLeftToRight(str string, wordToDigitMap map[string]string) string {
	for word, digit := range wordToDigitMap {
		str = strings.Replace(str, word, word+digit+word, -1)
	}
	return str
}

func (d One) replaceRightToLeft(str string, wordToDigitMap map[string]string) string {
	str = reverse(str)
	for word, digit := range wordToDigitMap {
		str = strings.Replace(str, reverse(word), reverse(word)+digit+reverse(word), -1)
	}
	return reverse(str)
}
