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

	input, _ := os.Open(path.Dir(file) + "/input")

	defer input.Close()
	scanner := bufio.NewScanner(input)

	currentMask := ""

	memPart1 := map[int]int{}
	memPart2 := map[int]int{}
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

			for _,realAddress := range applyMemoryDecoder(address, currentMask) {
				memPart2[realAddress] = value
			}
		}
	}

	// Part 1
	sumPart1 := 0
	for _,value := range memPart1 {
		sumPart1 += value
	}

	// Part 2
	sumPart2 := 0
	for _,value := range memPart2 {
		sumPart2 += value
	}


	fmt.Printf("Part 1: %d\n", sumPart1)
	fmt.Printf("Part 2: %d\n", sumPart2)
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

func applyMemoryDecoder(address int, mask string) []int {

	numberBits := strconv.FormatInt(int64(address), 2)

	stringMask := ""
	runes := []rune(mask)
	at := 0

	addresses := []int{}
	addressesBits := []string{}

	for i := len(runes)-1; i >= 0; i-- {

		bit := ""

		if string(mask[i]) == "1" {
			stringMask = fmt.Sprintf("%s%s", stringMask, "1")
			bit = "1"
		}

		if string(mask[i]) == "0" {
			if at < len(numberBits) {
				bit = string(numberBits[len(numberBits) - 1 - at])
				stringMask = fmt.Sprintf("%s%s", stringMask, string(numberBits[len(numberBits) - 1 - at]))
			} else {
				bit = "0"
			}
		}

		if string(mask[i]) == "X" {
			bit = "X"
			stringMask = fmt.Sprintf("%s%s", stringMask, "X")
		}

		addressesBits = addBit(bit, addressesBits)

		at++
	}

	for _, addressBit := range addressesBits {
		address,_ := strconv.ParseInt(reverse(addressBit), 2, 64)
		addresses = append(addresses, int(address))
	}

	return addresses
}

func addBit(bit string, addresses []string) []string {

	newAddresses := []string{}

	if bit == "0" {
		if len(addresses) < 1 {
			newAddresses = append(newAddresses, "0")
		} else {
			for _,address := range addresses {
				newAddresses = append(newAddresses, fmt.Sprintf("%s%s", address, "0"))
			}
		}
	}

	if bit == "1" {
		if len(addresses) < 1 {
			newAddresses = append(newAddresses, "1")
		} else {
			for _,address := range addresses {
				newAddresses = append(newAddresses, fmt.Sprintf("%s%s", address, "1"))
			}
		}
	}

	if bit == "X" {
		newAddresses = append(newAddresses, addBit("0", addresses)...)
		newAddresses = append(newAddresses, addBit("1", addresses)...)
	}

	return newAddresses
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

func reverse(s string) string {
    runes := []rune(s)
    for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
        runes[i], runes[j] = runes[j], runes[i]
    }
    return string(runes)
}
