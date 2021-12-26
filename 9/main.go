package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	//inputFile := "test_input.txt"
	inputFile := "input.txt"
	input := getInput(inputFile)
	radar := createRadar(input)
	lowSpots := radar.GetAllLowSpotDepths()

	sum := 0
	for _, value := range lowSpots {
		sum += value + 1
	}

	fmt.Printf("Sum: %d\n", sum)

	basins := radar.GetAllBasinSizes()
	sort.Sort(sort.Reverse(sort.IntSlice(basins)))

	fmt.Println(basins)

	product := basins[0] * basins[1] * basins[2]

	fmt.Printf("Basins Product: %d\n", product)
}

type Radar struct {
	floorDepth []string
	maxX       int
	maxY       int
	floodMap   [][]bool
}

var adjacentDirections = [][2]int{
	{-1, 0},
	{1, 0},
	{0, -1},
	{0, 1},
}

func (radar *Radar) IsLowSpot(x int, y int) (int, bool) {
	depth, isValid := radar.GetFloor(x, y)
	if !isValid {
		return -1, false
	}

	for _, dir := range adjacentDirections {
		adjacentDepth, adjacentFound := radar.GetFloor(x+dir[0], y+dir[1])
		if adjacentFound {
			if adjacentDepth <= depth {
				return depth, false
			}
		}
	}

	return depth, true
}

func (radar *Radar) GetFloor(x int, y int) (int, bool) {
	//fmt.Printf("maxX: %d, maxY: %d\n", radar.maxX, radar.maxY)
	if x < 0 || x >= radar.maxX || y < 0 || y >= radar.maxY {
		return -1, false
	}
	temp := radar.floorDepth[y][x]
	return intLookup[temp], true
}

func (radar *Radar) GetAllBasinSizes() []int {
	results := make([]int, 0)

	for y, row := range radar.floorDepth {
		for x := range row {
			_, isLow := radar.IsLowSpot(x, y)
			if isLow {
				if !radar.floodMap[y][x] {
					results = append(results, radar.getBasinSize(x, y))
				}
			}
		}
	}

	return results
}

func (radar *Radar) GetAllLowSpotDepths() []int {
	results := make([]int, 0)

	for y, row := range radar.floorDepth {
		for x := range row {
			depth, isLow := radar.IsLowSpot(x, y)
			if isLow {
				results = append(results, depth)
			}
		}
	}

	return results
}

var intLookup = map[byte]int{
	'0': 0,
	'1': 1,
	'2': 2,
	'3': 3,
	'4': 4,
	'5': 5,
	'6': 6,
	'7': 7,
	'8': 8,
	'9': 9,
}

func createRadar(input []string) *Radar {
	newRadar := Radar{
		floorDepth: input,
		maxX:       len(input[0]),
		maxY:       len(input),
	}
	newRadar.floodMap = make([][]bool, newRadar.maxY)
	for index := range newRadar.floodMap {
		newRadar.floodMap[index] = make([]bool, newRadar.maxX)
	}
	return &newRadar
}

func (radar *Radar) getBasinSize(x int, y int) int {
	depth, isValid := radar.GetFloor(x, y)
	if !isValid || depth >= 9 {
		return 0
	}

	if radar.floodMap[y][x] {
		return 0
	}

	// Mark as visited
	radar.floodMap[y][x] = true

	sum := 1
	for _, dir := range adjacentDirections {
		sum += radar.getBasinSize(x+dir[0], y+dir[1])
	}

	return sum
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
