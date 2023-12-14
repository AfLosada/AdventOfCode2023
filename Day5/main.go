package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"slices"
	"strconv"
	"strings"
	"time"
)

func main() {
	defer timer("main")()
	input, err := os.Open("./input.txt")
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(input)
	scanner.Scan()
	seedRanges := strings.Fields(strings.Split(scanner.Text(), ":")[1])
	seeds := buildSeedsFromPairSlice(windowed(seedRanges, 2))
	mapOfMaps := map[string][]MapRange{}
	mapDirection := map[string]string{}
	mapOrder := []string{}
	var currentMap string
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}
		if strings.Contains(line, ":") {
			mapNames := strings.Split(strings.Split(line, " ")[0], "-to-")
			currentMap = mapNames[0]
			mapDirection[currentMap] = mapNames[1]
			mapOrder = append(mapOrder, currentMap)
			continue
		}
		mapOfMaps[currentMap] = append(mapOfMaps[currentMap], parseLineToMap(line, currentMap))
	}
	locationList := []int{}
	for _, seed := range seeds {
		locationList = append(locationList, recursiveMapping(seed, mapOrder[0], mapDirection, mapOfMaps)...)
	}
	lowestLocation := slices.Min(locationList)
	fmt.Printf("The lowest location number is: %d\n", lowestLocation)
}

func timer(name string) func() {
	start := time.Now()
	return func() {
		fmt.Printf("%s took %v\n", name, time.Since(start))
	}
}

type Range interface {
	start() int
	size() int
}

type MapRange struct {
	destinationStart int
	sourceStart      int
	length           int
	name             string
}
type SeedRange struct {
	destinationStart int
	length           int
}

func (mR MapRange) start() int {
	return mR.destinationStart
}

func (sR SeedRange) start() int {
	return sR.destinationStart
}

func (mR MapRange) size() int {
	return mR.length
}

func (sR SeedRange) size() int {
	return sR.length
}

func parseLineToMap(line string, name string) MapRange {
	splitLine := strings.Fields(line)
	sourceStart, _ := strconv.Atoi(splitLine[0])
	destinationStart, _ := strconv.Atoi(splitLine[1])
	length, _ := strconv.Atoi(splitLine[2])
	return MapRange{
		destinationStart: destinationStart,
		sourceStart:      sourceStart,
		length:           length,
		name:             name,
	}
}

func windowed(slice []string, size int) [][]string {
	var result [][]string

	for i := 0; i <= len(slice)-size; i += size {
		result = append(result, slice[i:i+size])
	}

	return result
}

func buildSeedsFromPairSlice(sliceOfPairs [][]string) []SeedRange {
	result := []SeedRange{}
	for _, pair := range sliceOfPairs {
		start, _ := strconv.Atoi(pair[0])
		amount, _ := strconv.Atoi(pair[1])
		result = append(result, SeedRange{destinationStart: start, length: amount})
	}
	return result
}

func recursiveMapping(r Range, mapName string, mapDirection map[string]string, mapOfMaps map[string][]MapRange) []int {
	dir, ok := mapDirection[mapName]
	if !ok {
		return []int{r.start()}
	}
	containedRanges := []MapRange{}
	for _, mr := range mapOfMaps[dir] {
		containedRange := findContainedRange(r, mr)
		if containedRange.size() == 0 {
			continue
		}
		containedRangeDelta := int(math.Abs(float64(containedRange.start() - mr.destinationStart)))
		containedRanges = append(
			containedRanges,
			MapRange{
				destinationStart: containedRange.start(),
				length:           containedRange.size(),
				sourceStart:      mr.sourceStart + containedRangeDelta,
				name:             mr.name})
	}
	locations := []int{}
	for _, cr := range containedRanges {
		sr := translateRangeToSourceRange(r, cr)
		nextLocations := recursiveMapping(sr, mapDirection[mapName], mapDirection, mapOfMaps)
		locations = append(locations, nextLocations...)
	}
	if len(containedRanges) == 0 {
		locations = append(locations, recursiveMapping(r, mapDirection[mapName], mapDirection, mapOfMaps)...)
	}
	return locations
}

func findContainedRange(r1 Range, r2 Range) Range {
	rMin := r1.start()
	rMax := r1.start() + r1.size()
	mRMin := r2.start()
	mRMax := r2.start() + r2.size()
	if rMin > mRMax || rMax < mRMin {
		return SeedRange{}
	}
	maxMin := max(rMin, mRMin)
	minMax := min(rMax, mRMax)
	rangeDelta := minMax - maxMin

	return SeedRange{destinationStart: maxMin, length: rangeDelta}
}

func translateRangeToSourceRange(r Range, mr MapRange) Range {
	return SeedRange{destinationStart: mr.sourceStart, length: mr.length}
}
