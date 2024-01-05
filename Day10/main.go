package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
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
	input, err := os.Open("./smol_input.txt")
	if err != nil {
		panic(err)
	}
	nodeMatrix, edgeMap, startNode := readLines(input)
	inferredStartNode := Node{position: startNode.position, form: inferFormForStart(startNode.position, nodeMatrix)}
	nodeMatrix[inferredStartNode.position.x][inferredStartNode.position.y] = inferredStartNode
	edgeMap[inferredStartNode.position] = inferredStartNode
	path := findAnimalPath(nodeMatrix, edgeMap, inferredStartNode)
	divisor := int(math.Ceil(float64(len(path)) / 2.0))
	leftPath := path[0:divisor]
	rightPath := path[divisor-1 : len(path)-1]
	minLen := math.Max(float64(len(leftPath)), float64(len(rightPath)))
	fmt.Printf("The path is: %v.\nThe farthest point from the start is: %d\n", path, int(minLen))
	totalAreaInside := calculcateAreaUnderPolygon(path)
	fmt.Printf("The total area inside the polygon is: %d\n", totalAreaInside)
}

func calculcateAreaUnderPolygon(path []Node) int {
	lines := buildLineList(path)
	totalArea := 0
	for _, line := range lines {
		totalArea += calculateAreaOfLine(line)
	}
	return totalArea
}

func calculateAreaOfLine(line Line) int {
	// y = dx + c
	var slope = 1
	if line.end.y != line.start.y {
		slope = (line.end.x - line.start.x) / (line.end.y - line.start.y)
	}
	// y - y1 = d(x - x1) -> y = dx -dx1 +y1 -> c = -dx1 + y1
	//c := -slope*line.start.x + line.start.y
	height := (line.start.x + line.end.x) / 2
	width := math.Abs(float64(line.start.y) - float64(line.end.y))
	area := height * int(width) * slope
	return area

}

func buildLineList(nodeList []Node) []Line {
	nodePairs := windowed(nodeList, 2)
	result := []Line{}
	for _, pair := range nodePairs {
		start := pair[0].position
		end := pair[1].position
		result = append(result, Line{start: start, end: end})
	}
	return result
}

func findAnimalPath(matrix [][]Node, edgeMap map[Position]Node, startNode Node) []Node {
	adjancencyMatrix := [][]bool{}
	for i := range matrix {
		adjancencyMatrix = append(adjancencyMatrix, make([]bool, len(matrix[i])))
	}
	return visitNeighbors(&adjancencyMatrix, &edgeMap, &startNode, false)
}

func visitNeighbors(visitedMatrix *[][]bool, edgeMap *map[Position]Node, current *Node, isExit bool) []Node {
	currentPosition := current.position
	if (*visitedMatrix)[currentPosition.x][currentPosition.y] {
		return []Node{}
	} else {
		(*visitedMatrix)[currentPosition.x][currentPosition.y] = true
		entrance, exit := calculatePositionsFromForm(current.position, current.form, isExit)
		var exitPosition Position
		IsExit := false
		if (*visitedMatrix)[entrance.x][entrance.y] {
			exitPosition = exit
			IsExit = true
		} else {
			exitPosition = entrance
		}
		exitNode := (*edgeMap)[exitPosition]
		visitNeighbor(currentPosition, visitedMatrix, current, &exitNode)
		return append([]Node{*current}, visitNeighbors(visitedMatrix, edgeMap, &exitNode, IsExit)...)
	}
}

func visitNeighbor(position Position, visitedMatrix *[][]bool, current *Node, leftNode *Node) {
	if current.leftNeighbor == nil {
		current.leftNeighbor = leftNode
		leftNode.rightNeghbor = current
	} else {
		current.rightNeghbor = leftNode
		leftNode.leftNeighbor = current
	}
}

type Line struct {
	start Position
	end   Position
}

type Position struct {
	x int
	y int
}

type Node struct {
	position     Position
	leftNeighbor *Node
	rightNeghbor *Node
	form         Form
}

type Form int

const (
	VERTICAL   Form = 0
	HORIZONTAL Form = 1
	NE_BEND    Form = 2
	NW_BEND    Form = 3
	SW_BEND    Form = 4
	SE_BEND    Form = 5
	START      Form = 6
	GROUND     Form = 7
)

