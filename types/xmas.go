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

func (x *XMAS) BreakNumber() int {

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

func (x *XMAS) Break(breakNumber int) int {
	for index1, value1 := range x.series {

		sum := value1
		at := index1 + 1
		for sum < breakNumber {
			sum += x.series[at]
			at++
		}

		if sum == breakNumber && (at - 1 != index1) {
			min, max := x.GetMinMax(x.series[index1:at])

			return min + max
		}
	}

	return 0
}

func (x *XMAS) GetMinMax(sequence []int) (int, int) {
	min := 0
	max := 0

	for _, value := range sequence {
		if value < min || min == 0 {
			min = value
		}

		if value > max {
			max = value
		}
	}

	return min, max
}
