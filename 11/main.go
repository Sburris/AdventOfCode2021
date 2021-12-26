package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	//inputFile := "input_test.txt"
	inputFile := "input.txt"
	input := getInput(inputFile)
	cavern := createCavern(input)

	// steps := 100
	// for step := 0; step < steps; step++ {
	// 	cavern.passTime()
	// }

	// fmt.Printf("Flashes: %d\n", cavern.flashes)

	step := 0
	for {
		step++
		cavern.passTime()
		if cavern.hasAllFlashed() {
			break
		}
	}

	fmt.Printf("Step: %d\n", step)
}

// Read a list of values from the given file
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

type Cavern struct {
	grid    [10][10]int
	flashes int
}

func createCavern(data []string) *Cavern {
	cavern := Cavern{
		flashes: 0,
	}

	for rowIndex, row := range data {
		for columnIndex, value := range row {
			v, _ := strconv.Atoi(string(value))
			cavern.grid[rowIndex][columnIndex] = v
		}
	}

	return &cavern
}

func (cavern *Cavern) passTime() {
	// Increment all energy
	for x := 0; x < 10; x++ {
		for y := 0; y < 10; y++ {
			cavern.grid[x][y]++
		}
	}

	// Check for flashes
	wasFlash := true
	for wasFlash {
		wasFlash = false
		for x := 0; x < 10; x++ {
			for y := 0; y < 10; y++ {
				if cavern.grid[x][y] > 9 {
					wasFlash = true
					cavern.flashes++
					cavern.handleFlash(x, y)
				}
			}
		}
	}
}

func (cavern *Cavern) handleFlash(x int, y int) {
	cavern.grid[x][y] = 0

	for _, dir := range adjacentDir {
		cavern.flashEnergyTransfer(x+dir[0], y+dir[1])
	}
}

func (cavern *Cavern) flashEnergyTransfer(x int, y int) {
	if x < 0 || x > 9 || y < 0 || y > 9 {
		return // out of bounds
	}

	// transfer energy to an octopus that has already flashed this turn
	if cavern.grid[x][y] > 0 {
		cavern.grid[x][y]++
	}
}

func (cavern *Cavern) hasAllFlashed() bool {
	for x := 0; x < 10; x++ {
		for y := 0; y < 10; y++ {
			if cavern.grid[x][y] > 0 {
				return false
			}
		}
	}

	return true
}

var adjacentDir = [][2]int{
	{-1, -1},
	{-1, 0},
	{-1, 1},
	{1, -1},
	{1, 0},
	{1, 1},
	{0, -1},
	{0, 1},
}
