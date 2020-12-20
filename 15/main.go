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
)

type Number struct {
	PreviousSaid int
	LastSaid int
	TimesSaid int
}

func main() {

	// Get Data
	_, file, _,  _ := runtime.Caller(0)

	input, _ := os.Open(path.Dir(file) + "/input")

	defer input.Close()
	scanner := bufio.NewScanner(input)

	scanner.Scan()
	numbers := strings.Split(scanner.Text(), ",")

	memoryPart1 := map[int]Number{}
	lastNumberPart1 := 0
	for index,value := range numbers {
		lastNumberPart1, _ = strconv.Atoi(value)
		memoryPart1[lastNumberPart1] = newNumber(index)
	}

	memoryPart2 := map[int]Number{}
	lastNumberPart2 := 0
	for index,value := range numbers {
		lastNumberPart2, _ = strconv.Atoi(value)
		memoryPart2[lastNumberPart2] = newNumber(index)
	}

	fmt.Printf("Part 1: %d\n", startGame(memoryPart1, lastNumberPart1, len(numbers), 2020))
	fmt.Printf("Part 2: %d\n", startGame(memoryPart2, lastNumberPart2, len(numbers), 30000000))
}

func startGame(memory map[int]Number, lastSaid int, at int, stopAt int) int {

	for at := at; at < stopAt; at++ {
		said := say(memory, lastSaid, at, stopAt)

		if current,exists := memory[said]; exists {
			current.TimesSaid++
			current.PreviousSaid = current.LastSaid
			current.LastSaid = at
			memory[said] = current
		} else {
			memory[said] = newNumber(at)
		}

		lastSaid = said
	}
	return lastSaid
}

func say(memory map[int]Number, lastSaid int, at int, stopAt int) int {

	number, hasBeenSaid := memory[lastSaid]

	if !hasBeenSaid || number.TimesSaid == 1 {
		return 0
	}

	return number.LastSaid - number.PreviousSaid
}

func newNumber(saidAt int) Number {
	return  Number{-1, saidAt, 1}
}
