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

	fmt.Println("Safe reports: ", countSafeReports(reports))
}

func countSafeReports(reports [][]int) int {
	safeReports := 0

	for _, report := range reports {
		if isSafeSituation(report) {
			safeReports++
			printReport(report, "safe   ")
		} else if isDampenedSafe(report) {
			safeReports++
			printReport(report, "sa damp") //safe dampened
		} else {
			printReport(report, "unsafe ")
		}
	}

	return safeReports
}

func isDampenedSafe(report []int) bool {
	dampenerSuggestions := getDampenerSuggestions(report)
	for i := 0; i < len(dampenerSuggestions); i++ {
		isDampenedSafe := isSafeSituation(dampenerSuggestions[i])
		if isDampenedSafe {

			return true
		}
	}
	return false
}

func getDampenerSuggestions(report []int) [][]int {
	var dampenerSuggestions [][]int

	for i := 0; i < len(report); i++ {
		dampenerSuggestions = append(dampenerSuggestions, removeIndex(report, i))
	}

	return dampenerSuggestions
}

func isSafeSituation(report []int) bool {
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
			return false
		}

		if !isLastDifInitialized {
			//Initialize lastDifWasPositive
			lastDifWasPositive = dif > 0
			isLastDifInitialized = true

			// Check dif is within bounds (1-3)
			absDif := absolute(dif)
			if absDif < 1 || absDif > 3 {
				return false
			}
		}

		if (dif < 0 && lastDifWasPositive) || (dif > 0 && !lastDifWasPositive) {
			// opposite increase/decrease from last
			return false
		}

		// Check dif is within bounds (1-3)
		absDif := absolute(dif)
		if absDif < 1 || absDif > 3 {
			return false
		}

		lastLevel = level
	}
	return true
}

func printReport(report []int, state string) {
	fmt.Print(state + " ")
	fmt.Printf("report: %v", report)
	fmt.Println()
}

func absolute(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func removeIndex(slice []int, index int) []int {
	ret := make([]int, 0)
	ret = append(ret, slice[:index]...)
	return append(ret, slice[index+1:]...)
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
