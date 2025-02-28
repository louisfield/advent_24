package scanner

import (
	"bufio"
)

func FetchLines(scanner *bufio.Scanner) []string {
	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines
}
