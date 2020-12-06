package types

import (
)

type Group struct {
	YesAnswers map[string]int
	Size int
}

func NewGroup() Group {
	return Group{
		map[string]int{},
		0,
	}
}
