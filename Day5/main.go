package main

import (
	"bufio"
	"os"
	"strings"
)

func main() {
	input, err := os.Open("./input.txt")
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(input)
	scanner.Scan()
	seeds := strings.Fields(strings.Split(scanner.Text(), ":")[1])
	mapOfMaps := map[string]map[string]string{
		"seed2Soil":            make(map[string]string),
		"soil2fertilizer":      make(map[string]string),
		"fertilizer2water":     make(map[string]string),
		"water2light":          make(map[string]string),
		"light2temperature":    make(map[string]string),
		"temperature2humidity": make(map[string]string),
		"humidity2location":    make(map[string]string),
	}
	var currentMap string

	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, ":") {
			currentMap = strings.Replace(strings.Split(line, " ")[0], "-to-", "2", 1)
			continue
		}
		mapOfMaps[currentMap] = readLine(line)

	}
}

type Ranges struct {
	destination []
}

func readLine(line string) map[string]string {

}
