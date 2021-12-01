package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	inputFile := "input.txt"
	inputData := getInput(inputFile)
	windowedData := createWindows(inputData, 3)
	increaseCount := countIncreases(windowedData)

	fmt.Printf("Number of Windows Increases: %d", increaseCount)
}

/// Read a list of values from the given file
func getInput(filename string) []int {
	results := make([]int, 0)

	file, err := os.Open(filename)
	checkErr(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		value, err := strconv.Atoi(scanner.Text())
		checkErr(err)
		results = append(results, value)
	}

	return results
}

// Takes an list of measurements and creates a sum of X consecutive values.
//
// data - the raw input list of values
// windowSize - the number of consecutive values to sum
func createWindows(data []int, windowSize int) []int {
	// Input gate check
	inputLength := len(data)
	if inputLength < windowSize || windowSize < 1 {
		return []int{}
	}

	finalSize := inputLength - (windowSize - 1)
	results := make([]int, finalSize)

	for index := 0; index < finalSize; index++ {
		sum := 0
		for offset := 0; offset < windowSize; offset++ {
			sum += data[index+offset]
		}
		results[index] = sum
	}

	return results
}

// Count how many one value increases compared to is next neighbor
func countIncreases(data []int) int {
	increaseCount := 0

	length := len(data) - 1 // not going to look at final index
	for index := 0; index < length; index++ {
		if data[index] < data[index+1] {
			increaseCount++
		}
	}

	return increaseCount
}

// Check if an error has occured
func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
