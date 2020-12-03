package util

import (
	"fmt"
)

type Grid struct {
	grid map[string]Coordinate
	MaxX int
	MaxY int
	MinX int
	MinY int
}

func (g *Grid) SetValue (x int, y int, value interface{}) {

	if g.grid == nil {
		g.grid = map[string]Coordinate{}
	}

	coordinate := Coordinate{x, y, value}
	g.SetCoordinate(coordinate)
}

func (g *Grid) DeleteValue (x int, y int) {
	delete(g.grid, fmt.Sprintf("%d,%d", x, y))
}

func (g *Grid) GetValue (x int, y int) interface{} {

	coordinate := g.grid[fmt.Sprintf("%d,%d", x, y)]

	return coordinate.Value
}


func (g *Grid) SetCoordinate(coordinate Coordinate) {
	g.grid[coordinate.String()] = coordinate

	if coordinate.X > g.MaxX {
		g.MaxX = coordinate.X
	}

	if coordinate.Y > g.MaxY {
		g.MaxY = coordinate.Y
	}

	if coordinate.X < g.MinX {
		g.MinX = coordinate.X
	}

	if coordinate.Y < g.MinY {
		g.MinY = coordinate.Y
	}
}

func (g *Grid) Print() {

	for y := g.MinX; y <= g.MaxY; y++ {
		row := ""
		for x := g.MinX; x <= g.MaxX; x++ {
			row += fmt.Sprintf("%v", g.GetValue(x, y))
		}

		fmt.Println(row)
	}
}
