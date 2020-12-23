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
	"regexp"
)

func main() {
	// Get Data
	_, file, _,  _ := runtime.Caller(0)

	input, _ := os.Open(path.Dir(file) + "/example")

	defer input.Close()
	scanner := bufio.NewScanner(input)

	rules := types.NewMessageRules()
	messages := []string{}

	rulesDone := false
	for scanner.Scan() {
		inputStr := scanner.Text()

		if inputStr == "" {
			rulesDone = true
			continue
		}

		if !rulesDone {
			ruleParts := strings.Split(inputStr, ": ")

			index,_ := strconv.Atoi(ruleParts[0])
			ruleItems := strings.Split(ruleParts[1], " ")
			hasQuote,_ := regexp.Compile("\"[a-fA-F]+\"")
			if hasQuote.MatchString(ruleParts[1]) {
				ruleItems = strings.Split(ruleParts[1], "")
			}
			rules.AddRule(index, types.NewMessageRule(ruleItems))
		} else {
			messages = append(messages, inputStr)
		}
	}

	ruleIndex := rules.Evaluate(0)

	// Part 1
	matching := 0
	for _,message := range messages {

		if _, ok := ruleIndex[message]; ok {
			matching++
		}
	}

	fmt.Printf("Part 1: %d\n", matching)
	fmt.Printf("Part 2: %d\n", 0)
}
