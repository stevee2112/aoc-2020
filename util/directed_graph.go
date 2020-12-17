package util

import (
)

type DirectedGraph struct {
	Map Grid
	At Coordinate
	Visits map[string]int
	Facing Direction
}

type Direction string
type RelativeDirection string

const (
	North = Direction("N")
	South = Direction("S")
	East  = Direction("E")
	West  = Direction("W")
	Forward = RelativeDirection("F")
	Left = RelativeDirection("L")
	Right = RelativeDirection("R")
)

func (dg *DirectedGraph) SetCoordinate(coordinate Coordinate) *DirectedGraph {
	dg.Map.SetCoordinate(coordinate)
	dg.Visits[coordinate.String()]++

	return dg
}

func (dg *DirectedGraph) Rotate(direction RelativeDirection, degrees int) *DirectedGraph {

	var directions  map[Direction]int

	directionsRight := map[Direction]int{
		North: 1,
		East: 2,
		South: 3,
		West: 4,
	}

	directionsLeft := map[Direction]int{
		North: 1,
		West: 2,
		South: 3,
		East: 4,
	}

	if direction == Right {
		directions = directionsRight
	}

	if direction == Left {
		directions = directionsLeft
	}

	rotate := degrees / 90;
	val := (directions[dg.Facing] + rotate) % 4
	if val == 0 {
		val = len(directions)
	}

	for final,directionIndex := range directions {
		if val == directionIndex {
			dg.Facing = final
		}
	}

	return dg
}

func (dg *DirectedGraph) RotateAroundPoint(direction RelativeDirection, degrees int, x int, y int) *DirectedGraph {

	rotate := degrees / 90;

	for rotate > 0 {
		dg.Rotate90AroundPoint(direction, x, y)
		rotate--
	}

	return dg
}

func (dg *DirectedGraph) Rotate90AroundPoint(direction RelativeDirection, x int, y int) *DirectedGraph {
	// take diff
	xDiff := dg.At.X - x
	yDiff := dg.At.Y - y

	if direction == Right {
		// apply diff and move
		dg.MoveTo(-yDiff + x , xDiff + y)
	}

	if direction == Left {
		// apply diff and move
		dg.MoveTo(yDiff + x , -xDiff + y)
	}

	return dg
}

func (dg *DirectedGraph) Move(direction interface{}) *DirectedGraph {
	return dg.MoveBy(direction, 1)
}

func (dg *DirectedGraph) MoveBy(givenDirection interface{}, length int) *DirectedGraph {

	var direction Direction

	switch given := givenDirection.(type) {
	case RelativeDirection:
		direction = GetDirectionRelative(dg.Facing, given)
	case Direction:
		direction = given
	}

	dg.Map.DeleteValue(dg.At.X, dg.At.Y)
	switch (direction) {
	case North:
		dg.At.Y -= length
	case South:
		dg.At.Y += length
	case East:
		dg.At.X += length
	case West:
		dg.At.X -= length
	}

	dg.SetCoordinate(dg.At)

	return dg
}

func (dg *DirectedGraph) MoveTo(x int, y int) *DirectedGraph {

	dg.Map.DeleteValue(dg.At.X, dg.At.Y)
	dg.At.X = x
	dg.At.Y = y

	return dg
}

func NewDirectedGraph (value interface{}, facing Direction) (*DirectedGraph) {
	dg := DirectedGraph{
		Map: Grid{},
		At: Coordinate{0, 0, value},
		Visits: map[string]int{},
		Facing: facing,
	}

	dg.SetCoordinate(dg.At)

	return &dg
}

func GetDirectionRelative(facing Direction, relative RelativeDirection) Direction {

	direction := facing

	if facing == North {
		switch (relative) {
		case Left:
			direction = West
		case Right:
			direction = East
		}
	}

	if facing == West {
		switch (relative) {
		case Left:
			direction = South
		case Right:
			direction = North
		}
	}

	if facing == South {
		switch (relative) {
		case Left:
			direction = East
		case Right:
			direction = West
		}
	}

	if facing == East {
		switch (relative) {
		case Left:
			direction = North
		case Right:
			direction = South
		}
	}

	return direction
}
