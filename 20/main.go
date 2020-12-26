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

	// Part 1
	sum := 1
	for _,tile := range tiles {
		connected := getConnected(tile, tiles)
		if len(connected) == 2 {
			sum *= tile.Id
		}
	}

	// Part 2
	fullGrid := buildImage(tiles)

	foundMonster := false

	rotations := 4
	for rotations > 0 {
		foundMonster, fullGrid = findMonster(fullGrid)


		if foundMonster {
			break
		}

		fullGrid.Rotate90()

		rotations--
	}

	if !foundMonster {
		fullGrid.FlipVertical()

		rotations = 4
		for rotations > 0 {
			foundMonster, fullGrid = findMonster(fullGrid)


			if foundMonster {
				break
			}

			fullGrid.Rotate90()

			rotations--
		}

	}

	if !foundMonster {
		fullGrid.FlipHorizontal()

		rotations = 4
		for rotations > 0 {
			foundMonster, fullGrid = findMonster(fullGrid)


			if foundMonster {
				break
			}

			fullGrid.Rotate90()

			rotations--
		}

	}

	roughness := 0
	fullGrid.Iterate(func (c util.Coordinate) {
		if c.Value == "#" {
			roughness++
		}
	})

	fullGrid.Print()
	fmt.Printf("Part 1: %d\n", sum)
	fmt.Printf("Part 2: %d\n", roughness)
}

func findMonster(image util.Grid) (bool, util.Grid) {

	monsterWidth := 19
	monsterHeight := 2
	hasMonster := false

	image.Iterate(func (coordinate util.Coordinate) {
		x := coordinate.X
		y := coordinate.Y

		subImage := image.Subset(x,y, x + monsterWidth, y + monsterHeight)
		check, imageWithMonster := checkForMonster(subImage)
		if check {
			hasMonster = true
			image.AddGrid(x, y, imageWithMonster)
		}
	})
	return hasMonster, image
}

func checkForMonster(image util.Grid) (bool, util.Grid) {

	mustMatch := []util.Coordinate{
		util.Coordinate{18,0,"#"},
		util.Coordinate{0,1,"#"},
		util.Coordinate{5,1,"#"},
		util.Coordinate{6,1,"#"},
		util.Coordinate{11,1,"#"},
		util.Coordinate{12,1,"#"},
		util.Coordinate{17,1,"#"},
		util.Coordinate{18,1,"#"},
		util.Coordinate{19,1,"#"},
		util.Coordinate{1,2,"#"},
		util.Coordinate{4,2,"#"},
		util.Coordinate{7,2,"#"},
		util.Coordinate{10,2,"#"},
		util.Coordinate{13,2,"#"},
		util.Coordinate{16,2,"#"},
	}

	for _,match := range mustMatch {
		if image.GetValue(match.X, match.Y) != match.Value {
			return false, image
		} else {
			image.SetValue(match.X, match.Y, "O")
		}
	}
	return true, image
}

func buildImage(tiles []Tile) util.Grid {

	tileGrid := util.Grid{}
	tileGridIndex := map[int]util.Coordinate{}
	firstTileCoordinate := util.Coordinate{0, 0, tiles[0]}
	tileGrid.SetCoordinate(firstTileCoordinate)
	tileGridIndex[tiles[0].Id] = firstTileCoordinate


	fullGrid := util.Grid{}
	fullGrid.AddGrid(0, 0, tiles[0].Grid.Subset(1,1,8,8))

	tileSize := 8

	for len(tileGridIndex) < len(tiles) {
		for _,tileCoordinate := range tileGridIndex {
			testTile := tileCoordinate.Value.(Tile)
			connected := getConnected(testTile, tiles)

			for direction,tile := range connected {

				x := tileCoordinate.X
				y := tileCoordinate.Y
				if direction == "TOP" {
					y--
				}

				if direction == "RIGHT" {
					x++
				}

				if direction == "BOTTOM" {
					y++
				}

				if direction == "LEFT" {
					x--
				}

				subGrid := connected[direction].Grid
				fullGrid.AddGrid(x * tileSize, y * tileSize, subGrid.Subset(1,1,8,8))
				position := util.Coordinate{x,y,tile}
				tileGridIndex[tile.Id] = position
				tileGrid.SetCoordinate(position)
			}
		}
	}

	fullGrid.Normalize()
	return fullGrid
}

func getConnected(tile Tile, tiles []Tile) map[string]Tile {

	connected := map[string]Tile{}

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
				connected["TOP"]= checkTile
			}

			if topRow.Checksum() == flipped.Checksum() {
				connected["TOP"]= Tile{checkTile.Id, checkTile.Grid.NewFlipHorizontal()}
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
				connected["RIGHT"]= checkTile
			}

			if RightCol.Checksum() == flipped.Checksum() {
				connected["RIGHT"] = Tile{checkTile.Id, checkTile.Grid.NewFlipVertical()}
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
				connected["BOTTOM"] = checkTile
			}

			if bottomRow.Checksum() == flipped.Checksum() {
				connected["BOTTOM"] = Tile{checkTile.Id, checkTile.Grid.NewFlipHorizontal()}
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
				connected["LEFT"] = checkTile
			}

			if leftCol.Checksum() == flipped.Checksum() {
				connected["LEFT"] = Tile{checkTile.Id, checkTile.Grid.NewFlipVertical()}
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
