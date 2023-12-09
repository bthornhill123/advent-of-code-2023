package advent

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type Three struct{}

func (d Three) PartOne() error {
	file, err := os.Open("inputs/three.txt")
	if err != nil {
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	// Fill matrix
	var matrix [][]rune
	for scanner.Scan() {
		line := scanner.Text()
		matrix = append(matrix, []rune(line))
	}

	sum := 0
	for rowIdx, row := range matrix {
		for colIdx := 0; colIdx < len(row); colIdx++ {
			char := row[colIdx]
			if d.isDigit(char) {
				startIndex := colIdx
				colIdx++
				numString := string(char)
				for ; colIdx < len(row); colIdx++ {
					char = row[colIdx]
					if !d.isDigit(char) {
						break
					}
					numString += string(char)
				}
				endIndex := colIdx - 1
				_, ok := d.isAdjacent(rowIdx, startIndex-1, endIndex+1, matrix, d.isSymbol)
				if ok {
					num, _ := strconv.Atoi(numString)
					sum += num
				}
			}
		}
	}

	println("Sum: ", sum)

	return nil
}

func (d Three) isDigit(c rune) bool {
	return c >= '0' && c <= '9'
}

func (d Three) isSymbol(c rune) bool {
	return !d.isDigit(c) && c != '.'
}

func (d Three) isGear(c rune) bool {
	return c == '*'
}

func (d Three) maxInt(x, y int) int {
	if x > y {
		return x
	}

	return y
}

func (d Three) minInt(x, y int) int {
	if x < y {
		return x
	}

	return y
}

func (d Three) isAdjacent(row, colStart, colEnd int, matrix [][]rune, isMatch func(rune) bool) ([]string, bool) {
	// Ensure everything is in bounds
	ids := make([]string, 0) // row + col as a string, like "10"
	ok := false
	colStart = d.maxInt(colStart, 0)
	colEnd = d.minInt(colEnd, len(matrix[row])-1)

	// Search row above from colStart to colEnd
	if row > 0 {
		for i := colStart; i <= colEnd; i++ {
			if isMatch(matrix[row-1][i]) {
				ids = append(ids, d.gearID(row-1, i))
				ok = true
			}
		}
	}

	// Search row at colStart and colEnd
	if isMatch(matrix[row][colStart]) {
		ids = append(ids, d.gearID(row, colStart))
		ok = true
	}
	if isMatch(matrix[row][colEnd]) {
		ids = append(ids, d.gearID(row, colEnd))
		ok = true
	}

	// Search row below from colStart to colEnd
	if row < len(matrix)-1 {
		for i := colStart; i <= colEnd; i++ {
			if isMatch(matrix[row+1][i]) {
				ids = append(ids, d.gearID(row+1, i))
				ok = true
			}
		}
	}

	return ids, ok
}

func (d Three) gearID(row, col int) string {
	return fmt.Sprintf("row:%v col:%v", row, col)
}

func (d Three) PartTwo() error {
	// row -> col -> partNumbers
	gears := make(map[string][]int, 0)

	file, err := os.Open("inputs/three.txt")
	if err != nil {
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	// Fill matrix
	var matrix [][]rune
	for scanner.Scan() {
		line := scanner.Text()
		matrix = append(matrix, []rune(line))
	}

	for rowIdx, row := range matrix {
		for colIdx := 0; colIdx < len(row); colIdx++ {
			char := row[colIdx]
			if d.isDigit(char) {
				startIndex := colIdx
				colIdx++
				numString := string(char)
				for ; colIdx < len(row); colIdx++ {
					char = row[colIdx]
					if !d.isDigit(char) {
						break
					}
					numString += string(char)
				}
				colIdx--

				gearIDs, ok := d.isAdjacent(rowIdx, startIndex-1, colIdx+1, matrix, d.isGear)
				if ok {
					num, _ := strconv.Atoi(numString)
					for _, gearID := range gearIDs {
						gears[gearID] = append(gears[gearID], num)
					}
				}
			}
		}
	}

	sum := 0
	for _, values := range gears {
		if len(values) != 2 {
			continue
		}

		gearRatio := values[0] * values[1]
		sum += gearRatio
	}

	fmt.Println("Sum: ", sum)

	return nil
}
