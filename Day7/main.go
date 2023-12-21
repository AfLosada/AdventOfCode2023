package main

import (
	"bufio"
	"cmp"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
	"golang.org/x/exp/maps"
	"golang.org/x/exp/slices"
)

func main() {
	defer timer("main")()
	input, err := os.Open("./input.txt")
	if err != nil {
		panic(err)
	}
	handList := readInput(input)
	handTypeMap := buildEqualHandTypeMap(handList)
	for _, val := range handTypeMap {
		slices.SortFunc(val, cardByCardComparison)
	}
	winningsArray := []Hand{}
	for i := 0; i < len(handTypeMap); i++ {
		handSlice := handTypeMap[HandType(i)]
		winningsArray = append(winningsArray, handSlice...)
	}
	fmt.Printf("The total winnings are: %d\n", calculateTotalWinnings(winningsArray))
}

func calculateTotalWinnings(handList []Hand) int {
	totalWinnings := 0
	for i, hand := range handList {
		totalWinnings += (i + 1) * hand.bid
	}
	return totalWinnings
}

type Hand struct {
	hand string
	bid  int
}

func buildEqualHandTypeMap(handList []Hand) map[HandType][]Hand {
	m := map[HandType][]Hand{
		HIGH_CARD:       {},
		ONE_PAIR:        {},
		TWO_PAIR:        {},
		THREE_OF_A_KIND: {},
		FULL_HOUSE:      {},
		FOUR_OF_A_KIND:  {},
		FIVE_OF_A_KIND:  {},
	}
	for _, hand := range handList {
		m[hand.handType()] = append(m[hand.handType()], hand)
	}
	return m
}

func compareTo(hand Hand, other Hand) int {
	currentType := hand.handType()
	otherType := other.handType()
	typeDifference := currentType - otherType
	if typeDifference == 0 {
		return cardByCardComparison(hand, other)
	}
	return int(typeDifference)
}

func cardByCardComparison(hand Hand, other Hand) int {
	comparisonMap := map[rune]int{
		'A': 5 + 9,
		'K': 4 + 9,
		'Q': 3 + 9,
		'J': 2 + 9,
		'T': 1 + 9,
	}
	comparisonHand := hand.hand
	comparisonOther := other.hand
	for i := range comparisonHand {
		handLetter := comparisonHand[i]
		otherLetter := comparisonOther[i]
		hlValue := 0
		olValue := 0
		if hl, ok := comparisonMap[rune(handLetter)]; ok {
			hlValue = hl
		} else {
			hlValue, _ = strconv.Atoi(string(handLetter))
		}
		if ol, ok := comparisonMap[rune(otherLetter)]; ok {
			olValue = ol
		} else {
			olValue, _ = strconv.Atoi(string(otherLetter))
		}
		diff := hlValue - olValue
		if diff != 0 {
			return hlValue - olValue
		}
	}
	return 0
}

type HandType int

const (
	HIGH_CARD       HandType = 0
	ONE_PAIR        HandType = 1
	TWO_PAIR        HandType = 2
	THREE_OF_A_KIND HandType = 3
	FULL_HOUSE      HandType = 4
	FOUR_OF_A_KIND  HandType = 5
	FIVE_OF_A_KIND  HandType = 6
)

func (hand Hand) handType() HandType {
	letterMap := hand.calculateLetterMap()
	return findTypeFromMap(letterMap)
}

func findTypeFromMap(m map[rune]int) HandType {
	hasFiveOfAKind := false
	hasFourOfAKind := false
	hasThreeOfAKind := false
	hasTwoOfAKind := false
	twoOfAKindCount := 0
	for _, value := range m {
		switch value {
		case 5:
			hasFiveOfAKind = true
		case 4:
			hasFourOfAKind = true
		case 3:
			hasThreeOfAKind = true
		case 2:
			hasTwoOfAKind = true
			twoOfAKindCount++
		default:
			continue
		}
	}
	if hasFiveOfAKind {
		return FIVE_OF_A_KIND
	} else if hasFourOfAKind {
		return FOUR_OF_A_KIND
	} else if hasThreeOfAKind && hasTwoOfAKind {
		return FULL_HOUSE
	} else if hasThreeOfAKind {
		return THREE_OF_A_KIND
	} else if hasTwoOfAKind && twoOfAKindCount == 2 {
		return TWO_PAIR
	} else if hasTwoOfAKind {
		return ONE_PAIR
	} else {
		return HIGH_CARD
	}
}

func (hand Hand) calculateLetterMap() map[rune]int {
	m := map[rune]int{}
	for _, val := range hand.hand {
		r := rune(val)
		m[r]++
	}
	keys := maps.Keys(m)
	if strings.Contains(hand.hand, "J") {
		keys = slices.DeleteFunc(keys, func(a rune) bool {
			return a == 'J'
		})
		if len(keys) > 0 {
			mostRepeteadLetter := slices.MaxFunc(keys, func(a, b rune) int {
				return cmp.Compare(m[a], m[b])
			})
			mostRepeatedLetterAmount := strings.Count(hand.hand, string("J"))
			m[mostRepeteadLetter] += mostRepeatedLetterAmount
		}
	}
	return m
}

func readInput(input *os.File) []Hand {
	scanner := bufio.NewScanner(input)
	hands := []Hand{}
	for scanner.Scan() {
		splitLine := strings.Fields(scanner.Text())
		hand := splitLine[0]
		bid, _ := strconv.Atoi(splitLine[1])
		hands = append(hands, Hand{hand: hand, bid: bid})
	}
	return hands
}

func timer(name string) func() {
	start := time.Now()
	return func() {
		fmt.Printf("%s took %v\n", name, time.Since(start))
	}
}
