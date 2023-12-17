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
	if err != nil {
		panic(err)
	}
	times, distancesToBeat := readInput(input)
	maxTime, err := strconv.Atoi(times)
	if err != nil {
		fmt.Print(err.Error())
	}
	distanceToBeat, err := strconv.Atoi(distancesToBeat)
	if err != nil {
		fmt.Print(err.Error())
	}
	numberOfWaysToBeat := calculateNumberOfWaysToBeat(maxTime, distanceToBeat)
	fmt.Printf("The multiplication of these numbers is %d\n", numberOfWaysToBeat)
}

func readInput(input *os.File) (string, string) {
	scanner := bufio.NewScanner(input)
	scanner.Scan()
	timeLine := scanner.Text()
	scanner.Scan()
	distanceLine := scanner.Text()
	times := strings.Join(strings.Fields(strings.Split(timeLine, ":")[1]), "")
	distances := strings.Join(strings.Fields(strings.Split(distanceLine, ":")[1]), "")
	return times, distances
}

func calculateNumberOfWaysToBeat(maxTime int, maxDistance int) int {
	numberOfWaysToBeat := 0
	maxTimeToPress := maxTime / 2
	for i := maxTimeToPress; i >= 0 && calculateDistance(i, maxTime) > maxDistance; i-- {
		numberOfWaysToBeat++
	}
	for i := maxTimeToPress + 1; i < maxTime && calculateDistance(i, maxTime) > maxDistance; i++ {
		numberOfWaysToBeat++
	}
	return numberOfWaysToBeat
}

func calculateDistance(timePressing int, totalTime int) int {
	return timePressing * (totalTime - timePressing)
}
