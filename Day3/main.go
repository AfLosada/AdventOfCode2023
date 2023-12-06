package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"slices"
	"strconv"
	"strings"
)

func main() {
	input, err := os.Open("./input.txt")
	adjacentNumbers := make([]int, 0)
	gearRatioList := make([]int, 0)

	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(input)
	schematicsMatrix := make([][]string, 0)
	visitedMatrix := make([][]bool, 0)
	for scanner.Scan() {
		line := scanner.Text()
		row := strings.Split(line, "")
		schematicsMatrix = append(schematicsMatrix, row)
		visitedMatrix = append(visitedMatrix, make([]bool, len(row)))
	}
	for i, _ := range schematicsMatrix {
		for j := range schematicsMatrix[i] {
			point := schematicsMatrix[i][j]
			nonNumberRegex := regexp.MustCompile(`[^\d.]`)
			if nonNumberRegex.Match([]byte(point)) {
				adjacentNumbersToPoint, err := findAdjacentNumbersToPoint(Position{x: i, y: j}, schematicsMatrix, visitedMatrix)
				if err != nil {
					fmt.Println(err.Error())
					break
				}
				if point == "*" && len(adjacentNumbersToPoint) == 2 {
					gearRatioList = append(gearRatioList, adjacentNumbersToPoint[0]*adjacentNumbersToPoint[1])
				}
				adjacentNumbers = append(adjacentNumbers, adjacentNumbersToPoint...)
			}
		}
	}
	partNumberSum := 0
	for _, partNumber := range adjacentNumbers {
		partNumberSum += partNumber
	}
	gearRatioSum := 0
	for _, gearRatio := range gearRatioList {
		gearRatioSum += gearRatio
	}
	fmt.Printf("The sum of part numbers is: %d\n", partNumberSum)
	fmt.Printf("The sum of gear ratios is: %d\n", gearRatioSum)
}

// TODO: I CAN USE A SLICE OF THE POINT IN EACH DIRECTION AND CHECK IF IT HAS DIGITS WITH A REGEX
func findAdjacentNumbersToPoint(position Position, schematicsMatrix [][]string, visitedMatrix [][]bool) ([]int, error) {
	x := position.x
	y := position.y
	hasTop := y > 0
	hasLeft := x > 0
	hasBottom := x < len(schematicsMatrix)
	hasRight := y < len(schematicsMatrix[x])
	numberRegex := regexp.MustCompile(`[\d]`)
	adjacentNumbersPositions := make([]Position, 0)
	if hasTop {
		currentPosition := Position{x, y - 1}
		value := schematicsMatrix[currentPosition.x][currentPosition.y]
		if numberRegex.Match([]byte(value)) {
			adjacentNumbersPositions = append(adjacentNumbersPositions, currentPosition)
		}
		if hasLeft {
			currentPosition = Position{x - 1, y - 1}
			value := schematicsMatrix[currentPosition.x][currentPosition.y]
			if numberRegex.Match([]byte(value)) {
				adjacentNumbersPositions = append(adjacentNumbersPositions, currentPosition)
			}
		}
		if hasRight {
			currentPosition = Position{x + 1, y - 1}
			value := schematicsMatrix[currentPosition.x][currentPosition.y]
			if numberRegex.Match([]byte(value)) {
				adjacentNumbersPositions = append(adjacentNumbersPositions, currentPosition)
			}
		}
	}
	if hasBottom {
		currentPosition := Position{x, y + 1}
		value := schematicsMatrix[currentPosition.x][currentPosition.y]
		if numberRegex.Match([]byte(value)) {
			adjacentNumbersPositions = append(adjacentNumbersPositions, currentPosition)
		}
		if hasLeft {
			currentPosition = Position{x - 1, y + 1}
			value := schematicsMatrix[currentPosition.x][currentPosition.y]
			if numberRegex.Match([]byte(value)) {
				adjacentNumbersPositions = append(adjacentNumbersPositions, currentPosition)
			}
		}
		if hasRight {
			currentPosition = Position{x + 1, y + 1}
			value := schematicsMatrix[currentPosition.x][currentPosition.y]
			if numberRegex.Match([]byte(value)) {
				adjacentNumbersPositions = append(adjacentNumbersPositions, currentPosition)
			}
		}
	}
	if hasLeft {
		currentPosition := Position{x - 1, y}
		value := schematicsMatrix[currentPosition.x][currentPosition.y]
		if numberRegex.Match([]byte(value)) {
			adjacentNumbersPositions = append(adjacentNumbersPositions, currentPosition)
		}
	}
	if hasRight {
		currentPosition := Position{x + 1, y}
		value := schematicsMatrix[currentPosition.x][currentPosition.y]
		if numberRegex.Match([]byte(value)) {
			adjacentNumbersPositions = append(adjacentNumbersPositions, currentPosition)
		}
	}
	adjacentNumbers := make([]int, 0)
	for _, position := range adjacentNumbersPositions {
		wholeNumber, err := findWholeNumberFromPosition(position, schematicsMatrix, visitedMatrix)
		if err != nil {
			return nil, err
		}
		if wholeNumber != 0 {
			adjacentNumbers = append(adjacentNumbers, wholeNumber)
		}
	}
	return adjacentNumbers, nil
}

func findWholeNumberFromPosition(position Position, schematicsMatrix [][]string, visitedMatrix [][]bool) (int, error) {
	i := position.x
	j := position.y
	if visitedMatrix[i][j] {
		return 0, nil
	}
	row := schematicsMatrix[i]
	currentNumber := schematicsMatrix[i][j]
	nonNumberRegex := regexp.MustCompile(`[^\d]`)
	leftNumber := ""
	for j := position.y - 1; j >= 0; j-- {
		currentLeftNumber := row[j]
		if visitedMatrix[i][j] {
			break
		}
		if nonNumberRegex.Match([]byte(currentLeftNumber)) {
			break
		}
		leftNumber += currentLeftNumber
		visitedMatrix[i][j] = true
	}
	splitLeftNumber := strings.Split(leftNumber, "")
	slices.Reverse(splitLeftNumber)
	leftNumber = strings.Join(splitLeftNumber, "")
	rightNumber := ""
	for j := position.y + 1; j < len(row); j++ {
		currentRightNumber := row[j]
		if visitedMatrix[i][j] {
			break
		}
		if nonNumberRegex.Match([]byte(currentRightNumber)) {
			break
		}
		rightNumber += currentRightNumber
		visitedMatrix[i][j] = true
	}
	visitedMatrix[i][j] = true
	return strconv.Atoi(leftNumber + currentNumber + rightNumber)
}

type Position struct {
	x int
	y int
}
