package main

import (
	"fmt"
	"os"
	"runtime"
	"path"
	"bufio"
	"stevee2112/aoc-2020/util"
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

	tiles := []Tile{}

	var tile Tile
	y := 0;
	for scanner.Scan() {
		inputStr := scanner.Text()

		isTileId,_ := regexp.Compile("^Tile")

		if isTileId.MatchString(inputStr) {
			id, _ := strconv.Atoi(strings.Split(strings.Trim(inputStr, ":"), " ")[1])
			tile = NewTile(id)
			continue
		}

		if inputStr == "" {
			y = 0
			tiles = append(tiles, tile)
			continue
		}

		chars := strings.Split(inputStr, "")

		for x, char := range chars {
			tile.Grid.SetValue(x, y, char)
		}

		y++
	}

	// Add last tile
	tiles = append(tiles, tile)


	sum := 1
	for _,tile := range tiles {
		connected := len(getConnected(tile, tiles))
		if connected == 2 {
			fmt.Println("2")
			sum *= tile.Id
		}
	}

	fmt.Printf("Part 1: %d\n", sum)
	// tile = tiles[5]

	// tile.Grid.Print()
	// tile.Grid.Rotate90()
	// fmt.Println("\n")
	// tile.Grid.Print()
}

func getConnected(tile Tile, tiles []Tile) []Tile {

	connected := []Tile{}


	// Top
	topRow := tile.Grid.GetRow(0)
	for _,checkTile := range tiles {

		if tile.Id == checkTile.Id {
			continue
		}

		rotations := 4
		for rotations > 0 {
			bottomRow := checkTile.Grid.GetRow(checkTile.Grid.MaxY)
			flipped := Flip(bottomRow)
			if bottomRow.Checksum() == topRow.Checksum() ||
			    topRow.Checksum() == flipped.Checksum() {
				connected = append(connected, checkTile)
				//fmt.Println("TOP MATCH", checkTile.Id)
			}

			checkTile.Grid.Rotate90()
			rotations--
		}
	}

	// Right
	RightCol := tile.Grid.GetCol(tile.Grid.MaxX)
	for _,checkTile := range tiles {

		if tile.Id == checkTile.Id {
			continue
		}

		rotations := 4
		for rotations > 0 {
			leftCol := checkTile.Grid.GetCol(0)
			flipped := Flip(leftCol)
			if leftCol.Checksum() == RightCol.Checksum() ||
				RightCol.Checksum() == flipped.Checksum() {
				connected = append(connected, checkTile)
				//fmt.Println("RIGHT MATCH", checkTile.Id)
			}

			checkTile.Grid.Rotate90()
			rotations--
		}
	}

	// Bottom
	bottomRow := tile.Grid.GetRow(tile.Grid.MaxY)
	for _,checkTile := range tiles {

		if tile.Id == checkTile.Id {
			continue
		}

		rotations := 4
		for rotations > 0 {
			topRow := checkTile.Grid.GetRow(0)
			flipped := Flip(topRow)

			// since the image can be flipped with should row flipped as well
			if bottomRow.Checksum() == topRow.Checksum() ||
				bottomRow.Checksum() == flipped.Checksum() {
				connected = append(connected, checkTile)
				//fmt.Println("BOTTOM MATCH", checkTile.Id)
			}

			checkTile.Grid.Rotate90()
			rotations--
		}
	}

	// Left
	leftCol := tile.Grid.GetCol(0)
	for _,checkTile := range tiles {

		if tile.Id == checkTile.Id {
			continue
		}

		rotations := 4
		for rotations > 0 {
			rightCol := checkTile.Grid.GetCol(tile.Grid.MaxX)
			flipped := Flip(rightCol)
			if leftCol.Checksum() == rightCol.Checksum() ||
				leftCol.Checksum() == flipped.Checksum() {
				connected = append(connected, checkTile)
				//fmt.Println("LEFT MATCH", checkTile.Id)
			}

			checkTile.Grid.Rotate90()
			rotations--
		}
	}


	return connected
}

func Flip (set util.Grid) util.Grid {

	flipped := util.Grid{}
	set.Iterate(func (coordinate util.Coordinate) {
		flipped.SetValue(set.MaxX - coordinate.X, set.MaxY - coordinate.Y, coordinate.Value)
	})

	return flipped
}

type Tile struct {
	Id int
	Grid util.Grid
}

func NewTile(id int) Tile {
	return Tile{
		id,
		util.Grid{},
	}
}
