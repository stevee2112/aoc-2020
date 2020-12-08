package main

import (
	"stevee2112/aoc-2020/types"
	"fmt"
	"os"
	"runtime"
	"path"
	"bufio"
	"strings"
	"strconv"
	//	"regexp"
)

func main() {

	// Get Data
	_, file, _,  _ := runtime.Caller(0)

	input, _ := os.Open(path.Dir(file) + "/example")

	defer input.Close()
	scanner := bufio.NewScanner(input)

	bootSequence := types.NewBootSequence()
	for scanner.Scan() {
		inputStr := scanner.Text()

		inputParts := strings.Split(inputStr, " ")
		instruction := inputParts[0]
		arg, _:= strconv.Atoi(inputParts[1])

		bootSequence.AddInstruction(types.NewInstruction(instruction, arg))
	}

	part1Accumulator, _ := bootSequence.Run()

	fmt.Printf("Part 1: %d\n", part1Accumulator)
	fmt.Printf("Part 2: %d\n", 0)
}
