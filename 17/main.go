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

	cubeSet3d := util.Grid3d{}
	cubeSet4d := util.Grid4d{}

	y := 0
	z := 0
	w := 0
	for scanner.Scan() {
		inputStr := scanner.Text()
		chars := strings.Split(inputStr, "")

		for x, char := range chars {
			cubeSet3d.SetValue(x, y, z, char)
			cubeSet4d.SetValue(x, y, z, w, char)
		}

		y++
	}

	cycleCount := 6


	// Part1
	for i := 0; i < cycleCount; i++ {
		cubeSet3d = cycle3d(cubeSet3d)
	}

	totalActivePart1 := 0
	cubeSet3d.Iterate(func (c util.Coordinate3d) {
		if c.Value == "#" {
			totalActivePart1++
		}
	})

	// Part2
	for i := 0; i < cycleCount; i++ {
		cubeSet4d = cycle4d(cubeSet4d)
	}

	totalActivePart2 := 0
	cubeSet4d.Iterate(func (c util.Coordinate4d) {
		if c.Value == "#" {
			totalActivePart2++
		}
	})

	fmt.Printf("Part 1: %d\n", totalActivePart1)
	fmt.Printf("Part 2: %d\n", totalActivePart2)
}

func cycle3d(cubeSet3d util.Grid3d) util.Grid3d {
	activeCubes := []util.Coordinate3d{}
	// get all active cubes
	cubeSet3d.Iterate(func (coordinate util.Coordinate3d) {
		if coordinate.Value == "#" {
			activeCubes = append(activeCubes, coordinate)
		}
	})

	toBeChanged := map[string]util.Coordinate3d{}
	// get all cubes adjacent to a active cube
	for _,cube := range activeCubes {
		for _,adjacent := range cubeSet3d.GetAdjacent(cube) {
			toBeChanged[adjacent.String()] = adjacent
		}

		// add self
		toBeChanged[cube.String()] = cube
	}

	// for each cube to a cube toggle
	newCubeSet := cubeSet3d.Clone()
	for _,cube := range toBeChanged {
		newCubeSet.SetCoordinate(getNewValue3d(cube, cubeSet3d))
	}

	return newCubeSet
}

func cycle4d(cubeSet4d util.Grid4d) util.Grid4d {
	activeCubes := []util.Coordinate4d{}
	// get all active cubes
	cubeSet4d.Iterate(func (coordinate util.Coordinate4d) {
		if coordinate.Value == "#" {
			activeCubes = append(activeCubes, coordinate)
		}
	})

	toBeChanged := map[string]util.Coordinate4d{}
	// get all cubes adjacent to a active cube
	for _,cube := range activeCubes {
		for _,adjacent := range cubeSet4d.GetAdjacent(cube) {
			toBeChanged[adjacent.String()] = adjacent
		}

		// add self
		toBeChanged[cube.String()] = cube
	}

	// for each cube to a cube toggle
	newCubeSet := cubeSet4d.Clone()
	for _,cube := range toBeChanged {
		newCubeSet.SetCoordinate(getNewValue4d(cube, cubeSet4d))
	}

	return newCubeSet
}

func getNewValue3d(current util.Coordinate3d, grid util.Grid3d) util.Coordinate3d {

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

func getNewValue4d(current util.Coordinate4d, grid util.Grid4d) util.Coordinate4d {

	newCube := util.Coordinate4d{current.X, current.Y, current.Z, current.W, current.Value}
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

