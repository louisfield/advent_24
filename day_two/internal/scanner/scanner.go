package scanner

import (
	"bufio"
	"strconv"
	"strings"
)

func ScanTextToInts(scanner *bufio.Scanner) [][]int {
	var outputList [][]int
	for scanner.Scan() {

		var convertedLine []int
		line := strings.Split(scanner.Text(), " ")

		for _, value := range line {

			convertedValue, err := strconv.Atoi(value)
			if err != nil {
				panic(err)
			}

			convertedLine = append(convertedLine, convertedValue)

		}

		outputList = append(outputList, convertedLine)

	}
	return outputList
}

func FindSafeReports(list [][]int) int {
	var safeReports int
	for _, line := range list {

		var prevNumber int = -1
		var ascending bool = true
		var descending bool = true
		var isSafe bool = true
		for _, num := range line {
			if prevNumber != -1 {
				ascending = num > prevNumber && ascending
				descending = num < prevNumber && descending

				isSafe = checkSafe(num, prevNumber, ascending, descending)
			}

			prevNumber = num
			if !isSafe {
				break
			}
		}

		if isSafe {
			safeReports++
		}
	}
	return safeReports
}

func checkSafe(num int, prevNumber int, descending bool, ascending bool) bool {
	if abs(num-prevNumber) > 3 || abs(num-prevNumber) < 1 {
		return false
	}
	return descending || ascending
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
