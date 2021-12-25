package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	inputFile := "input.txt"
	inputs := getInput(inputFile)

	sum := 0
	for _, line := range inputs {
		num := getNumber(line)
		sum += num
	}

	fmt.Printf("Sum: %d", sum)
}

func getNumber(line string) int {
	patterns, digits := processInput(line)
	lookup := createPatternLookup(patterns)

	number := ""
	for _, digit := range digits {
		number += string(lookup[digit])
	}

	num, _ := strconv.Atoi(number)
	return num
}

func processInput(line string) ([]string, []string) {
	parts := strings.Split(line, " | ")
	patterns := strings.Split(parts[0], " ")
	digits := strings.Split(parts[1], " ")

	for index := range patterns {
		patterns[index] = sortString(patterns[index])
	}

	for index := range digits {
		digits[index] = sortString(digits[index])
	}

	return patterns, digits
}

func sortString(input string) string {
	runes := strings.Split(input, "")
	sort.Strings(runes)
	return strings.Join(runes, "")
}

func countSegments(patterns []string) map[int][]rune {
	segmentCount := map[rune]int{
		'a': 0,
		'b': 0,
		'c': 0,
		'd': 0,
		'e': 0,
		'f': 0,
		'g': 0,
	}

	for _, pattern := range patterns {
		for _, r := range pattern {
			segmentCount[r]++
		}
	}

	countSegments := make(map[int][]rune)
	for key, value := range segmentCount {
		_, isFound := countSegments[value]
		if isFound {
			countSegments[value] = append(countSegments[value], key)
		} else {
			countSegments[value] = []rune{key}
		}
	}

	return countSegments
}

func createPatternLookup(patterns []string) map[string]rune {
	lookup := make(map[string]rune)
	digitLookup := make(map[int]string)

	for _, pattern := range patterns {
		switch len(pattern) {
		case 2:
			lookup[pattern] = '1'
			digitLookup[1] = pattern
		case 3:
			lookup[pattern] = '7'
			digitLookup[7] = pattern
		case 4:
			lookup[pattern] = '4'
			digitLookup[4] = pattern
		case 7:
			lookup[pattern] = '8'
			digitLookup[8] = pattern
		}
	}

	segmentCounts := countSegments(patterns)
	for _, pattern := range patterns {
		switch len(pattern) {
		case 5:
			index := strings.Index(pattern, string(segmentCounts[9]))
			if index == -1 {
				lookup[pattern] = '2'
				digitLookup[2] = pattern
			} else {
				topSegmentIndex := strings.Index(pattern, string(segmentCounts[6]))
				bottomSegmentIndex := strings.Index(pattern, string(segmentCounts[4]))

				if bottomSegmentIndex == -1 {
					if topSegmentIndex == -1 {
						lookup[pattern] = '3'
						digitLookup[3] = pattern
					} else {
						lookup[pattern] = '5'
						digitLookup[5] = pattern
					}
				}
			}
		case 6:
			index := strings.Index(pattern, string(segmentCounts[4]))
			if index == -1 {
				lookup[pattern] = '9'
				digitLookup[9] = pattern
			} else {
				for _, seg := range digitLookup[1] {
					index = strings.Index(pattern, string(seg))
					if index == -1 {
						lookup[pattern] = '6'
						digitLookup[6] = pattern
					}
				}
			}
		}
	}

	for _, pattern := range patterns {
		_, wasFound := lookup[pattern]
		if !wasFound {
			lookup[pattern] = '0'
			digitLookup[0] = pattern
		}
	}

	return lookup
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
