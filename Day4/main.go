package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	input, err := os.Open("./input.txt")
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(input)
	cardValues := make([][]int, 0)
	winningCards := make([][]int, 0)
	i := 0
	for scanner.Scan() {
		line := scanner.Text()
		splitLine := strings.Split(line, "|")
		winningPart := splitLine[0]
		splitWinningPart := strings.Split(winningPart, ":")
		cardName := splitWinningPart[0]
		winningNumbers := strings.Fields(splitWinningPart[1])
		numbersYouHave := strings.Fields(splitLine[1])
		winnerMap := make(map[string]bool, 0)
		for _, number := range winningNumbers {
			winnerMap[number] = true
		}
		numbersYouWin := make([]string, 0)
		for _, number := range numbersYouHave {
			if _, ok := winnerMap[number]; ok {
				numbersYouWin = append(numbersYouWin, number)
			}
		}
		cardValue := calculateCardValue(numbersYouWin, 1)
		cardValues = append(cardValues, []int{cardValue})
		winningCards = append(winningCards, []int{len(numbersYouWin)})
		fmt.Printf("The total value of the card %s is: %d\n", cardName, cardValue)
		i++
	}
	for key, valueArray := range winningCards {
		for _, value := range valueArray {
			for j := 0; j < value; j++ {
				positionToDuplicate := key + j + 1
				if positionToDuplicate >= len(winningCards) {
					break
				}
				if len(winningCards[positionToDuplicate]) > 0 {
					winningCards[positionToDuplicate] = append(winningCards[positionToDuplicate], winningCards[positionToDuplicate][0])
				}
			}
		}
	}

	fmt.Printf("The total value of the cards is: %d\n", sumValues(cardValues))
	fmt.Printf("The total amount of cards of the cards is: %d\n", sumLength(winningCards))
}

func sumLength(arrayOfArrays [][]int) int {
	result := 0
	for _, array := range arrayOfArrays {
		result += len(array)
	}
	return result
}

func sumValues(arrayOfArrays [][]int) int {
	result := 0
	for _, array := range arrayOfArrays {
		for _, number := range array {
			result += number
		}
	}
	return result
}

func calculateCardValue(numbers []string, currentValue int) int {
	if len(numbers) == 0 {
		return 0
	}
	if len(numbers) == 1 {
		return currentValue
	}
	return calculateCardValue(numbers[:len(numbers)-1], currentValue*2)
}
