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

	count := 0
	fillSize := len(inputStr)

	cups := ring.New(fillSize)
	cupsInt := []int{}

	for _,s := range inputStr {
		cup, _ := strconv.Atoi(s)
		cups.Value = cup
		cups = cups.Next()
		cupsInt = append(cupsInt, cup)
		count++
	}

	for at := 10; count < fillSize; at++ {
		cups.Value = at
		cups = cups.Next()
		cupsInt = append(cupsInt, at)
		count++
	}

	Print("0", cups)
	moves := 10
	for at :=1; at <= moves;at++ {
		cups = Move(cups)
		//moved := MoveRingTo(1, Clone(cups))
		Print(strconv.Itoa(at), cups)
		fmt.Println(strconv.Itoa(at), predict(cupsInt, at, fillSize))
		//fmt.Println(strconv.Itoa(at), moved.Next().Value)
	}

	cups = MoveRingTo(1, cups)

	fmt.Printf("Part 1: %s\n", String(cups.Unlink(GetMaxValue(cups) - 1)))
	fmt.Printf("Part 2: %d\n", 0)
}

func predict(set []int, iteration int, size int) []int {

	// (at - 1) - (iteration - 1) mod size
	values := make([]int, size)

    for position,value := range set {
		at := mod((position - 1 - (iteration - 1)), size)
		fmt.Println(at)
		values[at] = value
    }

	return values
}

func Move(cups *ring.Ring) *ring.Ring{

	removeCount := 1
	maxValue := GetMaxValue(cups)

	// Current value
	current := cups.Value.(int)

	// remove
	removed := cups.Unlink(removeCount)
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

func Clone(cups *ring.Ring) *ring.Ring {
	newCups := ring.New(cups.Len())
	cups.Do(func(p interface{}) {
		newCups.Value = p.(int)
		newCups = newCups.Next()

	})

	return newCups
}

func mod(a, b int) int {
    return (a % b + b) % b
}
