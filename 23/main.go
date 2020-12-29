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

	input, _ := os.Open(path.Dir(file) + "/input")

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

	moves := 100
	for moves > 0 {
		cups = Move(cups)
		moves--
	}

	cups = MoveRingTo(1, cups)

	fmt.Printf("Part 1: %s\n", String(cups.Unlink(GetMaxValue(cups) - 1)))
	fmt.Printf("Part 2: %d\n", 0)
}

func Move(cups *ring.Ring) *ring.Ring{

	maxValue := GetMaxValue(cups)

	// Current value
	current := cups.Value.(int)

	// remove 3
	removed := cups.Unlink(3)
	destination := GetDestination(current, removed, maxValue)

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

	return cups
}

func GetDestination(current int, removed *ring.Ring, maxValue int) int {
	current = current - 1
	isRemoved := false

	if current == 0 {
		current = maxValue
	}

	// if value is in removed try again
    removed.Do(func(p interface{}) {
        if p.(int) == current {
			isRemoved = true
        }
    })

	if isRemoved {
		current = GetDestination(current, removed, maxValue)
	}

	return current
}

func String(cups *ring.Ring) string {
	ring := []string{}
	// Iterate through the combined ring and print its contents
	cups.Do(func(p interface{}) {
		ring = append(ring, strconv.Itoa((p.(int))))
	})

	return strings.Join(ring, "")
}

func Print(label string, cups *ring.Ring) {
	ring := []string{}
	// Iterate through the combined ring and print its contents
	cups.Do(func(p interface{}) {
		ring = append(ring, strconv.Itoa((p.(int))))
	})

	fmt.Println(label, strings.Join(ring, ", "))
}

func GetMaxValue(cups *ring.Ring) int {

	maxValue := 0
	cups.Do(func(p interface{}) {
		if p.(int) > maxValue {
			maxValue = p.(int)
		}
	})

	return maxValue
}

func MoveRingTo(value int, cups *ring.Ring) *ring.Ring {

    for cups.Value.(int) != value {
        cups = cups.Move(1)
    }

	return cups
}
