package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type vent struct {
	x1 int
	y1 int
	x2 int
	y2 int
}

func main() {
	inputFile := "input.txt"
	//inputFile := "test_data/sample_input.txt"
	rawInput := getInput(inputFile)
	vents := processVentInput(rawInput)
	simpleMap := calculateVentMap(vents, true)
	complexMap := calculateVentMap(vents, false)

	dangerousPointsSimple := calculateDangerousPoints(simpleMap)
	fmt.Printf("Dangerous Points Simple: %d\n", dangerousPointsSimple)

	dangerousPointsComplex := calculateDangerousPoints(complexMap)
	fmt.Printf("Dangerous Points Complex: %d\n", dangerousPointsComplex)

	// displayMap := printVentMap(complexMap)
	// for _, row := range displayMap {
	// 	fmt.Println(row)
	// }
}

// Determine the number of points two or more vents cross eachother
func calculateDangerousPoints(ventMap [][]int) int {
	count := 0
	for _, row := range ventMap {
		for _, cell := range row {
			if cell > 1 {
				count++
			}
		}
	}

	return count
}

// Create a new vent object from raw input
func newVent(input string) *vent {
	vent := vent{}

	segments := strings.Split(input, " ")
	if len(segments) != 3 {
		panic(fmt.Sprintf("Invalid vent input detected: '%s'\n", input))
	}

	vent.x1, vent.y1 = processVentCoordinate(segments[0])
	vent.x2, vent.y2 = processVentCoordinate(segments[2])

	return &vent
}

// Take an coordinate string and turn it into x and y
func processVentCoordinate(input string) (int, int) {
	parts := strings.Split(input, ",")
	if len(parts) != 2 {
		panic(fmt.Sprintf("Unknown vent coordiantes: '%s'\n", input))
	}

	x, err := strconv.Atoi(parts[0])
	checkErr(err)
	y, err := strconv.Atoi(parts[1])
	checkErr(err)

	return x, y
}

// Converts the charted vents into a printable format
func printVentMap(input [][]int) []string {
	verticalSize := len(input)
	ventMap := make([]string, 0, verticalSize)

	for _, row := range input {
		output := ""

		for _, cell := range row {
			if cell == 0 {
				output += "."
			} else {
				output += strconv.Itoa(cell)
			}
		}

		ventMap = append(ventMap, output)
	}

	return ventMap
}

// Creates a 2d vent map showing where and how many vents are at a particular point
func calculateVentMap(input []*vent, isSimple bool) [][]int {
	// Calculate size of map
	maxX := 0
	maxY := 0
	for _, vent := range input {
		if vent.x1 > maxX { maxX = vent.x1 }
		if vent.y1 > maxY { maxY = vent.y1 }
		if vent.x2 > maxX { maxX = vent.x2 }
		if vent.y2 > maxY { maxY = vent.y2 }
	}

	// Account for 0 index counting
	maxX++
	maxY++

	ventMap := make([][]int, maxX)
	for index := 0; index < maxY; index++ {
		ventMap[index] = make([]int, maxX)
	}

	// Draw all horizontal/vertical vents
	for _, vent := range input {
		// Check if vent is horizontal/vertical
		if isSimple && vent.x1 != vent.x2 && vent.y1 != vent.y2 {
			continue; // not horizontal/vertical
		}

		slope, verticalErr := calculateSlope(vent)

		if verticalErr == ErrVerticalLine {
			// Vertical line
			yStart := 0
			yEnd := 0
			if vent.y1 > vent.y2 {
				yStart = vent.y2
				yEnd = vent.y1
			} else{
				yStart = vent.y1
				yEnd = vent.y2
			}

			for y := yStart; y <= yEnd; y++ {
				ventMap[y][vent.x1]++
			}
		} else {
			// any other line
			xStart := 0
			xEnd := 0
			yStart := 0

			if vent.x1 > vent.x2 {
				xStart = vent.x2
				xEnd = vent.x1
				yStart = vent.y2
			} else{
				xStart = vent.x1
				xEnd = vent.x2
				yStart = vent.y1
			}

			currentY := yStart
			for x := xStart; x <= xEnd; x++ {
				ventMap[currentY][x]++
				currentY += slope
			}
		}
	}

	return ventMap
}

var ErrVerticalLine = errors.New("Vertical line")

// Calculate the slope of a vent
func calculateSlope(vent *vent) (int, error) {
	// Check for a single point
	if vent.x1 == vent.x2 && vent.y1 == vent.y2 {
		return 0, nil
	}

	// Vertical Line
	if vent.x1 == vent.x2 {
		return 0, ErrVerticalLine
	}

	changeInX := vent.x2 - vent.x1
	changeInY := vent.y2 - vent.y1
	slope := changeInY / changeInX

	return slope, nil
}

// Turn raw readings into vent objects to work with
func processVentInput(input []string) []*vent {
	inputLength := len(input)
	results := make([]*vent, 0, inputLength)

	for _, ventInput := range input {
		vent := newVent(ventInput)
		results = append(results, vent)
	}

	return results
}

// Read in raw data so we can process it
func getInput(filename string) []string {
	results := make([]string, 0)

	file, err := os.Open(filename)
	checkErr(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		results = append(results, scanner.Text())
	}

	return results
}

// Check if an error has occured
func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}