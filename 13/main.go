package main

import (
	"fmt"
	"os"
	"runtime"
	"path"
	"bufio"
	"strings"
	"strconv"
)

var pathCountCache  map[string]int

func main() {

	// Get Data
	_, file, _,  _ := runtime.Caller(0)

	input, _ := os.Open(path.Dir(file) + "/input")

	defer input.Close()
	scanner := bufio.NewScanner(input)

	scanner.Scan()
	timestamp,_ := strconv.Atoi(scanner.Text())

	scanner.Scan()
	busesStr := strings.Split(scanner.Text(), ",")
	buses := []int{}
	for _,bus := range busesStr {

		busInt := 0
		if bus != "x" {
			busInt,_ = strconv.Atoi(bus)
		}

		buses = append(buses, busInt)
	}

	// Part 1
	bus, arrivesAt := getSoonestBus(timestamp, buses)

	// Part 2 - cheese it
	sequenceTimestamp := getSoonestSequence(406192816267680, buses)

	fmt.Printf("Part 1: %d\n", (arrivesAt - timestamp) * bus)
	fmt.Printf("Part 2: %d\n", sequenceTimestamp)
}

func getSoonestSequence(startTimestamp int, buses []int) int {

	inSequence := false
	timestamp := startTimestamp

	for !inSequence {

		fmt.Println(timestamp)

		inSequence = true
		if timestamp % buses[0] != 0 {
			inSequence = false
			timestamp += 2077047947
			fmt.Println(timestamp)
			continue
		}

		next := timestamp + 1

		for _,bus := range buses[1:] {
			if bus == 0 {
				next++
				continue
			}

			if next % bus != 0 {
				inSequence = false
				break
			}

			fmt.Println("good", bus)
			next++
		}

		if !inSequence {
			timestamp += 2077047947
		}
	}

	return timestamp
}

func getSoonestBus(timestamp int, buses []int) (int, int) {

	for {
		for _,bus := range buses {
			if bus != 0 && timestamp % bus == 0 {
				return bus, timestamp
			}
		}

		timestamp++
	}

	return 0, 0
}
