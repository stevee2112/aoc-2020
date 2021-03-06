package main

import (
	"stevee2112/aoc-2020/types"
	"fmt"
	"os"
	"runtime"
	"path"
	"bufio"
	"strings"
	//"strconv"
)

func main() {

	// Get Data
	_, file, _,  _ := runtime.Caller(0)

	input, _ := os.Open(path.Dir(file) + "/input")

	defer input.Close()
	scanner := bufio.NewScanner(input)

	passport := types.Passport{}
	currentPassport := &passport
	passports := []types.Passport{}
	for scanner.Scan() {
		inputStr := scanner.Text()

		if inputStr == "" {
			passports = append(passports, *currentPassport)
			passport = types.Passport{}
			currentPassport = &passport
		} else { // passport data
			passportFields := strings.Split(inputStr, " ")

			for _, passportField := range passportFields {
				passportKeyValue := strings.Split(passportField, ":")
				currentPassport.SetField(passportKeyValue[0], passportKeyValue[1])
			}
		}
	}

	// Add last passport
	passports = append(passports, *currentPassport)

	// Part 1
	validPassportsPart1 := 0
	for _, passport := range passports {
		if passport.HasRequiredFields() {
			validPassportsPart1++
		}
	}

	// Part 2
	validPassportsPart2 := 0
	for _, passport := range passports {
		if passport.HasRequiredFields() && passport.IsValid() {
			validPassportsPart2++
		}
	}

	fmt.Printf("Part 1: %d\n", validPassportsPart1)
	fmt.Printf("Part 2: %d\n", validPassportsPart2)
}
