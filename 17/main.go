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

	cubeSet := util.Grid3d{}

	y := 0
	z := 0
	for scanner.Scan() {
		inputStr := scanner.Text()
		chars := strings.Split(inputStr, "")

		for x, char := range chars {
			cubeSet.SetValue(x, y, z, char)
		}

		y++
	}

	cycleCount := 6

	for i := 0; i < cycleCount; i++ {
		cubeSet = cycle(cubeSet)
	}

	// Part1
	totalActive := 0
	cubeSet.Iterate(func (c util.Coordinate3d) {
		if c.Value == "#" {
			totalActive++
		}
	})

	fmt.Printf("Part 1: %d\n", totalActive)
	fmt.Printf("Part 2: %d\n", 0)
}

func cycle(cubeSet util.Grid3d) util.Grid3d {
	activeCubes := []util.Coordinate3d{}
	// get all active cubes
	cubeSet.Iterate(func (coordinate util.Coordinate3d) {
		if coordinate.Value == "#" {
			activeCubes = append(activeCubes, coordinate)
		}
	})

	toBeChanged := map[string]util.Coordinate3d{}
	// get all cubes adjacent to a active cube
	for _,cube := range activeCubes {
		for _,adjacent := range cubeSet.GetAdjacent(cube) {
			toBeChanged[adjacent.String()] = adjacent
		}

		// add self
		toBeChanged[cube.String()] = cube
	}

	// for each cube to a cube toggle
	newCubeSet := cubeSet.Clone()
	for _,cube := range toBeChanged {
		newCubeSet.SetCoordinate(getNewValue(cube, cubeSet))
	}

	return newCubeSet
}

func getNewValue(current util.Coordinate3d, grid util.Grid3d) util.Coordinate3d {

	newCube := util.Coordinate3d{current.X, current.Y, current.Z, current.Value}
	adjacent := grid.GetAdjacent(newCube)
	adjacentActive := 0

	for _,adjCube := range adjacent {
		if adjCube.Value == "#" {
			adjacentActive++
		}
	}

	if current.Value == "#" {
		if adjacentActive != 2 && adjacentActive != 3 {
			newCube.Value = "."
		}
	} else {
		if adjacentActive == 3 {
			newCube.Value = "#"
		}
	}

	return newCube
}
