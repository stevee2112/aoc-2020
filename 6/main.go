package main

import (
	"stevee2112/aoc-2020/types"
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

	groups := []types.Group{}
	group := types.NewGroup()
	for scanner.Scan() {
		inputStr := scanner.Text()

		if inputStr == "" {
			groups = append(groups, group)
			group = types.NewGroup()
		} else { // data
			group.Size++
			for _, question := range inputStr {
				group.YesAnswers[string(question)]++
			}
		}
	}

	// Add last passport
	groups = append(groups, group)

	// Part 1
	totalYes := 0
	for _, group := range groups {
		totalYes += len(group.YesAnswers)
	}

	// Part 2
	totalAllAnswered := 0
	for _, group := range groups {
		for _, total := range group.YesAnswers {
			if total == group.Size {
				totalAllAnswered++
			}
		}
	}

	fmt.Printf("Part 1: %d\n", totalYes)
	fmt.Printf("Part 2: %d\n", totalAllAnswered)
}
