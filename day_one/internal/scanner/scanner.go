package scanner

import (
	"bufio"
	"strconv"
	"strings"
)

func ScanTextToInts(scanner *bufio.Scanner) ([]int, []int) {
	var list_one []int
	var list_two []int
	for scanner.Scan() {
		line := strings.Split(scanner.Text(), "   ")

		value_one, err := strconv.Atoi(line[0])
		if err != nil {
			panic(err)
		}

		value_two, err := strconv.Atoi(line[1])
		if err != nil {
			panic(err)
		}

		list_one = append(list_one, value_one)
		list_two = append(list_two, value_two)

	}
	return list_one, list_two
}
