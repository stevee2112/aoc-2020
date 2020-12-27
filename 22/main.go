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

	p1Cards := list.New()
	p2Cards := list.New()

	var currentPlayer *list.List
	currentPlayer = p1Cards

	for scanner.Scan() {
		inputStr := scanner.Text()

		if inputStr == "" {
			currentPlayer = p2Cards
		}

		card, err := strconv.Atoi(inputStr)

		if err != nil {
			continue
		}

		currentPlayer.PushBack(card)
	}


	var winner *list.List
	for p1Cards.Len() > 0 && p2Cards.Len() > 0 {
		PlayHand(p1Cards, p2Cards)

		if p1Cards.Len() == 0 {
			winner = p2Cards
		}

		if p2Cards.Len() == 0 {
			winner = p1Cards
		}

	}

	fmt.Printf("Part 1: %d\n", Score(winner))
	fmt.Printf("Part 2: %d\n", 0)
}

func PlayHand(p1Cards *list.List, p2Cards *list.List) {
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
