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
