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

	rules := types.TicketRules{}
	tickets := []types.Ticket{}

	rulesDone := false
	yourTicketDone := false
	for scanner.Scan() {
		inputStr := scanner.Text()

		if inputStr == "" {

			if rulesDone == false {
				rulesDone = true
				continue
			}
		}

		isRule,_ := regexp.Compile(":")

		if isRule.MatchString(inputStr) && !rulesDone {
			parts := strings.Split(inputStr, ":")
			name := parts[0]

			ranges := strings.Split(strings.TrimSpace(parts[1]), " or ")

			range1 := strings.Split(ranges[0], "-")
			range2 := strings.Split(ranges[1], "-")

			range1Min, _ := strconv.Atoi(range1[0])
			range1Max, _ := strconv.Atoi(range1[1])
			range2Min, _ := strconv.Atoi(range2[0])
			range2Max, _ := strconv.Atoi(range2[1])

			rules.Fields = append(rules.Fields, types.NewFieldRule(name, range1Min, range1Max, range2Min, range2Max))
		}

		isTicket,_ := regexp.Compile("[0-9,]+")

		if isTicket.MatchString(inputStr) && rulesDone {

			if yourTicketDone == false {
				// Do noting with my ticket for now
				types.NewTicket(strToInt(strings.Split(inputStr,",")))
				yourTicketDone = true
			} else {
				tickets = append(tickets, types.NewTicket(strToInt(strings.Split(inputStr,","))))
			}
		}
	}

	// Part 1
	errorRate := 0
	for _,ticket := range tickets{
		ticket.MatchFieldsToRules(rules)

		for _,field := range ticket.Fields {
			if len(field.RulesMatched) == 0 {
				errorRate += field.Value
			} else {
				fmt.
			}
		}
	}

	fmt.Printf("Part 1: %d\n", errorRate)
	fmt.Printf("Part 2: %d\n", 0)
}

func strToInt(slice []string) []int {
	intSlice := []int{}
	for _,value := range slice {
		intValue,_ := strconv.Atoi(value)
		intSlice = append(intSlice, intValue)
	}

	return intSlice
}
