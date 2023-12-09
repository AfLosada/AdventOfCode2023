package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	input, err := os.Open("./input.txt")
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(input)
	scanner.Scan()
	seedRanges := strings.Fields(strings.Split(scanner.Text(), ":")[1])
	seeds := windowed(seedRanges, 2)
	mapOfMaps := map[string][]Range{
		"seed2soil":            nil,
		"soil2fertilizer":      nil,
		"fertilizer2water":     nil,
		"water2light":          nil,
		"light2temperature":    nil,
		"temperature2humidity": nil,
		"humidity2location":    nil,
	}
	mapOrder := []string{}
	var currentMap string
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}
		if strings.Contains(line, ":") {
			currentMap = strings.Replace(strings.Split(line, " ")[0], "-to-", "2", 1)
			mapOrder = append(mapOrder, currentMap)
			continue
		}
		mapOfMaps[currentMap] = append(mapOfMaps[currentMap], parseLineToMap(line))
	}
	locationList := []int{}

	for _, seed := range seeds {
		location := seed
		for _, m := range mapOrder {
			location = mapFromRanges(location, mapOfMaps[m])
		}
		locationList = append(locationList, location)
	}
	lowestLocation := slices.Min(locationList)
	fmt.Printf("The lowest location number is: %d\n", lowestLocation)
}

type Range struct {
	destinationStart int
	sourceStart      int
	length           int
}

func parseLineToMap(line string) Range {
	splitLine := strings.Fields(line)
	destinationStart, _ := strconv.Atoi(splitLine[0])
	sourceStart, _ := strconv.Atoi(splitLine[1])
	length, _ := strconv.Atoi(splitLine[2])
	return Range{
		destinationStart: destinationStart,
		sourceStart:      sourceStart,
		length:           length,
	}
}

func mapFromRanges(number int, ranges []Range) int {
	result := number
	for _, r := range ranges {
		result = mapFromRange(result, r)
		if result != number {
			return result
		}
	}
	return result
}

func mapFromRange(number int, r Range) int {
	inRange := (number >= r.sourceStart && number <= (r.sourceStart+r.length))
	if !inRange {
		return number
	}
	positionInRange := r.sourceStart - number
	return r.destinationStart + -positionInRange
}

func windowed(slice []string, size int) [][]string {
	var result [][]string

	for i := 0; i <= len(slice)-size; i += size {
		result = append(result, slice[i:i+size])
	}

	return result
}

func buildSeedsFromPairSlice(sliceOfPairs [][]string) []int {
	result := []int{}
	for _, pair := range sliceOfPairs {
		start, _ := strconv.Atoi(pair[0])
		amount, _ := strconv.Atoi(pair[1])
		for i := 0; i < amount; i++ {
			result = append(result, start+i)
		}
	}
	return result
}

func findMinOfIntersectionOfRanges(seedRange Range, mapRange Range){
	
}