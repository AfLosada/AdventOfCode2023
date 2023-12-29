package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

func timer(name string) func() {
	start := time.Now()
	return func() {
		fmt.Printf("%s took %v\n", name, time.Since(start))
	}
}

func main() {
	defer timer("main")()
	input, err := os.Open("./input.txt")
	if err != nil {
		panic(err)
	}
	lines, err := readLines(input)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	nextNumbers := []int{}
	firstNumbers := []int{}
	total := 0
	totalFirst := 0
	for _, line := range lines {
		nextNumber := calculateNextNumber(line)
		firstNumber := calculateFirstNumber(line)
		nextNumbers = append(nextNumbers, nextNumber)
		firstNumbers = append(firstNumbers, firstNumber)
		total += nextNumber
		totalFirst += firstNumber
	}
	fmt.Printf("The total sum of the next numbers: %d\n", total)
	fmt.Printf("The total sum of the first numbers: %d\n", totalFirst)
}

func readLines(input *os.File) ([][]int, error) {
	scanner := bufio.NewScanner(input)
	lines := [][]int{}
	for scanner.Scan() {
		line := scanner.Text()
		numbersStrings := strings.Fields(line)
		numbers := []int{}
		for _, str := range numbersStrings {
			number, err := strconv.Atoi(str)
			if err != nil {
				return nil, err
			}
			numbers = append(numbers, number)
		}
		lines = append(lines, numbers)
	}
	return lines, nil
}

func calculateNextNumber(line []int) int {
	if len(line) == 1 {
		return line[len(line)-1]
	}
	if allAreZeroes(line) {
		return line[len(line)-1]
	}
	newLine := make([]int, len(line)-1)
	for i := range newLine {
		newLine[i] = int(float64(line[i+1] - line[i]))
	}
	lastNumber := line[len(line)-1]
	nextNumber := calculateNextNumber(newLine)
	return lastNumber + nextNumber
}

func calculateFirstNumber(line []int) int {
	if len(line) == 1 {
		return line[0]
	}
	if allAreZeroes(line) {
		return line[0]
	}
	newLine := make([]int, len(line)-1)
	for i := range newLine {
		newLine[i] = int(float64(line[i+1] - line[i]))
	}
	lastNumber := line[0]
	fmt.Printf("Len: %d. %v\n", len(line), line)
	nextNumber := calculateFirstNumber(newLine)
	return lastNumber - nextNumber
}

func allAreZeroes(line []int) bool {
	for _, v := range line {
		if v != 0 {
			return false
		}
	}
	return true
}
