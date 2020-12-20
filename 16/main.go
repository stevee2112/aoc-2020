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

	rules := types.TicketRules{}
	tickets := []types.Ticket{}

	rulesDone := false
	yourTicketDone := false
	yourTicket := types.Ticket{}

	matches := map[string]map[int]int{}

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
				yourTicket = types.NewTicket(strToInt(strings.Split(inputStr,",")))
				yourTicketDone = true
			} else {
				tickets = append(tickets, types.NewTicket(strToInt(strings.Split(inputStr,","))))
			}
		}
	}

	goodTickets := []types.Ticket{}

	// Part 1
	errorRate := 0
	for _,ticket := range tickets{
		ticket.MatchFieldsToRules(rules)

		goodTicket := true
		for _, field := range ticket.Fields {
			if len(field.RulesMatched) == 0 {
				goodTicket = false
				errorRate += field.Value
			}
		}

		if goodTicket {
			goodTickets = append(goodTickets, ticket)
		}
	}

	// Part 2
	for _,ticket := range goodTickets {
		for index, field := range ticket.Fields {
			for _,rule := range field.RulesMatched {

			if _,ok := matches[rule.Name]; !ok {
				matches[rule.Name] = map[int]int{}
			}

			if _,ok := matches[rule.Name][index]; !ok {
				matches[rule.Name][index] = 0
			}

			matches[rule.Name][index]++
			}
		}
	}

	colsMatched := map[int]bool {}
	at := 1
	for at <= len(rules.Fields) {
		for name, match := range matches {
			if len(filter(match)) == at {
				for index,_ := range filter(match) {
					if _,ok := colsMatched[index]; !ok {
						yourTicket.Fields[index].Name = name
						colsMatched[index] = true
						break;
					}
				}
				break;
			}
		}

		at++
	}

	part2 := 1
	for _, field := range yourTicket.Fields {
		departureFields,_ := regexp.Compile("^departure")

		if departureFields.MatchString(field.Name) {
			part2 *= field.Value
		}
	}

	fmt.Printf("Part 1: %d\n", errorRate)
	fmt.Printf("Part 2: %d\n", part2)
}

func strToInt(slice []string) []int {
	intSlice := []int{}
	for _,value := range slice {
		intValue,_ := strconv.Atoi(value)
		intSlice = append(intSlice, intValue)
	}

	return intSlice
}

func filter(f map[int]int) map[int]int {
	max := 0
	filtered := map[int]int{}

	for _, value := range f {
		if value > max {
			max = value
		}
	}

	for key, value := range f {
		if value >= max {
			filtered[key] = value
		}
	}

	return filtered
}
