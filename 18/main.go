package main

import (
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

	totalSum := 0
	totalSum2 := 0
	for scanner.Scan() {
		inputs := strings.Split(strings.ReplaceAll(scanner.Text(), " ", ""), "")
		totalSum += evaluate(inputs)
		totalSum2 += evaluate2(inputs)
	}

	fmt.Printf("Part 1: %d\n", totalSum)
	fmt.Printf("Part 2: %d\n", totalSum2)
}

func evaluate2(inputs []string) int {

	if len(inputs) == 1 {
		num, _ := strconv.Atoi(inputs[0])
		return num
	}

	// replace all parenthesis first
	at := 0
	for at < len(inputs) {
		input := inputs[at]
		at++
		if input == "(" {
			startAt := at
			endAt := 0
			subMatchCounter := 1
			for _,subValue := range inputs[at:] {
				if subValue == "(" {
					subMatchCounter++
				}

				if subValue == ")" {
					subMatchCounter--
				}

				if subMatchCounter == 0 {
					endAt = at
					break
				}

				at++
			}

			newInput := append(inputs[:startAt - 1], strconv.Itoa(evaluate(inputs[startAt:endAt])))
			newInput = append(newInput, inputs[endAt + 1:]...)

			inputs = newInput
			at = 0 // reset
		}
	}

	at = 1
	currentValue,_ :=  strconv.Atoi(inputs[0])
	nextAction := ""
	for _,input := range inputs[at:] {

		num, err := strconv.Atoi(input)
		if err != nil {
			nextAction = input
		} else { // is number
			if nextAction == "+" {
				currentValue += num
			}

			if nextAction == "*" {
				currentValue *= num
			}

		}
	}

	return currentValue
}

func evaluate(inputs []string) int {

	if len(inputs) == 1 {
		num, _ := strconv.Atoi(inputs[0])
		return num
	}

	// replace all parenthesis first
	at := 0
	for at < len(inputs) {
		input := inputs[at]
		at++
		if input == "(" {
			startAt := at
			endAt := 0
			subMatchCounter := 1
			for _,subValue := range inputs[at:] {
				if subValue == "(" {
					subMatchCounter++
				}

				if subValue == ")" {
					subMatchCounter--
				}

				if subMatchCounter == 0 {
					endAt = at
					break
				}

				at++
			}

			newInput := append(inputs[:startAt - 1], strconv.Itoa(evaluate(inputs[startAt:endAt])))
			newInput = append(newInput, inputs[endAt + 1:]...)

			inputs = newInput
			at = 0 // reset
		}
	}

	at = 1
	currentValue,_ :=  strconv.Atoi(inputs[0])
	nextAction := ""
	for _,input := range inputs[at:] {

		num, err := strconv.Atoi(input)
		if err != nil {
			nextAction = input
		} else { // is number
			if nextAction == "+" {
				currentValue += num
			}

			if nextAction == "*" {
				currentValue *= num
			}

		}
	}

	return currentValue
}
