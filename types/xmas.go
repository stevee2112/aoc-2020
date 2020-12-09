package types

import (
)

type XMAS struct {
	series []int
	at int
	window int
}

func NewXMAS(preamble int, window int) XMAS {
	return XMAS{
		[]int{},
		preamble,
		window,
	}
}

func (x *XMAS) AddNumber(number int) {
	x.series = append(x.series, number)
}

func (x *XMAS) Break() int {

	at := x.at
	window := x.window

	for i := at; i < len(x.series); i++ {
		valid := x.SumExists(x.series[i], x.series[(i - window):i])
		if !valid {
			return x.series[i]
		}
	}

	return 0;
}

func (x *XMAS) SumExists(number int, numbers []int) bool {
	for index1, value1 := range numbers {
		for index2 := index1 + 1;index2 < len(numbers); index2++ {
			if value1 + numbers[index2] == number {
				return true
			}
		}
	}

	return false
}
