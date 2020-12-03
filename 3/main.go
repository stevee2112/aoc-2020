package main

import (
	"fmt"
	"os"
	"runtime"
	"path"
	"bufio"
	"stevee2112/aoc-2020/util"
	"strings"
	//"strconv"
)

func main() {

	// Get Data
	_, file, _,  _ := runtime.Caller(0)

	input, _ := os.Open(path.Dir(file) + "/input")

	defer input.Close()
	scanner := bufio.NewScanner(input)

	forest := util.Grid{}

	y := 0;
	for scanner.Scan() {
		inputStr := scanner.Text()
		chars := strings.Split(inputStr, "")

		for x, char := range chars {
			forest.SetValue(x, y, char)
		}

		y++
	}

	forest.Print()

	// part 1
	slopeX := 3
	slopeY := 1
	trees := 0

	x := slopeX;
	for y := slopeY; y <= forest.MaxY; y += slopeY {
		if forest.GetValue(x, y) == "#" {
			trees++
		}
		x += slopeX
		x = x % (forest.MaxX + 1)
	}

	fmt.Printf("Part 1: %d\n", trees)
	//fmt.Printf("Part 2: %d\n", part2Valid)
}
