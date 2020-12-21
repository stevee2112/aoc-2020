package util

import (
	"fmt"
	"crypto/md5"
)

type Grid4d struct {
	grid map[string]Coordinate4d
	MaxX int
	MaxY int
	MinX int
	MinY int
	MinZ int
	MaxZ int
	MinW int
	MaxW int
}

func (g *Grid4d) Clone() Grid4d {
	newGrid := Grid4d{}

	g.Iterate(func (coordinate Coordinate4d) {
		newGrid.SetCoordinate(coordinate)
	})

	return newGrid
}

func (g *Grid4d) SetValue (x int, y int, z int, w int, value interface{}) {

	coordinate := Coordinate4d{x, y, z, w, value}
	g.SetCoordinate(coordinate)
}

func (g *Grid4d) DeleteValue (x int, y int, z int, w int) {
	delete(g.grid, fmt.Sprintf("%d,%d,%d", x, y, z, w))
}

func (g *Grid4d) GetValue (x int, y int, z int, w int) interface{} {

	coordinate := g.grid[fmt.Sprintf("%d,%d,%d,%d", x, y, z,w)]

	return coordinate.Value
}

func (g *Grid4d) GetCoordinate (x int, y int, z int, w int) Coordinate4d {

	coordinate := g.grid[fmt.Sprintf("%d,%d,%d,%d", x, y, z, w)]

	if coordinate.Value == nil {
		coordinate = Coordinate4d{x, y, z, w, nil}
	}

	return coordinate
}

func (g *Grid4d) SetCoordinate(coordinate Coordinate4d) {

	if g.grid == nil {
		g.grid = map[string]Coordinate4d{}
	}

	g.grid[coordinate.String()] = coordinate

	if coordinate.X > g.MaxX {
		g.MaxX = coordinate.X
	}

	if coordinate.Y > g.MaxY {
		g.MaxY = coordinate.Y
	}

	if coordinate.Z > g.MaxZ {
		g.MaxZ = coordinate.Z
	}

	if coordinate.W > g.MaxW {
		g.MaxW = coordinate.W
	}

	if coordinate.X < g.MinX {
		g.MinX = coordinate.X
	}

	if coordinate.Y < g.MinY {
		g.MinY = coordinate.Y
	}

	if coordinate.Z < g.MinZ {
		g.MinZ = coordinate.Z
	}

	if coordinate.W < g.MinW {
		g.MinW = coordinate.W
	}
}

func (g *Grid4d) GetAdjacent(coordinate Coordinate4d) []Coordinate4d {

	adjacent := []Coordinate4d{}

	for _,w := range []int{-1, 0, 1} {
		for _,z := range []int{-1, 0, 1} {
			for _,y := range []int{-1, 0, 1} {
				for _,x := range []int{-1, 0, 1} {
					if x != 0 || y != 0 || z != 0 || w != 0 {
						adjacent = append(
							adjacent,
							g.GetCoordinate(coordinate.X + x, coordinate.Y + y, coordinate.Z + z, coordinate.W + w),
						)
					}
				}
			}
		}
	}

	return adjacent
}

func (g *Grid4d) Iterate(callback func (coordinate Coordinate4d)) {
	for w := g.MinW; w <= g.MaxW; w++ {
		for z := g.MinZ; z <= g.MaxZ; z++ {
			for y := g.MinY; y <= g.MaxY; y++ {
				for x := g.MinX; x <= g.MaxX; x++ {
					callback(g.GetCoordinate(x, y, z, w))
				}
			}
		}
	}
}

func (g *Grid4d) Checksum() string {
	checksum := ""
	g.Iterate(func(coordinate Coordinate4d) {
		checksum = fmt.Sprintf("%s%v", checksum, coordinate.Value)
	})

	return fmt.Sprintf("%x", md5.Sum([]byte(checksum)))
}
