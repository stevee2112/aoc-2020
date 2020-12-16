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
	checksum := seats.Checksum()

	for {
		seats = changeAllSeats(seats)
		if seats.Checksum() == checksum {
			break
		}

		checksum = seats.Checksum()
	}

	totalOccupied := 0
	seats.Iterate(func(coordinate util.Coordinate) {
		if fmt.Sprintf("%v", coordinate.Value) == "#" {
			totalOccupied++
		}
	})

	fmt.Printf("Part 1: %d\n", totalOccupied)
	fmt.Printf("Part 2: %d\n", 0)
}

func changeAllSeats(seats util.Grid) util.Grid {
	newSeats := util.Grid{}

	seats.Iterate(func(coordinate util.Coordinate) {
		adjacent := seats.GetAdjacent(coordinate)
		newSeatValue := seatChange(coordinate, adjacent)
		newSeats.SetCoordinate(newSeatValue)
	})

	return newSeats
}

func seatChange(seat util.Coordinate, adjacent []util.Coordinate) util.Coordinate {

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
