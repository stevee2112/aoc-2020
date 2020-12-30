package main

import (
	"stevee2112/aoc-2020/util"
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

	current := 0
	fillSize := 9
	arraySize := fillSize+ 1
	cups := make([]int, arraySize)

	for i := 0 ; i < fillSize; i++ {
		if i < len(inputStr) {
			s := inputStr[i]
			cup, _ := strconv.Atoi(s)
			var value int

			if i == len(inputStr) - 1 {
				value,_ = strconv.Atoi(inputStr[0])
			} else {
				value,_ = strconv.Atoi(inputStr[i + 1])
			}
			cups[cup] = value

			if i == 0 { // set current to first cup
				current = cup
			}
		} else {

			var value int
			if i == fillSize - 1 {
				value,_ = strconv.Atoi(inputStr[0])
			} else {
				value = i + 1
			}

			cups[i] = value
		}
	}

	cupsPtr := &cups

	moves := 100
	for at := 0; at < moves; at++ {
		current, cupsPtr = Move(current, cupsPtr)
	}

	endLabels := GetLabels(cups[1], fillSize - 1, cups)

	strLabels := []string{}
	for _,label := range endLabels {
		strLabels = append(strLabels, strconv.Itoa(label))
	}

	fmt.Printf("Part 1: %s\n", strings.Join(strLabels, ""))
	// fmt.Printf("Part 2: %d\n", 0)
}

func Move(current int, cups *[]int) (int, *[]int) {

	removeCount := 3
	maxValue := util.Max(*cups)

	// Get start of removed
	destination := GetDestination(current, GetLabels((*cups)[current], 3, *cups), maxValue)

	// get what current is pointing to (conceptually, the start of the removed set)
	currentPointer := (*cups)[current]

	// get what end if removed is current pointing to
	removedEndPointer := GetNext(current, removeCount, *cups)

	// get and store what destination is currently pointing do
	desPointer := (*cups)[destination]

	// Point current to end of removed (conceptually remove remove set)
	(*cups)[current] = (*cups)[removedEndPointer]

	// Point destinatation to end of removed set
	(*cups)[destination] = currentPointer


	// Point end remove set to what destination WAS pointing to
	(*cups)[removedEndPointer] = desPointer

	return (*cups)[current], cups
}

func GetDestination(current int, removed []int, maxValue int) int {
	current = current - 1
	isRemoved := false

	if current == 0 {
		current = maxValue
	}

	// if value is in removed try again

	for _,removedValue := range removed {
        if removedValue == current {
			isRemoved = true
        }
    }

	if isRemoved {
		current = GetDestination(current, removed, maxValue)
	}

	return current
}

func GetNext(current int, size int, cups []int) int {

	at := current
	for size > 0 {
		at = cups[at]
		size--
	}

	return at
}

func GetLabels(current int, size int, cups []int) []int {

	labels := []int{}

	at := current
	for size > 0 {
		labels = append(labels, at)
		at = cups[at]
		size--
	}

	return labels
}


func Print(label string, current int, cups []int) {

	cupsStr := []string{strconv.Itoa(current)}

	at := cups[current]
	for {
		if at == current {  // reached end of loop
			break
		}

		cupsStr = append(cupsStr, strconv.Itoa(at))
		at = cups[at]
	}

	fmt.Println(label, strings.Join(cupsStr, ", "))
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
