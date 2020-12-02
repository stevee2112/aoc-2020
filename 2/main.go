package main

import (
	"fmt"
	"os"
	"runtime"
	"path"
	"bufio"
	"strings"
	 "strconv"
)

func main() {

	// Get Data
	_, file, _,  _ := runtime.Caller(0)

	input, _ := os.Open(path.Dir(file) + "/input")

	defer input.Close()
	scanner := bufio.NewScanner(input)

	part1Valid := 0
	for scanner.Scan() {
		inputStr := scanner.Text()
		inputParts := strings.Split(inputStr, ":")
		policyStr := strings.TrimSpace(inputParts[0])
		policyParts := strings.Split(policyStr, " ")
		policyRange := strings.Split(policyParts[0], "-")
		min, _ := strconv.Atoi(policyRange[0])
		max, _ := strconv.Atoi(policyRange[1])
		letter := policyParts[1];
		password := strings.TrimSpace(inputParts[1])

		count := strings.Count(password, letter)

		if count >= min && count <= max {
			part1Valid++
		}
	}


	fmt.Printf("Part 1: %d\n", part1Valid)
	//fmt.Printf("Part 2: %d\n", part2Val1 * part2Val2 * part2Val3)
}

