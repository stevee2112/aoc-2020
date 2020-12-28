package main

import (
	//"stevee2112/aoc-2020/types"
	"fmt"
	"os"
	"runtime"
	"path"
	"bufio"
	//"strings"
	"strconv"
	//	"regexp"
)

func main() {

	// Get Data
	_, file, _,  _ := runtime.Caller(0)

	input, _ := os.Open(path.Dir(file) + "/input")

	defer input.Close()
	scanner := bufio.NewScanner(input)

	scanner.Scan()
	cardPublicKey,_ := strconv.Atoi(scanner.Text())

	scanner.Scan()
	doorPublicKey,_ := strconv.Atoi(scanner.Text())

	encryptionKey := 0
	cardLoopSize := FindLoopSize(7, cardPublicKey)
	doorLoopSize := FindLoopSize(7, doorPublicKey)

	if Transform(doorPublicKey, cardLoopSize) == Transform(cardPublicKey, doorLoopSize) {
		encryptionKey = Transform(doorPublicKey, cardLoopSize)
	}

	fmt.Printf("Part 1: %d\n", encryptionKey)
	fmt.Printf("Part 2: %d\n", 0)
}

func FindLoopSize(subjectNumber int, publicKey int) int {

	value := 1
	loopSize := 0
	for at := 1; value != publicKey ; at++ {
		value *= subjectNumber
		value = value % 20201227
		loopSize++
	}

	return loopSize
}

func Transform(subjectNumber int, loopSize int) int {

	value := 1

	for at := 0; at < loopSize; at++ {
		value *= subjectNumber
		value = value % 20201227
	}

	return value
}
