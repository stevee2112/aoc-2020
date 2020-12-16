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

	input, _ := os.Open(path.Dir(file) + "/example")

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
	seatsPart1.Print()

	for {
		fmt.Printf("\f\f")
		seatsPart1 = changeAllSeatsPart1(seatsPart1)
		seatsPart1.Print()
		if seatsPart1.Checksum() == checksumPart1 {
			break
		}

		checksumPart1 = seatsPart1.Checksum()
	}

	totalOccupied := 0
	seatsPart1.Iterate(func(coordinate util.Coordinate) {
		if fmt.Sprintf("%v", coordinate.Value) == "#" {
			totalOccupied++
		}
	})

	fmt.Printf("Part 1: %d\n", totalOccupied)
	fmt.Printf("Part 2: %d\n", 0)
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
