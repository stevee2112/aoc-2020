package main

import (
	"fmt"
	"os"
	"runtime"
	"path"
	"bufio"
	//"strings"
	//"strconv"
)

func main() {

	// Get Data
	_, file, _,  _ := runtime.Caller(0)

	input, _ := os.Open(path.Dir(file) + "/input")

	defer input.Close()
	scanner := bufio.NewScanner(input)

	groups := []map[string]bool{}
	group := map[string]bool{}
	for scanner.Scan() {
		inputStr := scanner.Text()

		if inputStr == "" {
			groups = append(groups, group)
			group = map[string]bool{}
		} else { // passport data

			for _, question := range inputStr {
				group[string(question)] = true
			}
		}
	}

	// Add last passport
	groups = append(groups, group)

	// Part 1
	totalYes := 0
	for _, group := range groups {
		totalYes += len(group)
	}


	fmt.Printf("Part 1: %d\n", totalYes)
	fmt.Printf("Part 2: %d\n", 0)
}
