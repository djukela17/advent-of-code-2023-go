package puzzles

import (
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
