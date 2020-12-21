package util

import (
	"fmt"
)

type Coordinate struct {
	X int
	Y int
	Value interface{}
}

func (c *Coordinate) String() string {
	return fmt.Sprintf("%d,%d", c.X, c.Y)
}

type Coordinate3d struct {
	X int
	Y int
	Z int
	Value interface{}
}

func (c *Coordinate3d) String() string {
	return fmt.Sprintf("%d,%d,%d", c.X, c.Y, c.Z)
}
