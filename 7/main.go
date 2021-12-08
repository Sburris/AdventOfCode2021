package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main (){
	inputFile := "input.txt"
	//inputFile := "test_data/sample_input.txt"
	input := getInput(inputFile)
	crabs := parseInput(input)

	// I don't want to try every possibility, so I figure the most cost effective 
	// would be within one standard devication.  So I only check those.
	avg := calculateAverage(crabs)
	std := calculateStandardDeviation(crabs)

	lower := avg - std
	if lower < 0 {
		lower = 0
	}

	upper := avg + std

	fuelUsed := findMostEfficient(crabs, lower, upper)

	fmt.Println(crabs)
	fmt.Println(avg)
	fmt.Println(std)
	fmt.Println(fuelUsed)
}

func findMostEfficient(crabs []int, lowerBound int, upperBounds int) int {
	bestFuelUsed := 2147483647

	for targetPosition := lowerBound; targetPosition <= upperBounds; targetPosition++ {
		fuelUsed := 0
		for _, sub := range crabs {
			distance := int(math.Abs(float64(targetPosition - sub)))
			fuelUsed += calculateFuel(distance)
		}

		if fuelUsed < bestFuelUsed {
			bestFuelUsed = fuelUsed
		}
	}
	
	return bestFuelUsed
}

func calculateFuel(distance int) int {
	// Part 1
	// return distance

	return (distance * (distance + 1)) / 2
}

func calculateAverage(data []int) int {
	sum := 0
	for _, number := range data {
		sum += number
	}
	average := sum / len(data)
	return average
}

func calculateStandardDeviation(data []int) int {
	mean := calculateAverage(data)

	temp := float64(0)
	for _, number := range data {
		temp += math.Pow(float64(number - mean), 2)
	}
	sd := math.Sqrt(float64(temp/float64(len(data))))
	return int(sd)
}

func parseInput(input string) []int {
	crabs := strings.Split(input, ",")
	results := make([]int, 0, len(crabs))

	for _, crab := range crabs {
		position, err := strconv.Atoi(crab)
		checkErr(err)
		results = append(results, position)
	}

	return results
}

// Read in raw data so we can process it
func getInput(filename string) string {
	file, err := os.Open(filename)
	checkErr(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	results := scanner.Text()

	return results
}

// Check if an error has occured
func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}