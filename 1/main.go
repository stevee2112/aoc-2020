package main

import (
	"fmt"
	"os"
	"runtime"
	"path"
    "bufio"
	"strconv"
)

func main() {

	// Get Data
	_, file, _,  _ := runtime.Caller(0)

	input, _ := os.Open(path.Dir(file) + "/input")

	defer input.Close()
	scanner := bufio.NewScanner(input)

	expenseIndex := map[int]int{}

	for scanner.Scan() {
		intStr := scanner.Text()

		expence, _ := strconv.Atoi(intStr)

		expenseIndex[expence]++
	}

	var part1Val1, part1Val2 int

	for expense, _ := range expenseIndex {
		needed := 2020 - expense

		if _, ok := expenseIndex[needed]; ok {
			part1Val1 = expense
			part1Val2 = needed
			break
		}
	}

	fmt.Printf("Part 1: %d\n", part1Val1 * part1Val2)
	//fmt.Printf("Part 2: %d\n", atBasement)
}
