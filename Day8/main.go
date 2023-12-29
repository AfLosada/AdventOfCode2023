package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"

	"golang.org/x/exp/maps"
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
	scanner := bufio.NewScanner(input)
	scanner.Scan()
	firstLine := scanner.Text()
	instructions := parseInstructionsToArray(firstLine)
	graph := parseGraph(scanner)
	keys := maps.Keys(graph)
	filter := func(s []string) []string {
		result := []string{}
		for _, v := range s {
			if v[len(v)-1] == 'A' {
				result = append(result, v)
			}
		}
		return result
	}
	keysToStart := filter(keys)
	minDistance := walkGraphGivenInstructions(instructions, keysToStart, graph)
	fmt.Printf("The path from AAA to ZZZ has a length of: %d\n", minDistance)
}

type Node struct {
	name  string
	left  *Node
	right *Node
}

type Instruction int

const (
	RIGHT Instruction = 0
	LEFT  Instruction = 1
)

func parseInstructionsToArray(inst string) []Instruction {
	result := make([]Instruction, len(inst))
	for i, v := range inst {
		ins := string(v)
		if ins == "R" {
			result[i] = RIGHT
		} else {
			result[i] = LEFT
		}
	}
	return result
}

func walkGraphGivenInstructions(
	instructions []Instruction,
	origins []string,
	m map[string]Node) int {
	stepState := map[string]string{}
	originLCM := map[string]int{}
	stepsToRepeat := map[string]map[string]int{}
	for i := 0; ; i++ {
		currentStep := []string{}
		for _, v := range origins {
			nodeToWalk := v
			val, ok := stepState[nodeToWalk]
			if ok {
				nodeToWalk = val
			}
			directionToWalk := instructions[i%len(instructions)]
			currentNode := walkOneStep(directionToWalk, m[nodeToWalk], m)
			currentNodeToWalk := currentNode.name
			_, existsInSteps := stepsToRepeat[v]
			_, existsInLCM := originLCM[v]
			existsInPath := false
			if currentNodeToWalk[len(currentNodeToWalk)-1] == 'Z' {
				if !existsInSteps {
					stepsToRepeat[v] = map[string]int{}
					stepsToRepeat[v][currentNodeToWalk]++
				} else {
					_, existsInPath = stepsToRepeat[v][currentNodeToWalk]
				}
				if !existsInLCM && existsInPath {
					stepsToRepeat[v][currentNodeToWalk] = i
					originLCM[v] = (i + 1) / 2
				}
			}
			stepState[v] = currentNodeToWalk
			currentStep = append(currentStep, currentNodeToWalk)
		}
		containsAllKeys := true
		for _, origin := range origins {
			if !contains(maps.Keys(originLCM), origin) {
				containsAllKeys = false
				break
			}
		}
		if containsAllKeys {
			vals := maps.Values(originLCM)
			return LCM(vals[0], vals[1], vals[1:]...)
		}
	}
}

func allWordsEndWithZ(word []string) bool {
	for _, v := range word {
		if v[len(v)-1] != 'Z' {
			return false
		}
	}
	return true
}

func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func walkOneStep(instruction Instruction, node Node, m map[string]Node) Node {
	var currentNode string
	if instruction == RIGHT {
		currentNode = node.right.name
	} else if instruction == LEFT {
		currentNode = node.left.name
	}
	return m[currentNode]
}

func parseGraph(scanner *bufio.Scanner) map[string]Node {
	adjacencyMap := map[string]Node{}
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}
		parsedNode := parseNodeFromLine(line)
		if val, ok := adjacencyMap[parsedNode.left.name]; ok {
			parsedNode.left = &val
		}
		if val, ok := adjacencyMap[parsedNode.right.name]; ok {
			parsedNode.right = &val
		}
		adjacencyMap[parsedNode.name] = parsedNode
	}
	connectGraph(&adjacencyMap)
	return adjacencyMap
}

func connectGraph(m *map[string]Node) *map[string]Node {
	for _, v := range *m {
		left := (*m)[v.left.name]
		right := (*m)[v.right.name]
		v.left = &left
		v.right = &right
	}
	return m
}

func parseNodeFromLine(line string) Node {
	splitLine := strings.Split(line, "=")
	node := splitLine[0]
	adjacentSplit := strings.Split(strings.Trim(splitLine[1], "("), ",")
	left := strings.Trim(adjacentSplit[0], " ")
	right := strings.Trim(adjacentSplit[1], " ")
	return Node{
		name:  strings.Trim(node, " "),
		left:  &Node{name: strings.Trim(left, "(")},
		right: &Node{name: strings.Trim(right, ")")}}
}

func GCD(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

// find Least Common Multiple (LCM) via GCD
func LCM(a, b int, integers ...int) int {
	result := a * b / GCD(a, b)

	for i := 0; i < len(integers); i++ {
		result = LCM(result, integers[i])
	}

	return result
}
