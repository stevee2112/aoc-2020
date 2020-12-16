package main

import (
	"stevee2112/aoc-2020/util"
	"fmt"
	"os"
	"runtime"
	"path"
	"bufio"
	"regexp"
	//"strings"
	"strconv"
	//	"regexp"
)

var pathCountCache  map[string]int

func main() {

	// Get Data
	_, file, _,  _ := runtime.Caller(0)

	input, _ := os.Open(path.Dir(file) + "/input")

	defer input.Close()
	scanner := bufio.NewScanner(input)

	chart := util.NewDirectedGraph("*", util.East)

	for scanner.Scan() {
		alphareg, _ := regexp.Compile("[^a-zA-Z]+")
		direction := alphareg.ReplaceAllString(scanner.Text(), "")

		numreg, _ := regexp.Compile("[^0-9]+")
		value,_ := strconv.Atoi(numreg.ReplaceAllString(scanner.Text(), ""))

		if (direction == "L") {
		chart.Rotate(util.Left, value)
		}

		if (direction == "R") {
			chart.Rotate(util.Right, value)
		}

		if (direction == "F") {
			chart.MoveBy(util.Forward, value)
		}

		if (direction == "N") {
			chart.MoveBy(util.North, value)
		}


		if (direction == "S") {
			chart.MoveBy(util.South, value)
		}

		if (direction == "E") {
			chart.MoveBy(util.East, value)
		}

		if (direction == "W") {
			chart.MoveBy(util.West, value)
		}
	}

	fmt.Printf("Part 1: %d\n", util.Abs(chart.At.X) + util.Abs(chart.At.Y))
	fmt.Printf("Part 2: %d\n", 0)
}
