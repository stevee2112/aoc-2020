package main

import (
	"fmt"
	"os"
	"runtime"
	"path"
	"bufio"
	"sort"
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

	adapters := []int{}

	for scanner.Scan() {
		inputStr := scanner.Text()
		int, _ := strconv.Atoi(inputStr)

		adapters = append(adapters, int)
	}

	sort.Ints(adapters)

	// Add device adapters
	adapters = append(adapters, adapters[len(adapters) - 1] + 3)

	differences, _ := getJoltRating(adapters)
	pathCount := getPathCount(adapters)

	fmt.Printf("Part 1: %d\n", differences[1] * differences[3])
	fmt.Printf("Part 2: %d\n", pathCount)
}

func getPathCount(adapters []int) int {

	sort.Sort(sort.Reverse(sort.IntSlice(adapters)))
	adapters = append(adapters, 0)

	return  pathCount(adapters)

}

func pathCount(adapters []int) int {

	var paths int

	if len(adapters) == 1 {
		paths = 1
	} else {
		paths = 0

		value := adapters[0]

		for i, value2 := range adapters[1:] {
			if value - value2 <= 3 && value != value2 {
				paths += pathCount(adapters[i + 1:])
			} else {
				break
			}
		}
	}

	fmt.Println(paths)

	return paths
}

func getJoltRating(adapters []int) (map[int]int, int) {

	jolts := 0
	differences := map[int]int{}

	for _, value := range adapters {
		if (value - jolts) > 3 {
			break
		}

		differences[value - jolts]++
		jolts = value
	}

	return differences, jolts
}
