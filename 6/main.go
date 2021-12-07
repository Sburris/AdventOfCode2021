package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type schoolOfFish struct {
	timer int
	count int
}

func main () {
	inputFile := "input.txt"
	//inputFile := "test_data/sample_input.txt"
	rawInput := getInput(inputFile)
	schools := processInput(rawInput)

	days := 256
	for day := 0; day < days; day++ {
		//fmt.Printf("Day %d\n", day)
		schools = processDay(schools)
	}

	printSchool(schools)
	count := countFish(schools)
	fmt.Printf("Total Fish %d\n", count)
}

func printSchool (schools []*schoolOfFish) {
	for _, school := range schools {
		fmt.Printf("School: Timer %d\tCount: %d\n", school.timer, school.count)
	}
}

func countFish(schools []*schoolOfFish) int {
	sumOfFish := 0

	for _, school := range schools {
		sumOfFish += school.count
	}

	return sumOfFish
}

func processDay(schools []*schoolOfFish) []*schoolOfFish {
	birthingSchoolIndex := -1
	for index, school := range schools {
		school.timer--
		if school.timer < 0 {
			birthingSchoolIndex = index
		}
	}

	if birthingSchoolIndex >= 0 {
		schools = addFishToSchool(schools, 6, schools[birthingSchoolIndex].count)
		schools[birthingSchoolIndex].timer = 8
	}

	return schools
}

func addFishToSchool(schools []*schoolOfFish, timer int, count int) []*schoolOfFish {
	
	for _, school := range schools {		
		if school.timer == timer {
			school.count += count
			return schools
		}
	}

	newSchool := schoolOfFish{timer: timer, count: count}
	schools = append(schools, &newSchool)
	return schools
}

func processInput(input string) []*schoolOfFish {
	schools := make([]*schoolOfFish, 0, 9)
	fishes := strings.Split(input, ",")

	for _, fish := range fishes {
		timer, err := strconv.Atoi(fish)
		checkErr(err)
		schools = addNewFish(schools, timer)
	}

	return schools
}

func addNewFish(schools []*schoolOfFish, fishTimer int) []*schoolOfFish {
	for _, school := range schools {
		if school.timer == fishTimer {
			school.count++
			return schools
		}
	}

	newSchool := schoolOfFish{timer: fishTimer, count: 1}
	schools = append(schools, &newSchool)
	return schools
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

