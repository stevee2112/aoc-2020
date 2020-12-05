package main

import (
	"fmt"
	"os"
	"runtime"
	"path"
	"bufio"
	"math"
)

func main() {

	// Get Data
	_, file, _,  _ := runtime.Caller(0)

	input, _ := os.Open(path.Dir(file) + "/input")

	defer input.Close()
	scanner := bufio.NewScanner(input)

	highestSeatId := 0
	for scanner.Scan() {
		inputStr := scanner.Text()

		// Part 1
		if seatId := getSeatId(inputStr);seatId > highestSeatId {
			highestSeatId = seatId
		}
	}

	fmt.Printf("Part 1: %d\n", highestSeatId)
	fmt.Printf("Part 2: %d\n", 0)
}

func getSeatId(seatCode string) int {
	return (getParitionValue(seatCode[:7]) * 8) + getParitionValue(seatCode[7:])
}

func getParitionValue(seatCode string) int {

	max := int(math.Pow(2,float64(len(seatCode))))
	set := Range{0,(max - 1)}

	for _,direction := range seatCode {
		if string(direction) == "F" || string(direction) == "L" {
			set.Lower()
		}

		if string(direction) == "B" || string(direction) == "R" {
			set.Upper()
		}

	}

	if set.min == set.max {
		return set.min
	}

	return 0
}

type Range struct {
	min int
	max int
}

func (r *Range) Lower() {
	r.max = r.max - ((r.max - r.min + 1) / 2)
}

func (r *Range) Upper() {
	r.min = r.min + ((r.max - r.min + 1 ) / 2)
}
