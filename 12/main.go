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
)

var pathCountCache  map[string]int

func main() {

	// Get Data
	_, file, _,  _ := runtime.Caller(0)

	input, _ := os.Open(path.Dir(file) + "/input")

	defer input.Close()
	scanner := bufio.NewScanner(input)

	chart := util.NewDirectedGraph("*", util.East)

	chart2 := util.NewDirectedGraph("*", util.East)
	waypoint := util.NewDirectedGraph("W", util.East)

	waypoint.MoveTo(10,-1)

	for scanner.Scan() {
		alphareg, _ := regexp.Compile("[^a-zA-Z]+")
		direction := alphareg.ReplaceAllString(scanner.Text(), "")

		numreg, _ := regexp.Compile("[^0-9]+")
		value,_ := strconv.Atoi(numreg.ReplaceAllString(scanner.Text(), ""))

		if (direction == "L") {
			waypoint.RotateAroundPoint(util.Left, value, chart2.At.X, chart2.At.Y)
			chart.Rotate(util.Left, value)
		}

		if (direction == "R") {
			waypoint.RotateAroundPoint(util.Right, value, chart2.At.X, chart2.At.Y)
			chart.Rotate(util.Right, value)
		}

		if (direction == "F") {
			waypointXDiff := waypoint.At.X - chart2.At.X
			waypointYDiff := waypoint.At.Y - chart2.At.Y
			chart2.MoveTo(chart2.At.X + (value * waypointXDiff), chart2.At.Y + (value * waypointYDiff))
			waypoint.MoveTo(waypoint.At.X + (value * waypointXDiff), waypoint.At.Y + (value * waypointYDiff))
			chart.MoveBy(util.Forward, value)
		}

		if (direction == "N") {
			waypoint.MoveBy(util.North, value)
			chart.MoveBy(util.North, value)
		}

		if (direction == "S") {
			waypoint.MoveBy(util.South, value)
			chart.MoveBy(util.South, value)
		}

		if (direction == "E") {
			waypoint.MoveBy(util.East, value)
			chart.MoveBy(util.East, value)
		}

		if (direction == "W") {
			waypoint.MoveBy(util.West, value)
			chart.MoveBy(util.West, value)
		}
	}

	fmt.Printf("Part 1: %d\n", util.Abs(chart.At.X) + util.Abs(chart.At.Y))
	fmt.Printf("Part 2: %d\n", util.Abs(chart2.At.X) + util.Abs(chart2.At.Y))
}
