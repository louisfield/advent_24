package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	"day_three/internal"
	"day_three/internal/scanner"
)

func main() {
	file, err := os.Open("mults")

	if err != nil {
		panic(err)
	}

	lines := scanner.FetchLines(bufio.NewScanner(file))

	var outcome int

	for _, line := range lines {
		mulValues := internal.FetchRegexValues(line, `mul\((\d+),(\d+)\)`)

		for _, value := range mulValues {
			numbersToMult := internal.FetchRegexValues(value, `(\d+)`)
			valueOne, err := strconv.Atoi(numbersToMult[0])
			if err != nil {
				panic(err)
			}
			valueTwo, err := strconv.Atoi(numbersToMult[1])
			if err != nil {
				panic(err)
			}

			outcome += valueOne * valueTwo

		}
	}

	fmt.Println(outcome)

	var outcometwo int
	var performMult bool = true
	for _, line := range lines {
		mulValues := internal.FetchRegexValues(line, `mul\((\d+),(\d+)\)|do\(\)|don't\(\)`)

		for _, value := range mulValues {

			if value == "do()" {
				performMult = true
			} else if value == "don't()" {
				performMult = false
			} else if performMult {
				numbersToMult := internal.FetchRegexValues(value, `(\d+)`)
				valueOne, err := strconv.Atoi(numbersToMult[0])
				if err != nil {
					panic(err)
				}
				valueTwo, err := strconv.Atoi(numbersToMult[1])
				if err != nil {
					panic(err)
				}

				outcometwo += valueOne * valueTwo
			}

		}
	}

	fmt.Println(outcometwo)

}
