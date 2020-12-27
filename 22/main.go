package main

import (
	"fmt"
	"os"
	"runtime"
	"path"
	"bufio"
	"container/list"
	"strings"
	"strconv"
	//	"regexp"
)

func main() {

	// Get Data
	_, file, _,  _ := runtime.Caller(0)

	input, _ := os.Open(path.Dir(file) + "/input")

	defer input.Close()
	scanner := bufio.NewScanner(input)

	p1CardsPart1 := list.New()
	p2CardsPart1 := list.New()

	var currentPlayerPart1 *list.List
	currentPlayerPart1 = p1CardsPart1

	for scanner.Scan() {
		inputStr := scanner.Text()

		if inputStr == "" {
			currentPlayerPart1 = p2CardsPart1
		}

		card, err := strconv.Atoi(inputStr)

		if err != nil {
			continue
		}

		currentPlayerPart1.PushBack(card)
	}


	// Part 1
	var winnerPart1 *list.List
	for p1CardsPart1.Len() > 0 && p2CardsPart1.Len() > 0 {
		PlayHandPart1(p1CardsPart1, p2CardsPart1)

		if p1CardsPart1.Len() == 0 {
			winnerPart1 = p2CardsPart1
		}

		if p2CardsPart1.Len() == 0 {
			winnerPart1 = p1CardsPart1
		}

	}

	fmt.Printf("Part 1: %d\n", Score(winnerPart1))
	fmt.Printf("Part 2: %d\n", 0)
}

func PlayHandPart1(p1Cards *list.List, p2Cards *list.List) {
	p1Card := GetCard(p1Cards)
	p2Card := GetCard(p2Cards)

	if p1Card > p2Card {
		p1Cards.PushBack(p1Card)
		p1Cards.PushBack(p2Card)
	}

	if p2Card > p1Card {
		p2Cards.PushBack(p2Card)
		p2Cards.PushBack(p1Card)
	}
}

func GetCard(deck *list.List) int {
	card := deck.Front()
	deck.Remove(card)

	return card.Value.(int)
}

func PrintDeck(deck *list.List) {
	deckStr := []string{}
	for e := deck.Front(); e != nil; e = e.Next() {
		deckStr = append(deckStr, strconv.Itoa(e.Value.(int)))
	}

	fmt.Println(strings.Join(deckStr, ", "))
}

func Score(deck *list.List) int {
	score := 0
	at := 1
	for e := deck.Back(); e != nil; e = e.Prev() {
		score += (at * e.Value.(int))
		at++
	}

	return score
}
