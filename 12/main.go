package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	inputFile := "input.txt"
	input := getInput(inputFile)
	caveSystem := createCaveSystem(input)
	total := caveSystem.processPath("0,start")
	fmt.Printf("Total: %d\n", total)
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

type CaveSystem struct {
	segments map[string][]string
}

func createCaveSystem(segments []string) *CaveSystem {
	caveSystem := CaveSystem{
		segments: make(map[string][]string),
	}

	for _, segment := range segments {
		parts := strings.Split(segment, "-")

		caveSystem.addSegment(parts[0], parts[1])
		caveSystem.addSegment(parts[1], parts[0])
	}

	return &caveSystem
}

func (cave *CaveSystem) addSegment(from string, to string) {
	if from == "end" || to == "start" {
		return
	}

	_, isFound := cave.segments[from]

	if !isFound {
		cave.segments[from] = make([]string, 0)
	}

	cave.segments[from] = append(cave.segments[from], to)
}

func (cave *CaveSystem) processPath(path string) int {
	fmt.Printf("Processing Path: %s\n", path)
	parts := strings.Split(path, ",")
	lastNode := parts[len(parts)-1]

	if lastNode == "end" {
		fmt.Printf("Path: %s\n", path)
		return 1
	}

	sum := 0
	nextSteps := cave.segments[lastNode]
	for _, node := range nextSteps {
		fmt.Printf("nextStep: %s\n", node)
		if isSmallCave(node) {
			index := strings.Index(path, ","+node+",")
			if index >= 0 {
				if path[0] == '0' {
					//path = "1" + path
					newPath := path + "," + node
					sum += cave.processPath("1" + newPath)
					continue
				} else {
					fmt.Printf("no: %s\n", path)
					continue
				}
			}
		}
		newPath := path + "," + node
		sum += cave.processPath(newPath)
	}

	return sum
}

func isSmallCave(segment string) bool {
	//fmt.Printf("isSmallCave: %s (%t)\n", segment, (segment[0] >= 'a'))
	return (segment[0] >= 'a')
	//return true
}
