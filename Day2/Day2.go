package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	input, err := os.Open("./input.txt")
	elfBag := map[string]int{"red": 12, "green": 13, "blue": 14}

	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(input)
	possibleGames := make([]string, 0)
	var idSum int
	for scanner.Scan() {
		gameLine := scanner.Text()
		gameId, boxCounterDictPerPlay, err := parseGameLine(gameLine)
		if err != nil {
			fmt.Printf(err.Error())
		}
		if validateBoxCounterArray(boxCounterDictPerPlay, elfBag) {
			possibleGames = append(possibleGames, gameId)
			gameIdNumber, err := strconv.Atoi(gameId)
			if err != nil {
				fmt.Printf(err.Error())
			}
			idSum += gameIdNumber
		}
	}
	fmt.Printf("The sum of the ids of the games is: %d\n", idSum)
}

func validateBoxCounterArray(boxCounterArray []map[string]int, elfBag map[string]int) bool {
	for _, boxCounter := range boxCounterArray {
		for key, value := range boxCounter {
			if value > elfBag[key] {
				return false
			}
		}
	}
	return true
}

func parseGameLine(line string) (string, []map[string]int, error) {
	splitGameLine := strings.Split(line, ":")
	splitGamePlays := strings.Split(splitGameLine[1], ";")
	subsetBoxPlayDict := make([]map[string]int, 0)
	for i, play := range splitGamePlays {
		playsByBox := strings.Split(play, ",")
		subsetBoxPlayDict = append(subsetBoxPlayDict, make(map[string]int))
		for _, playByBox := range playsByBox {
			playAndBox := strings.Split(strings.Trim(playByBox, " "), " ")
			play, err := strconv.Atoi(playAndBox[0])
			if err != nil {
				return "", nil, err
			}
			box := playAndBox[1]
			boxPlayDict := subsetBoxPlayDict[i]
			if boxPlayDict == nil {
				boxPlayDict = make(map[string]int)
			}
			boxPlayDict[box] += play
			subsetBoxPlayDict[i] = boxPlayDict
		}
	}
	gameName := splitGameLine[0]
	splitGameName := strings.Split(gameName, " ")
	return splitGameName[1], subsetBoxPlayDict, nil
}
