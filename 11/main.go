package main

import (
	"fmt"
	"os"
	"runtime"
	"path"
	"bufio"
	"stevee2112/aoc-2020/util"
	"strings"
	//"strconv"
)

func main() {

	// Get Data
	_, file, _,  _ := runtime.Caller(0)

	input, _ := os.Open(path.Dir(file) + "/input")

	defer input.Close()
	scanner := bufio.NewScanner(input)

	seats := util.Grid{}

	y := 0;
	for scanner.Scan() {
		inputStr := scanner.Text()
		chars := strings.Split(inputStr, "")

		for x, char := range chars {
			seats.SetValue(x, y, char)
		}

		y++
	}

	// Part 1
	seatsPart1 := seats
	checksumPart1 := seatsPart1.Checksum()

	for {
		seatsPart1 = changeAllSeatsPart1(seatsPart1)
		if seatsPart1.Checksum() == checksumPart1 {
			break
		}

		checksumPart1 = seatsPart1.Checksum()
	}

	totalOccupiedPart1 := 0
	seatsPart1.Iterate(func(coordinate util.Coordinate) {
		if fmt.Sprintf("%v", coordinate.Value) == "#" {
			totalOccupiedPart1++
		}
	})

	// Part 2
	seatsPart2 := seats
	checksumPart2 := seatsPart2.Checksum()

	for {
		seatsPart2 = changeAllSeatsPart2(seatsPart2)
		if seatsPart2.Checksum() == checksumPart2 {
			break
		}

		checksumPart2 = seatsPart2.Checksum()
	}

	totalOccupiedPart2 := 0
	seatsPart2.Iterate(func(coordinate util.Coordinate) {
		if fmt.Sprintf("%v", coordinate.Value) == "#" {
			totalOccupiedPart2++
		}
	})

	fmt.Printf("Part 1: %d\n", totalOccupiedPart1)
	fmt.Printf("Part 2: %d\n", totalOccupiedPart2)
}

func changeAllSeatsPart1(seats util.Grid) util.Grid {
	newSeats := util.Grid{}

	seats.Iterate(func(coordinate util.Coordinate) {
		adjacent := seats.GetAdjacent(coordinate)
		newSeatValue := seatChangePart1(coordinate, adjacent)
		newSeats.SetCoordinate(newSeatValue)
	})

	return newSeats
}

func changeAllSeatsPart2(seats util.Grid) util.Grid {
	newSeats := util.Grid{}

	seats.Iterate(func(coordinate util.Coordinate) {
		visible := getVisible(coordinate, seats)
		newSeatValue := seatChangePart2(coordinate, visible)
		newSeats.SetCoordinate(newSeatValue)
	})

	return newSeats
}


func seatChangePart1(seat util.Coordinate, adjacent []util.Coordinate) util.Coordinate {

	currentValue := fmt.Sprintf("%v", seat.Value)
	newValue := currentValue
	occupiedSeats := 0

	for _,seat := range adjacent {
		seatValue := fmt.Sprintf("%v", seat.Value)
		if seatValue == "#" {
			occupiedSeats++
		}
	}

	if currentValue == "L" && occupiedSeats == 0 {
		newValue = "#"
	}

	if currentValue == "#" && occupiedSeats >= 4 {
		newValue = "L"
	}

	newSeat := util.Coordinate{seat.X, seat.Y, newValue}

	return newSeat;
}

func seatChangePart2(seat util.Coordinate, visible []util.Coordinate) util.Coordinate {

	currentValue := fmt.Sprintf("%v", seat.Value)
	newValue := currentValue
	occupiedSeats := 0

	for _,seat := range visible {
		seatValue := fmt.Sprintf("%v", seat.Value)
		if seatValue == "#" {
			occupiedSeats++
		}
	}

	if currentValue == "L" && occupiedSeats == 0 {
		newValue = "#"
	}

	if currentValue == "#" && occupiedSeats >= 5 {
		newValue = "L"
	}

	newSeat := util.Coordinate{seat.X, seat.Y, newValue}

	return newSeat;
}

func getVisible(seat util.Coordinate, seats util.Grid) []util.Coordinate {
	visible := []util.Coordinate{}

	//N
	for _, pathItem := range seats.GetFromSlope(seat, 0, -1) {
		itemValue := fmt.Sprintf("%v", pathItem.Value)

		if itemValue == "L" || itemValue == "#" {
			visible = append(visible, pathItem)
			break;
		}
	}

	//S
	for _, pathItem := range seats.GetFromSlope(seat, 0, 1) {
		itemValue := fmt.Sprintf("%v", pathItem.Value)

		if itemValue == "L" || itemValue == "#" {
			visible = append(visible, pathItem)
			break;
		}
	}

	//E
	for _, pathItem := range seats.GetFromSlope(seat, 1, 0) {
		itemValue := fmt.Sprintf("%v", pathItem.Value)

		if itemValue == "L" || itemValue == "#" {
			visible = append(visible, pathItem)
			break;
		}
	}

	//W
	for _, pathItem := range seats.GetFromSlope(seat, -1, 0) {
		itemValue := fmt.Sprintf("%v", pathItem.Value)

		if itemValue == "L" || itemValue == "#" {
			visible = append(visible, pathItem)
			break;
		}
	}

	//NE
	for _, pathItem := range seats.GetFromSlope(seat, 1, -1) {
		itemValue := fmt.Sprintf("%v", pathItem.Value)

		if itemValue == "L" || itemValue == "#" {
			visible = append(visible, pathItem)
			break;
		}
	}

	//SE
	for _, pathItem := range seats.GetFromSlope(seat, 1, 1) {
		itemValue := fmt.Sprintf("%v", pathItem.Value)

		if itemValue == "L" || itemValue == "#" {
			visible = append(visible, pathItem)
			break;
		}
	}

	//SW
	for _, pathItem := range seats.GetFromSlope(seat, -1, 1) {
		itemValue := fmt.Sprintf("%v", pathItem.Value)

		if itemValue == "L" || itemValue == "#" {
			visible = append(visible, pathItem)
			break;
		}
	}

	//NE
	for _, pathItem := range seats.GetFromSlope(seat, -1, -1) {
		itemValue := fmt.Sprintf("%v", pathItem.Value)

		if itemValue == "L" || itemValue == "#" {
			visible = append(visible, pathItem)
			break;
		}
	}

	return visible
}
