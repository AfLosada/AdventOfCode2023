package main

import (
	"bufio"
	"os"
	"regexp"
	"strings"
)

func main() {
	input, err := os.Open("./input.txt")

	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(input)
	schematicsMatrix := make([][]string, 0)
	adjacencyMatrix := make([][]string, 0)
	for scanner.Scan() {
		line := scanner.Text()
		row := strings.Split(line, "")
		schematicsMatrix = append(schematicsMatrix, row)
		adjacencyMatrix = append(adjacencyMatrix, make([]string, len(row)))
	}
}
//TODO: I CAN USE A SLICE OF THE POINT IN EACH DIRECTION AND CHECK IF IT HAS DIGITS WITH A REGEX
func findAdjacentNumbersToPoint(x int, y int, schematicsMatrix [][]string, adjacencyMatrix [][]string) {
	var topLeft string
	var topRight string
	var bottomLeft string
	var bottomRight string
	var left string
	var right string
	var top string
	var bottom string
	hasTop := y > 1
	hasLeft := x > 1
	hasBottom := x < len(schematicsMatrix)
	hasRight := y < len(schematicsMatrix[x])
	numberRegex := regexp.MustCompile(`\d`)
	adjacentNumbersSlots := make([]string, 0)
	if hasTop {
		top = schematicsMatrix[x][y-1]
		if (numberRegex.Match([]byte(top))){
			adjacentNumbersSlots = append(adjacentNumbersSlots, )
		}
		if hasLeft {
			topLeft = schematicsMatrix[x-1][y-1]
		}
		if hasRight {
			topLeft = schematicsMatrix[x+1][y-1]
		}
	}
	if hasBottom {
		bottom = schematicsMatrix[x][y+1]
		if hasLeft {
			bottom = schematicsMatrix[x-1][y+1]
		}
		if hasRight {
			bottom = schematicsMatrix[x+1][y+1]
		}
	}
	if hasLeft {
		left = schematicsMatrix[x-1][y]
	}
	if hasRight {
		left = schematicsMatrix[x+1][y]
	}
	for _, slot := range adjacentSlots {
		if numberRegex.Match([]byte(slot)) {
			adjacentNumbers = append(adjacentNumbers, slot)
		}
	}
	return adjacentNumbers
}
