package main

import (
	"stevee2112/aoc-2020/util"
	"fmt"
	"os"
	"runtime"
	"path"
	"bufio"
	"strings"
	//"strconv"
	//	"regexp"
)

func main() {

	// Get Data
	_, file, _,  _ := runtime.Caller(0)

	input, _ := os.Open(path.Dir(file) + "/input")

	defer input.Close()
	scanner := bufio.NewScanner(input)

	hexGrid := util.Grid{}
	tileList := [][]string{}

	for scanner.Scan() {
		inputStr := scanner.Text()
		inputChars := strings.Split(inputStr, "")

		tileDirections := []string{}
		for at := 0;at < len(inputChars); at++ {

			len := 1

			if inputChars[at] == "n" || (inputChars[at] == "s") {
				len = 2
			}

			tileDirections = append(tileDirections, (strings.Join(inputChars[at:at + len], "")))
			at += (len - 1)
		}

		tileList = append(tileList, tileDirections)
	}

	for _, tileDirections := range tileList {
		tile := GetTileFromDirections(hexGrid.GetCoordinate(0,0), tileDirections, hexGrid)

		if tile.Value == nil || tile.Value == "W" {
			tile.Value = "B"
		} else {
			tile.Value = "W"
		}

		hexGrid.SetCoordinate(tile)
	}

	hexGrid.PrintWithFill(".")

	sumBlack := 0
	hexGrid.Iterate(func (c util.Coordinate) {
		if c.Value == "B" {
			sumBlack++
		}
	})

	fmt.Printf("Part 1: %d\n", sumBlack)
	fmt.Printf("Part 2: %d\n", 0)
}

func GetTileFromDirections(start util.Coordinate, directions []string, hexGrid util.Grid) util.Coordinate {

	x := start.X
	y := start.Y


	for _,direction := range directions {
		switch (direction) {
		case "e":
			x += 2
		case "ne":
			x++
			y--
		case "nw":
			x--
			y--
		case "w":
			x -= 2
		case "sw":
			x--
			y++
		case "se":
			x++
			y++
		}
	}

	finalAt := hexGrid.GetCoordinate(x, y)

	if finalAt.Value == nil {
		finalAt = util.Coordinate{x, y, nil}
	}

	return finalAt
}
