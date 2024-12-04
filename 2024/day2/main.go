package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	input, err := readInput()
	if err != nil {
		panic(err)
	}
	reports := createReportsFromInput(input)

	safeReports := 0

	for _, report := range reports {
		if countUnsafeSituationsInReport(report) == 0 {
			safeReports++
			fmt.Print("safe   ")
		} else {
			fmt.Print("unsafe ")
		}
		fmt.Printf("report: %v", report)
		fmt.Println()
	}

	fmt.Println("Safe reports: ", safeReports)
}

func countUnsafeSituationsInReport(report []int) int {
	unsafeSituations := 0
	var lastLevel int
	var lastDifWasPositive bool
	var isLastDifInitialized bool
	for i, level := range report {
		if i == 0 {
			//First level
			lastLevel = level
			continue
		}

		// Increase/Decrease fault checking
		dif := level - lastLevel
		if dif == 0 {
			unsafeSituations++
			lastLevel = level
			continue
		}

		if !isLastDifInitialized {
			//Initialize lastDifWasPositive
			lastDifWasPositive = dif > 0
			isLastDifInitialized = true

			// Check dif is within bounds (1-3)
			absDif := absolute(dif)
			if absDif < 1 || absDif > 3 {
				unsafeSituations++
			}

			lastLevel = level
			continue
		}

		if (dif < 0 && lastDifWasPositive) || (dif > 0 && !lastDifWasPositive) {
			// opposite increase/decrease from last
			unsafeSituations++
			lastLevel = level
			continue
		}

		// Check dif is within bounds (1-3)
		absDif := absolute(dif)
		if absDif < 1 || absDif > 3 {
			unsafeSituations++
			lastLevel = level
			continue
		}

		lastLevel = level
	}
	return unsafeSituations
}

func absolute(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func createReportsFromInput(input string) [][]int {
	reportAsStrings := strings.Split(input, "\r\n")

	reports := [][]int{}

	for i := 0; i < len(reportAsStrings); i++ {
		levelsStr := strings.Split(reportAsStrings[i], " ")
		var levels []int
		for _, levelAsString := range levelsStr {
			level, err := strconv.Atoi(levelAsString)
			if err != nil {
				fmt.Errorf("Could not parse level as string to an integer: ", err)
			}
			levels = append(levels, level)
		}
		reports = append(reports, levels)
	}

	return reports
}

func readInput() (string, error) {
	data, err := os.ReadFile("input.txt")

	if err != nil {
		return "", fmt.Errorf("File reading error: %w", err)
	}

	content := string(data)

	return content, nil
}
