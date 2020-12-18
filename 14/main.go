package main

import (
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

	currentMask := ""

	memPart1 := map[int]int{}
	for scanner.Scan() {
		inputStr := scanner.Text()

		alphareg, _ := regexp.Compile("mask")
		if alphareg.MatchString(inputStr) {
			currentMask = strings.Split(inputStr, " = ")[1]
		} else {
			memreg, _ := regexp.Compile("[^0-9=]+")
			memParts := strings.Split(memreg.ReplaceAllString(inputStr, ""), "=")

			address,_ := strconv.Atoi(memParts[0])
			value ,_ := strconv.Atoi(memParts[1])

			memPart1[address] = applyMask(value, currentMask)
		}
	}

	// Part 1
	sumPart1 := 0
	for _,value := range memPart1 {
		sumPart1 += value
	}

	fmt.Printf("Part 1: %d\n", sumPart1)
	fmt.Printf("Part 2: %d\n", 0)
}

func applyMask(value int, mask string) int {

	runes := []rune(mask)
	at := 0
	for i := len(runes)-1; i >= 0; i-- {
		if string(mask[i]) == "1" {
			value = setBit(value, uint(at))
		}

		if string(mask[i]) == "0" {
			value = clearBit(value, uint(at))
		}

		at++
	}

	return value
}

// Sets the bit at pos in the integer n.
func setBit(n int, pos uint) int {
    n |= (1 << pos)
    return n
}

// Clears the bit at pos in n.
func clearBit(n int, pos uint) int {
    mask := ^(1 << pos)
    n &= mask
    return n
}
