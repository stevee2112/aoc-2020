package util

import (
	"fmt"
	"crypto/md5"
)

type Grid3d struct {
	grid map[string]Coordinate3d
	MaxX int
	MaxY int
	MinX int
	MinY int
	MinZ int
	MaxZ int
}

func (g *Grid3d) Clone() Grid3d {
	newGrid := Grid3d{}

	g.Iterate(func (coordinate Coordinate3d) {
		newGrid.SetCoordinate(coordinate)
	})

	return newGrid
}

func (g *Grid3d) SetValue (x int, y int, z int, value interface{}) {

	coordinate := Coordinate3d{x, y, z, value}
	g.SetCoordinate(coordinate)
}

func (g *Grid3d) DeleteValue (x int, y int, z int) {
	delete(g.grid, fmt.Sprintf("%d,%d,%d", x, y, z))
}

func (g *Grid3d) GetValue (x int, y int, z int) interface{} {

	coordinate := g.grid[fmt.Sprintf("%d,%d,%d", x, y, z)]

	return coordinate.Value
}

func (g *Grid3d) GetCoordinate (x int, y int, z int) Coordinate3d {

	coordinate := g.grid[fmt.Sprintf("%d,%d,%d", x, y, z)]

	if coordinate.Value == nil {
		coordinate = Coordinate3d{x, y, z, nil}
	}

	return coordinate
}

func (g *Grid3d) SetCoordinate(coordinate Coordinate3d) {

	if g.grid == nil {
		g.grid = map[string]Coordinate3d{}
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

	if coordinate.X < g.MinX {
		g.MinX = coordinate.X
	}

	if coordinate.Y < g.MinY {
		g.MinY = coordinate.Y
	}

	if coordinate.Z < g.MinZ {
		g.MinZ = coordinate.Z
	}
}

func (g *Grid3d) GetAdjacent(coordinate Coordinate3d) []Coordinate3d {

	// TODO
	adjacent := []Coordinate3d{}

	for _,z := range []int{-1, 0, 1} {
		for _,y := range []int{-1, 0, 1} {
			for _,x := range []int{-1, 0, 1} {
				if x != 0 || y != 0 || z != 0 {
					adjacent = append(adjacent, g.GetCoordinate(coordinate.X + x, coordinate.Y + y, coordinate.Z + z))
				}
			}
		}
	}

	return adjacent
}

func (g *Grid3d) Iterate(callback func (coordinate Coordinate3d)) {
	for z := g.MinZ; z <= g.MaxZ; z++ {
		for y := g.MinY; y <= g.MaxY; y++ {
			for x := g.MinX; x <= g.MaxX; x++ {
				callback(g.GetCoordinate(x, y, z))
			}
		}
	}
}

func (g *Grid3d) Checksum() string {
	checksum := ""
	g.Iterate(func(coordinate Coordinate3d) {
		checksum = fmt.Sprintf("%s%v", checksum, coordinate.Value)
	})

	return fmt.Sprintf("%x", md5.Sum([]byte(checksum)))
}

func (g *Grid3d) Print() {

	for z := g.MinZ; z <= g.MaxZ; z++ {
		fmt.Printf("\nz=%d\n", z)
		for y := g.MinY; y <= g.MaxY; y++ {
			row := ""
			for x := g.MinX; x <= g.MaxX; x++ {
				row += fmt.Sprintf("%v", g.GetValue(x, y, z))
			}

			fmt.Println(row)
		}
	}
}

func (g *Grid3d) PrintWithFill(fill string) {

	for z := g.MinZ; z <= g.MaxZ; z++ {
		fmt.Printf("\nz=%d\n", z)
		for y := g.MinY; y <= g.MaxY; y++ {
			row := ""
			for x := g.MinX; x <= g.MaxX; x++ {
				value :=  g.GetValue(x, y, z)

				if value == nil {
					value = fill
				}

				row += fmt.Sprintf("%v", value)
			}

			fmt.Println(row)
		}
	}
}
