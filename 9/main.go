package main

import (
	"stevee2112/aoc-2020/types"
	"fmt"
	"os"
	"runtime"
	"path"
	"bufio"
	//"strings"
	"strconv"
	//	"regexp"
)

func main() {

	// Get Data
	_, file, _,  _ := runtime.Caller(0)

	input, _ := os.Open(path.Dir(file) + "/input")

	defer input.Close()
	scanner := bufio.NewScanner(input)

	xmas := types.NewXMAS(25, 25)

	for scanner.Scan() {
		inputStr := scanner.Text()

		number, _ := strconv.Atoi(inputStr)
		xmas.AddNumber(number)
	}

	// Part 1
	part1 := xmas.BreakNumber()

	// Part 2
	part2 := xmas.Break(part1)

	fmt.Printf("Part 1: %d\n", part1)
	fmt.Printf("Part 2: %d\n", part2)
}
