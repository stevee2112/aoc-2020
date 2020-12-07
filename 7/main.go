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

	bagRules := types.NewRules();

	for scanner.Scan() {
		inputStr := scanner.Text()
		bagRules.AddRule(parseAndMakeRule(inputStr))
	}

	fmt.Printf("Part 1: %d\n", len(bagRules.PossibleBags("shiny gold")))
	fmt.Printf("Part 2: %d\n", bagRules.BagCount("shiny gold"))
}

func parseAndMakeRule(ruleString string) types.Rule {
	ruleString = strings.ReplaceAll(ruleString, " no other bags.", "")
	ruleString = strings.ReplaceAll(ruleString, "bags", "")
	ruleString = strings.ReplaceAll(ruleString, "bag", "")
	ruleParts := strings.Split(ruleString, "contain")

	bagColor := strings.TrimSpace(ruleParts[0])

	children := strings.Split(ruleParts[1], ",")

	intReg, _ := regexp.Compile("[^0-9]+")
	strReg, _ := regexp.Compile("[^a-zA-Z ]+")
	childColors := map[string]int{}

	for _,child := range children {
		countInt, _ := strconv.Atoi(intReg.ReplaceAllString(child, ""))

		color := strings.TrimSpace(strReg.ReplaceAllString(child, ""))
		childColors[color] = countInt
	}

	return types.NewRule(bagColor, childColors)
}
