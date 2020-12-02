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

	const sum = 2020
	var part1Val1, part1Val2, part2Val1, part2Val2, part2Val3 int

	// Part 1
	_, part1Val1, part1Val2 = getSummers(sum, expenseIndex)

	// Part 2
	found := false
	for value, _ := range expenseIndex {
		needed := sum - value

		found, part2Val2, part2Val3 = getSummers(needed, expenseIndex)

		if found {
			part2Val1 = value
			break;
		}
	}

	fmt.Printf("Part 1: %d\n", part1Val1 * part1Val2)
	fmt.Printf("Part 2: %d\n", part2Val1 * part2Val2 * part2Val3)
}

func getSummers(sum int, values map[int]int) (bool, int, int) {
	for value, _ := range values {
		needed := sum - value

		if _, ok := values[needed]; ok {
			return true, value, needed
		}
	}

	return false, 0, 0
}
