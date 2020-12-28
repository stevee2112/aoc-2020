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

	hexGrid.Normalize()

	// Part 1
	sumBlack := 0
	hexGrid.Iterate(func (c util.Coordinate) {
		if c.Value == "B" {
			sumBlack++
		}
	})

	// Part 2
	days := 100
	for at := 0; at < days; at++ {
		hexGrid = FlipFloor(hexGrid)
	}

	sumBlackPart2 := 0
	hexGrid.Iterate(func (c util.Coordinate) {
		if c.Value == "B" {
			sumBlackPart2++
		}
	})

	fmt.Printf("Part 1: %d\n", sumBlack)
	fmt.Printf("Part 2: %d\n", sumBlackPart2)
}

func FlipFloor(hexGrid util.Grid) util.Grid {
	newHexGrid := util.Grid{}
	check := []util.Coordinate{}
	// Get all Black Coordinates
	hexGrid.Iterate(func (c util.Coordinate) {
		if c.Value == "B" {
			check = append(check, c)

			// Get all coordinates adjacent to all black coordinates and combine with all black coordinates
			// This is all the coordinates we need to check

			check = append(check, GetAdjacent(c, hexGrid)...)
		}
	})

	//for each apply rules and set value in NEW grid
	for _,tile := range check {
		newHexGrid.SetValue(tile.X, tile.Y, Flip(tile, hexGrid))
	}

	return newHexGrid
}

func Flip(at util.Coordinate, hexGrid util.Grid) string {

	adjacent := GetAdjacent(at, hexGrid)
	blkAdjCount := 0

	for _,adj := range adjacent {
		if adj.Value == "B" {
			blkAdjCount++
		}
	}

	if at.Value == "B" {
		if blkAdjCount == 0 || blkAdjCount > 2 {
			return "W"
		}
	} else { // W
		if blkAdjCount == 2 {
			return "B"
		}
	}

	if at.Value == nil {
		at.Value = "W"
	}

	return at.Value.(string)
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

func GetAdjacent(at util.Coordinate, hexGrid util.Grid) []util.Coordinate {

	adjacent := []util.Coordinate{}
	x := at.X
	y := at.Y

	// e
	east := hexGrid.GetCoordinate(x + 2, y)
	if east.Value == nil {
		east = util.Coordinate{x + 2, y, nil}
	}

	adjacent = append(adjacent, east)

	// ne
	northEast := hexGrid.GetCoordinate(x + 1, y - 1)
	if northEast.Value == nil {
		northEast = util.Coordinate{x + 1, y - 1, nil}
	}

	adjacent = append(adjacent, northEast)

	// nw
	northWest := hexGrid.GetCoordinate(x - 1, y - 1)
	if northWest.Value == nil {
		northWest = util.Coordinate{x - 1, y - 1, nil}
	}

	adjacent = append(adjacent, northWest)

	// w
	west := hexGrid.GetCoordinate(x - 2, y)
	if west.Value == nil {
		west = util.Coordinate{x - 2, y, nil}
	}

	adjacent = append(adjacent, west)

	// sw
	southWest := hexGrid.GetCoordinate(x - 1, y + 1)
	if southWest.Value == nil {
		southWest = util.Coordinate{x - 1, y + 1, nil}
	}

	adjacent = append(adjacent, southWest)

	// se
	southEast := hexGrid.GetCoordinate(x + 1, y + 1)
	if southEast.Value == nil {
		southEast = util.Coordinate{x + 1, y + 1, nil}
	}

	adjacent = append(adjacent, southEast)

	return adjacent
}
