package main

import (
	"bufio"
	"day_one/internal/scanner"
	"fmt"
	"log"
	"os"
	"sort"
)

func main() {
	file, err := os.Open("location_ids")

	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	bufioScanner := bufio.NewScanner(file)

	listOne, listTwo := scanner.ScanTextToInts(bufioScanner)

	var acc int
	var acc2 int

	uniqueValuesListTwo := createUniqueValueDict(listTwo)

	sort.Ints(listOne)
	sort.Ints(listTwo)

	for i, num := range listOne {
		acc2 += num * uniqueValuesListTwo[listOne[i]]
		acc += abs(listOne[i] - listTwo[i])
	}
	fmt.Println("The sum of the absolute differences between the two lists is: ", acc)

	fmt.Println(acc2)

}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func createUniqueValueDict(list []int) map[int]int {
	uniqueValues := make(map[int]int)

	for _, num := range list {
		uniqueValues[num] = uniqueValues[num] + 1
	}

	return uniqueValues
}
