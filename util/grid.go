package util

import (
	"fmt"
	"crypto/md5"
)

type Grid struct {
	grid map[string]Coordinate
	MaxX int
	MaxY int
	MinX int
	MinY int
}

func (g *Grid) SetValue (x int, y int, value interface{}) {

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

func (g *Grid) GetCoordinate (x int, y int) Coordinate {
	coordinate := g.grid[fmt.Sprintf("%d,%d", x, y)]
	return coordinate
}

func (g *Grid) SetCoordinate(coordinate Coordinate) {

	if g.grid == nil {
		g.grid = map[string]Coordinate{}
	}

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

func (g *Grid) GetAdjacent(coordinate Coordinate) []Coordinate {

	adjacent := []Coordinate{}

	// N
	if c := g.GetCoordinate(coordinate.X, coordinate.Y - 1); c.Value != nil {
		adjacent = append(adjacent, c)
	}

	// S
	if c := g.GetCoordinate(coordinate.X, coordinate.Y + 1); c.Value != nil {
		adjacent = append(adjacent, c)
	}

	// E
	if c := g.GetCoordinate(coordinate.X + 1, coordinate.Y); c.Value != nil {
		adjacent = append(adjacent, c)
	}

	// W
	if c := g.GetCoordinate(coordinate.X - 1, coordinate.Y); c.Value != nil {
		adjacent = append(adjacent, c)
	}

	// NE
	if c := g.GetCoordinate(coordinate.X + 1, coordinate.Y - 1); c.Value != nil {
		adjacent = append(adjacent, c)
	}

	// SE
	if c := g.GetCoordinate(coordinate.X + 1, coordinate.Y + 1); c.Value != nil {
		adjacent = append(adjacent, c)
	}

	// NW
	if c := g.GetCoordinate(coordinate.X - 1, coordinate.Y - 1); c.Value != nil {
		adjacent = append(adjacent, c)
	}

	// SW
	if c := g.GetCoordinate(coordinate.X - 1, coordinate.Y + 1); c.Value != nil {
		adjacent = append(adjacent, c)
	}

	return adjacent
}

// Returns all coordinates from the given location and slope till the end of grid
// The first value in the returned slice is the closest to the provided coordinate
func (g *Grid) GetFromSlope(coordinate Coordinate, slopeX int, slopeY int) []Coordinate {
	coordinates := []Coordinate{}

	// No slope given
	if slopeX == 0 && slopeY == 0 {
		return coordinates;
	}

	atX := coordinate.X + slopeX
	atY := coordinate.Y + slopeY

	for {
		newCoordinate := g.GetCoordinate(atX, atY)

		if newCoordinate.Value == nil {
			break
		}

		coordinates = append(coordinates, newCoordinate)
		atX += slopeX
		atY += slopeY
	}

	return coordinates
}

func (g *Grid) Iterate(callback func (coordinate Coordinate)) {
	for y := g.MinY; y <= g.MaxY; y++ {
		for x := g.MinX; x <= g.MaxX; x++ {
			callback(g.GetCoordinate(x, y))
		}
	}
}

func (g *Grid) GetRow(y int) Grid {
	row := Grid{}
	for x := g.MinX; x <= g.MaxX; x++ {
		row.SetValue(x,0, g.GetValue(x, y))
	}

	return row
}

func (g *Grid) GetCol(x int) Grid {
	row := Grid{}
	for y := g.MinY; y <= g.MaxY; y++ {
		row.SetValue(0,y, g.GetValue(x, y))
	}

	return row
}

func (g *Grid) Rotate90() {
	rotated := map[string]Coordinate{}

	g.Iterate(func (coordinate Coordinate) {
		newCoordinate := Coordinate{g.MaxY - coordinate.Y, coordinate.X, coordinate.Value}
		rotated[newCoordinate.String()] = newCoordinate
	});

	g.grid = rotated
}

func (g *Grid) Normalize() {
	normalized := Grid{}
    for y := g.MinY; y <= g.MaxY; y++ {
        for x := g.MinX; x <= g.MaxX; x++ {
			coordinate := g.GetCoordinate(x, y)
			newCoordinate := Coordinate{x - g.MinX, y - g.MinY, coordinate.Value}
			normalized.SetCoordinate(newCoordinate)
        }
    }

	*g = normalized
}

func (g *Grid) Subset(topX int, topY int, bottomX int, bottomY int) Grid {
	subset := Grid{}
    for y := topY; y <= bottomY; y++ {
        for x := topX; x <= bottomX; x++ {
			coordinate := g.GetCoordinate(x, y)
			newCoordinate := Coordinate{x - topX, y - topY, coordinate.Value}
			subset.SetCoordinate(newCoordinate)
        }
    }

	return subset
}

func (g *Grid) FlipHorizontal() {
	g.grid = g.NewFlipHorizontal().grid
}

func (g *Grid) FlipVertical() {
	g.grid = g.NewFlipVertical().grid
}

func (g *Grid) NewFlipHorizontal() Grid {
	flipped := Grid{}
    for y := g.MinY; y <= g.MaxY; y++ {
        for x := g.MinX; x <= g.MaxX; x++ {
			coordinate := g.GetCoordinate(x, y)
			newCoordinate := Coordinate{g.MaxX - coordinate.X, y, coordinate.Value}
			flipped.SetCoordinate(newCoordinate)
        }
    }

	return flipped
}

func (g *Grid) NewFlipVertical() Grid {
	flipped := Grid{}
    for y := g.MinY; y <= g.MaxY; y++ {
        for x := g.MinX; x <= g.MaxX; x++ {
			coordinate := g.GetCoordinate(x, y)
			newCoordinate := Coordinate{x, g.MaxY - coordinate.Y, coordinate.Value}
			flipped.SetCoordinate(newCoordinate)
        }
    }

	return flipped
}

func (g *Grid) Checksum() string {
	checksum := ""
	g.Iterate(func(coordinate Coordinate) {
		checksum = fmt.Sprintf("%s%v", checksum, coordinate.Value)
	})

	return fmt.Sprintf("%x", md5.Sum([]byte(checksum)))
}

func (g *Grid) Print() {

	for y := g.MinY; y <= g.MaxY; y++ {
		row := ""
		for x := g.MinX; x <= g.MaxX; x++ {
			row += fmt.Sprintf("%v", g.GetValue(x, y))
		}

		fmt.Println(row)
	}
}

func (g *Grid) PrintWithFill(fill string) {

	for y := g.MinY; y <= g.MaxY; y++ {
		row := ""
		for x := g.MinX; x <= g.MaxX; x++ {
			value :=  g.GetValue(x, y)

			if value == nil {
				value = fill
			}

			row += fmt.Sprintf("%v", value)
		}

		fmt.Println(row)
	}
}

func (g *Grid) AddGrid(x int, y int, grid Grid) {
	grid.Iterate(func (coordinate Coordinate) {
		g.SetValue(x + coordinate.X, y + coordinate.Y, coordinate.Value)
	});
}
