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
	"crypto/md5"
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


	p1CardsPart2 := CloneDeck(p1CardsPart1, p1CardsPart1.Len())
	p2CardsPart2 := CloneDeck(p2CardsPart1, p2CardsPart1.Len())

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

	// Part 2
	var winnerPart2 *list.List
	winner := PlayPart2(p1CardsPart2, p2CardsPart2)
	if winner == 1 {
		winnerPart2 = p1CardsPart2
	}

	if winner == 2 {
		winnerPart2 = p2CardsPart2
	}

	fmt.Printf("Part 1: %d\n", Score(winnerPart1))
	fmt.Printf("Part 2: %d\n", Score(winnerPart2))
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

func PlayPart2(p1Cards *list.List, p2Cards *list.List) int {
	deckCache := map[string]bool{}
	winner := 0

	for p1Cards.Len() > 0 && p2Cards.Len() > 0 {

		// Check cache here, if exist player 1 wins
		if _, ok := deckCache[Checksum(p1Cards, p2Cards)]; ok {
			winner = 1
			break
		}

		deckCache[Checksum(p1Cards, p2Cards)] = true

		PlayHandPart2(p1Cards, p2Cards)

		if p1Cards.Len() == 0 {
			winner = 2
		}

		if p2Cards.Len() == 0 {
			winner = 1
		}
	}

	return winner
}

func PlayHandPart2(p1Cards *list.List, p2Cards *list.List) int {

	winner := 0
	p1Card := GetCard(p1Cards)
	p2Card := GetCard(p2Cards)

	if p1Cards.Len() >= p1Card && p2Cards.Len() >= p2Card {
		// Recursive combat
		winner := PlayPart2(CloneDeck(p1Cards, p1Card), CloneDeck(p2Cards, p2Card))

		if winner == 1 {
			p1Cards.PushBack(p1Card)
			p1Cards.PushBack(p2Card)
		}

		if winner == 2 {
			p2Cards.PushBack(p2Card)
			p2Cards.PushBack(p1Card)
		}
	} else {
		if p1Card > p2Card {
			winner = 1
			p1Cards.PushBack(p1Card)
			p1Cards.PushBack(p2Card)
		}

		if p2Card > p1Card {
			winner = 2
			p2Cards.PushBack(p2Card)
			p2Cards.PushBack(p1Card)
		}
	}

	return winner
}

func CloneDeck(deck *list.List, size int) *list.List {
	newDeck := list.New()
	for e := deck.Front(); e != nil; e = e.Next() {
		newDeck.PushBack(e.Value.(int))
		size--
		if size < 1 {
			break
		}
	}

	return newDeck
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

func Checksum(deck1 *list.List, deck2 *list.List) string {

	// Add p1 and p2 start to prevent p1: 1,2 p2: 3,4,5
	// from bein the same as p1: 1,2,3 p2: 4,55

	deckStr := []string{}
	deckStr = append(deckStr, "p1")
	for e := deck1.Front(); e != nil; e = e.Next() {
		deckStr = append(deckStr, strconv.Itoa(e.Value.(int)))
	}

	deckStr = append(deckStr, "p2")
	for e := deck2.Front(); e != nil; e = e.Next() {
		deckStr = append(deckStr, strconv.Itoa(e.Value.(int)))
	}

	return fmt.Sprintf("%x", md5.Sum([]byte(strings.Join(deckStr, ","))))
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
