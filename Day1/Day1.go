package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func main() {
	input, err := os.Open("./input.txt")

	calibrationValues := []int{}

	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		line := scanner.Text()
		value, err := readCalibrationValue(line)
		if err != nil {
			fmt.Printf(err.Error())
		}
		fmt.Printf("Calibration values: %d\n", value)
		calibrationValues = append(calibrationValues, value)
	}
	calibrationResult := calculateCalibrationValues(calibrationValues)
	fmt.Printf("Calibration sum: %d\n", calibrationResult)
}

func calculateCalibrationValues(values []int) int {
	result := 0
	for _, v := range values {
		result += v
	}
	return result
}

func readCalibrationValue(line string) (int, error) {
	replacedString := replaceWordNumbersWithNumbers(line)
	numberRegex := regexp.MustCompile(`\d`)
	numbers := numberRegex.FindAllString(string(replacedString), -1)
	firstNumber := numbers[0]
	secondNumber := numbers[len(numbers)-1]
	return strconv.Atoi(firstNumber + secondNumber)
}

func replaceWordNumbersWithNumbers(line string) string {
	numberNameRegex := regexp.MustCompile(`one|two|three|four|five|six|seven|eight|nine`)
	lineByteArray := numberNameRegex.ReplaceAllFunc([]byte(line), func(b []byte) []byte {
		word := string(b)
		switch word {
		case "one":
			return []byte("1")
		case "two":
			return []byte("2")
		case "three":
			return []byte("3")
		case "four":
			return []byte("4")
		case "five":
			return []byte("5")
		case "six":
			return []byte("6")
		case "seven":
			return []byte("7")
		case "eight":
			return []byte("8")
		case "nine":
			return []byte("9")
		default:
			return []byte("")
		}
	})
	return string(lineByteArray)
}

type Pair struct {
	a, b interface{}
}
