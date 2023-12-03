package puzzles

import (
	"math"
	"strings"
)

type TrebuchetPuzzle struct{}

func (p *TrebuchetPuzzle) SolvePartOne(input string) int {
	const None = -1

	total := 0

	for _, line := range strings.Split(input, "\n") {
		//for _, line := range bytes.Split(contents, []byte("\n")) {
		// skip empty line, usually the last line in the input file
		if len(line) == 0 {
			continue
		}

		firstDigit := None
		lastDigit := None

		for _, c := range line {
			if c > 47 && c < 58 {
				n := int(c - 48)

				if firstDigit == None {
					firstDigit = n
				}

				lastDigit = n
			}
		}

		num := firstDigit*10 + lastDigit

		total += num
	}

	return total
}

func (p *TrebuchetPuzzle) SolvePartTwo(input string) int {
	const (
		One   = "one"
		Two   = "two"
		Three = "three"
		Four  = "four"
		Five  = "five"
		Six   = "six"
		Seven = "seven"
		Eight = "eight"
		Nine  = "nine"
	)

	total := 0

	for _, line := range strings.Split(input, "\n") {
		// find the first digit
		firstDigit := 0
		buff := ""

		for _, c := range line {
			if c > 47 && c < 58 {
				firstDigit = int(c - 48)

				break
			}

			buff += string(c)
		}

		if len(buff) > 2 {
			positions := []int{
				strings.Index(buff, One),
				strings.Index(buff, Two),
				strings.Index(buff, Three),
				strings.Index(buff, Four),
				strings.Index(buff, Five),
				strings.Index(buff, Six),
				strings.Index(buff, Seven),
				strings.Index(buff, Eight),
				strings.Index(buff, Nine),
			}

			lowestPositionValue := math.MaxInt

			for i, pos := range positions {
				if pos > -1 && pos < lowestPositionValue {
					lowestPositionValue = pos
					firstDigit = i + 1
				}
			}
		}

		// find the last digit
		lastDigit := 0
		buff = ""

		for i := len(line) - 1; i >= 0; i-- {
			c := line[i]
			if c > 47 && c < 58 {
				lastDigit = int(c - 48)

				break
			}

			buff = string(c) + buff
		}

		if len(buff) > 2 {
			positions := []int{
				strings.LastIndex(buff, One),
				strings.LastIndex(buff, Two),
				strings.LastIndex(buff, Three),
				strings.LastIndex(buff, Four),
				strings.LastIndex(buff, Five),
				strings.LastIndex(buff, Six),
				strings.LastIndex(buff, Seven),
				strings.LastIndex(buff, Eight),
				strings.LastIndex(buff, Nine),
			}

			highestPositionValue := -1

			for i, pos := range positions {
				if pos > highestPositionValue {
					highestPositionValue = pos
					lastDigit = i + 1
				}
			}
		}

		total += firstDigit*10 + lastDigit
	}

	return total
}
