package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
)

func main() {
	input, err := readInput()
	if err != nil {
		panic(err)
	}

	memSlice := strings.Split(input, "")

	calculations := findMuls(memSlice)
	answers := performCalcs(calculations)

	fmt.Print("Total: ", sum(answers))
}

func sum(ints []int) int {
	total := 0
	for _, val := range ints {
		total += val
	}
	return total
}

func performCalcs(calculations [][]int) []int {
	var answers []int
	for _, calculation := range calculations {
		answers = append(answers, calculation[0]*calculation[1])
	}
	return answers
}

func findMuls(memSlice []string) [][]int {
	mulStart := [4]string{"m", "u", "l", "("}
	muls := [][]int{}
	for i := 0; i < len(memSlice)-8; i++ {
		partialMem, largerPartial := getNextPartOfMemSlice(memSlice, i)
		fmt.Println()
		fmt.Printf("%s Lpartial: %v", strconv.Itoa(i)+" ", largerPartial)

		// Check start
		if !arraysEqual(mulStart, partialMem) {
			continue
		}
		fmt.Print(" goodStart ")

		//Check num1
		possibleNum1Arr := memSlice[i+4 : i+7]
		fmt.Printf(" possibleNum1: %v", possibleNum1Arr)
		num1, num1Length := getNumberFromStart(possibleNum1Arr)
		fmt.Print(" - num1: ", num1)

		//Check ,
		if memSlice[i+num1Length+4] != "," {
			continue
		}
		fmt.Print(" - found: ','")

		//Check num2
		possibleNum2Arr := memSlice[i+num1Length+5 : i+num1Length+8]
		fmt.Printf(" possibleNum2: %v", possibleNum2Arr)
		num2, num2Length := getNumberFromStart(possibleNum2Arr)
		fmt.Print(" - num2: ", num2)

		//Check )
		if memSlice[i+num1Length+num2Length+5] != ")" {
			continue
		}
		fmt.Print(" - found: ')'")

		muls = append(muls, []int{num1, num2})
	}
	return muls
}

func getNumberFromStart(combinedIntsAndStrings []string) (int, int) {
	var digits []string
	for i := 0; i < len(combinedIntsAndStrings); i++ {
		if len(combinedIntsAndStrings[i]) != 1 {
			panic("string length was not 1 character")
		}
		if unicode.IsDigit(rune(combinedIntsAndStrings[i][0])) {
			digits = append(digits, combinedIntsAndStrings[i])
		} else {
			return getNumFromStringArray(digits), i
		}
	}
	return getNumFromStringArray(digits), len(combinedIntsAndStrings)
}

func getNumFromStringArray(strArr []string) int {
	numStr := ""
	for i := 0; i < len(strArr); i++ {
		numStr += strArr[i]
	}

	num, err := strconv.Atoi(numStr)
	if err != nil {
		panic("Could not convert string array to number at getNumFromStringArray")
	}

	return num
}

func getNextPartOfMemSlice(memSlice []string, startIndex int) ([4]string, [12]string) {
	var partialMem [4]string
	copy(partialMem[:], memSlice[startIndex:startIndex+4])

	var largerPartial [12]string
	lengthLeft := len(memSlice) - startIndex
	if lengthLeft > 12 {
		copy(largerPartial[:], memSlice[startIndex:startIndex+12])
	} else {
		copy(largerPartial[:], memSlice[startIndex:startIndex+lengthLeft])
	}

	return partialMem, largerPartial
}

func arraysEqual(a1 [4]string, a2 [4]string) bool {
	for i := range a1 {
		if a1[i] != a2[i] {
			return false
		}
	}

	return true
}

func readInput() (string, error) {
	data, err := os.ReadFile("input.txt")

	if err != nil {
		return "", fmt.Errorf("File reading error: %w", err)
	}

	content := string(data)

	return content, nil
}
