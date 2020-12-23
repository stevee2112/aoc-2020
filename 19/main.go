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

	input, _ := os.Open(path.Dir(file) + "/input")

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
	rule42 := rules.Evaluate(42)
	rule31 := rules.Evaluate(31)

	// Part 1
	matching := 0
	for _,message := range messages {

		if _, ok := ruleIndex[message]; ok {
			matching++
		}
	}
	// Part 2
	// 42 repeating then ending with 31 repeating
	part2Matching := 0
	_, windowSize := getSize(rule42)
	_, rule31Size := getSize(rule31)

	MESSAGE:
	for _,message := range messages {

		at := 0
		count42 := 0
		count31 := 0

		for at < len(message) {
			substring42 := string([]rune(message)[at:at + windowSize])
			if matches(substring42, rule42) {
				count42++
			} else { // switch to 31

				if count42 < 2 {
					continue MESSAGE
				}

				// get remaining message
				for at < len(message) {
					substring31 := string([]rune(message)[at:at + rule31Size])
					if matches(substring31, rule31) {
						count31++

						if (at + rule31Size) == len(message) && (count31 < count42){
							part2Matching++
							continue MESSAGE
						}
					} else {
						continue MESSAGE
					}

					at += rule31Size
				}

				continue MESSAGE
			}

			at += windowSize
		}
	}

	fmt.Printf("Part 1: %d\n", matching)
	fmt.Printf("Part 2: %d\n", part2Matching)
}

func matches(message string, ruleSet map[string]bool) bool {
	if _, ok := ruleSet[message]; ok {
		return true
	} else {
		return false
	}
}

func getSize(rules map[string]bool) (map[int]int, int) {
	sizes := map[int]int{}
	size := 0

	for rule, _ := range rules {
		sizes[len(rule)]++
		size = len(rule)
	}

	return sizes, size
}
