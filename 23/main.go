package main

import (
	"fmt"
	"os"
	"runtime"
	"path"
	"bufio"
	"strings"
	"strconv"
	//	"regexp"
    "container/ring"
)

func main() {

	// Get Data
	_, file, _,  _ := runtime.Caller(0)

	input, _ := os.Open(path.Dir(file) + "/example")

	defer input.Close()
	scanner := bufio.NewScanner(input)

	scanner.Scan()
	inputStr := strings.Split(scanner.Text(), "")

	cups := ring.New(len(inputStr))

	for _,s := range inputStr {
		cup, _ := strconv.Atoi(s)
		cups.Value = cup
		cups = cups.Next()
	}

	moves := 2
	for moves > 0 {
		Print("cups:", cups)
		cups = Move(cups)
		Print("cups after move:", cups)
		fmt.Println("------")
		moves--
	}

	fmt.Printf("Part 1: %d\n", 0)
	fmt.Printf("Part 2: %d\n", 0)
}

func Move(cups *ring.Ring) *ring.Ring{

	// Current value
	current := cups.Value.(int)
	fmt.Println("current", current)

	// remove 3
	removed := cups.Unlink(3)
	Print("pick up:", removed)
	destination := current - 1
	// Check if part of what we just picked up
	// TODO this should be its own function and also handle if it drop below zero
	// GetDestination(current, removed, MAX VALUE OF CUPS)

	fmt.Println("destination", destination)
	//TODO handle if below zero etc

	//move to desination
	for cups.Value.(int) != destination {
		cups = cups.Move(1)
	}

	// add removed cups
	cups = cups.Link(removed)

	//move back to destination
	for cups.Value.(int) != current {
		cups = cups.Move(1)
	}

	// Finally move current forward 1
	cups = cups.Move(1)


	fmt.Println("final destination", cups.Value)

	return cups
}

func Print(label string, cups *ring.Ring) {
	ring := []string{}
	// Iterate through the combined ring and print its contents
	cups.Do(func(p interface{}) {
		ring = append(ring, strconv.Itoa((p.(int))))
	})

	fmt.Println(label, strings.Join(ring, ", "))
}
