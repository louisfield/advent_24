package main

import (
	"bufio"
	"day_two/internal/scanner"
	"fmt"
	"log"
	"os"
)

func main() {
	file, err := os.Open("reports.txt")

	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	bufioScanner := bufio.NewScanner(file)

	outputList := scanner.ScanTextToInts(bufioScanner)

	safeReports := scanner.FindSafeReports(outputList)
	fmt.Print(safeReports)

}