func readLines(input *os.File) ([][]Node, map[Position]Node, Node) {
	scanner := bufio.NewScanner(input)
	nodeMatrix := [][]Node{}
	edgeMap := map[Position]Node{}
	var startNode Node
	for i := 0; scanner.Scan(); i++ {
		line := scanner.Text()
		points := []rune(line)
		if len(nodeMatrix) < 1 {
			nodeMatrix = make([][]Node, len(points))
		}
		if len(nodeMatrix[i]) < 1 {
			nodeMatrix[i] = make([]Node, len(points))
		}
		for j, point := range points {
			currentPosition := Position{x: i, y: j}
			node := parseNode(point, currentPosition)
			if node.form == START {
				startNode = node
			}
			edgeMap[currentPosition] = node
			nodeMatrix[i][j] = node
		}
	}
	return nodeMatrix, edgeMap, startNode
}

func parseNode(input rune, position Position) Node {
	return Node{
		position: position,
		form:     getForm(input)}
}

func calculatePositionsFromForm(position Position, form Form, isExit bool) (Position, Position) {
	var entrance Position
	var exit Position
	switch form {
	case HORIZONTAL:
		entrance = Position{x: position.x, y: position.y - 1}
		exit = Position{x: position.x, y: position.y + 1}
	case VERTICAL:
		entrance = Position{x: position.x - 1, y: position.y}
		exit = Position{x: position.x + 1, y: position.y}
	case NE_BEND:
		entrance = Position{x: position.x - 1, y: position.y}
		exit = Position{x: position.x, y: position.y + 1}
	case NW_BEND:
		entrance = Position{x: position.x - 1, y: position.y}
		exit = Position{x: position.x, y: position.y - 1}
	case SE_BEND:
		entrance = Position{x: position.x + 1, y: position.y}
		exit = Position{x: position.x, y: position.y + 1}
	case SW_BEND:
		entrance = Position{x: position.x + 1, y: position.y}
		exit = Position{x: position.x, y: position.y - 1}
	}
	if isExit {
		return exit, entrance
	}
	return entrance, exit
}

func getForm(point rune) Form {
	switch point {
	case '|':
		return VERTICAL
	case '-':
		return HORIZONTAL
	case 'L':
		return NE_BEND
	case 'J':
		return NW_BEND
	case '7':
		return SW_BEND
	case 'F':
		return SE_BEND
	case 'S':
		return START
	default:
		return GROUND
	}
}

func inferFormForStart(position Position, matrix [][]Node) Form {
	positionsThatConnect := []Position{}
	for i := position.x - 1; i < position.x+2; i++ {
		if i < 0 || i > len(matrix) {
			continue
		}
		for j := position.y - 1; j < position.y+2; j++ {
			if j < 0 || j > len(matrix[0]) {
				continue
			}
			if len(positionsThatConnect) > 1 {
				return inferFormFromPositions(position, positionsThatConnect)
			}
			if i == position.x && j == position.y {
				continue
			}
			node := matrix[i][j]
			entrance, exit := calculatePositionsFromForm(node.position, node.form, false)
			if entrance == position {
				positionsThatConnect = append(positionsThatConnect, node.position)
			}
			if exit == position {
				positionsThatConnect = append(positionsThatConnect, node.position)
			}
		}
	}
	return inferFormFromPositions(position, positionsThatConnect)
}
func inferFormFromPositions(currentPosition Position, positions []Position) Form {
	leftPosition := positions[0]
	rightPosition := positions[1]
	tempPosLeft := Position{x: currentPosition.x - leftPosition.x, y: currentPosition.y - leftPosition.y}
	tempPosRight := Position{x: currentPosition.x - rightPosition.x, y: currentPosition.y - rightPosition.y}
	positionDirection := Position{x: tempPosLeft.x + tempPosRight.x, y: tempPosLeft.y + tempPosRight.y}
	switch positionDirection {
	case Position{0, 0}:
		if tempPosLeft.x == 0 {
			return HORIZONTAL
		} else if tempPosRight.y == 0 {
			return VERTICAL
		}
	case Position{1, 1}:
		return NW_BEND
	case Position{-1, 1}:
		return SW_BEND
	case Position{1, -1}:
		return NE_BEND
	case Position{-1, -1}:
		return SE_BEND
	}
	return GROUND
}

func windowed(slice []Node, size int) [][]Node {
	var result [][]Node

	for i := 0; i <= len(slice)-size; i += size {
		result = append(result, slice[i:i+size])
	}

	return result
}
