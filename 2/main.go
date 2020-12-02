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
	part2Valid := 0
	for scanner.Scan() {
		inputStr := scanner.Text()
		inputParts := strings.Split(inputStr, ":")
		policyStr := strings.TrimSpace(inputParts[0])
		policyParts := strings.Split(policyStr, " ")
		policyRange := strings.Split(policyParts[0], "-")
		index1, _ := strconv.Atoi(policyRange[0])
		index2, _ := strconv.Atoi(policyRange[1])
		letter := policyParts[1];
		password := strings.TrimSpace(inputParts[1])

		// part 1
		count := strings.Count(password, letter)
		if count >= index1 && count <= index2 {
			part1Valid++
		}

		// part 2
		pos1 := string(password[index1 - 1])
		pos2 := string(password[index2 - 1])

		if (pos1 == letter || pos2 == letter) && !(pos1 == pos2) {
				part2Valid++
		}
	}

	fmt.Printf("Part 1: %d\n", part1Valid)
	fmt.Printf("Part 2: %d\n", part2Valid)
}
