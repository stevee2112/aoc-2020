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

	input, _ := os.Open(path.Dir(file) + "/example")

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

	// TODO implement flip func(s) horizontal and vertical probably
	// To connected tiles in proper form
	// then start with a tile and build out from there ending when ALL tiles have been connected
	// connected being, getConnected was called for them
	// this can be done by keeping a map.  and ending when the map size equals the size of all tiles
	// at that point we should have everything we need to build the real image

	blah := getConnected(tiles[2], tiles) // 1171
	blah2 := getConnected(blah[0], tiles) // 2473

	for _,connected := range blah2 {
		fmt.Println("\n")
		connected.Grid.Print()
	}

	// sum := 1
	// tileConnections := map[int][]Tile{}
	// for _,tile := range tiles {
	// 	connected := getConnected(tile, tiles)
	// 	tileConnections[tile.Id] = connected
	// 	if len(connected) == 2 {
	// 		sum *= tile.Id
	// 	}
	// 	count++
	// }

	// for id, connections := range tileConnections {
	// 	fmt.Println(id, len(connections)
	// }

	//fmt.Printf("Part 1: %d\n", sum)
}

func getConnected(tile Tile, tiles []Tile) []Tile {

	connected := []Tile{}


	fmt.Println(tile.Id)
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
			if bottomRow.Checksum() == topRow.Checksum() {
				connected = append(connected, checkTile)
				fmt.Println("TOP MATCH", checkTile.Id, rotations)
			}

			if topRow.Checksum() == flipped.Checksum() {
				connected = append(connected, checkTile)
				fmt.Println("TOP MATCH", checkTile.Id, rotations, "FLIPPED")
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
			if leftCol.Checksum() == RightCol.Checksum() {
				connected = append(connected, checkTile)
				fmt.Println("RIGHT", checkTile.Id, rotations)
			}

			if RightCol.Checksum() == flipped.Checksum() {
				connected = append(connected, checkTile)
				fmt.Println("RIGHT", checkTile.Id, rotations, "flipped")
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
			if bottomRow.Checksum() == topRow.Checksum() {
				connected = append(connected, checkTile)
				fmt.Println("BOTTOM", checkTile.Id, rotations)
			}


			if bottomRow.Checksum() == flipped.Checksum() {
				connected = append(connected, checkTile)
				fmt.Println("BOTTOM", checkTile.Id, rotations, "flipped")
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
			if leftCol.Checksum() == rightCol.Checksum() {
				connected = append(connected, checkTile)
				fmt.Println("LEFT", checkTile.Id, rotations)
			}

			if leftCol.Checksum() == flipped.Checksum() {
				connected = append(connected, checkTile)
				fmt.Println("LEFT", checkTile.Id, rotations, "flipped")
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
