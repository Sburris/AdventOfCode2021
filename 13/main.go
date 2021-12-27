package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	inputFile := "input.txt"
	coords, instructions := getInput(inputFile)
	paper := createPaper(coords)

	for _, instruction := range instructions {
		paper.RunInstructions(instruction)
	}

	paper.DrawDots()
}

/// ---------------------------------------------------------------------------
///   Paper Functions
/// --------------------------------------------------------------------------

type Paper struct {
	dots []Dot
}

func createPaper(dots []string) Paper {
	paper := Paper{dots: make([]Dot, 0, len(dots))}

	for _, dot := range dots {
		d := createDot(dot)
		paper.dots = append(paper.dots, d)
	}

	return paper
}

func (paper *Paper) FoldOnY(foldy int) {
	newDots := make([]Dot, 0, len(paper.dots))

	for _, dot := range paper.dots {
		dot.FoldOnY(foldy)
		doesContain := false
		for _, d := range newDots {
			if d == dot {
				doesContain = true
				break
			}
		}

		if !doesContain {
			newDots = append(newDots, dot)
		}

		paper.dots = newDots
	}
}

func (paper *Paper) FoldOnX(foldx int) {
	newDots := make([]Dot, 0, len(paper.dots))

	for _, dot := range paper.dots {
		dot.FoldOnX(foldx)
		doesContain := false
		for _, d := range newDots {
			if d == dot {
				doesContain = true
				break
			}
		}

		if !doesContain {
			newDots = append(newDots, dot)
		}

		paper.dots = newDots
	}
}

func (paper *Paper) RunInstructions(instructions string) {
	parts := strings.Split(instructions, " ")
	instrucs := strings.Split(parts[2], "=")

	if instrucs[0] == "x" {
		line, _ := strconv.Atoi(instrucs[1])
		paper.FoldOnX(line)
	} else {
		line, _ := strconv.Atoi(instrucs[1])
		paper.FoldOnY(line)
	}
}

func (paper *Paper) DrawDots() {
	maxX := 0
	maxY := 0

	for _, dot := range paper.dots {
		if dot.x > maxX {
			maxX = dot.x
		}
		if dot.y > maxY {
			maxY = dot.y
		}
	}

	grid := make([][]string, maxY+1)
	for index := range grid {
		grid[index] = make([]string, maxX+1)
	}

	for _, dot := range paper.dots {
		grid[dot.y][dot.x] = "#"
	}

	for _, row := range grid {
		line := ""
		for _, spot := range row {

			if len(spot) == 0 {
				line += " "
			} else {
				line += "#"
			}
		}
		fmt.Println(line)

	}
}

/// ---------------------------------------------------------------------------
///   Dot Structure
/// ---------------------------------------------------------------------------

type Dot struct {
	x int
	y int
}

func createDot(coordinate string) Dot {
	parts := strings.Split(coordinate, ",")

	x, err := strconv.Atoi(parts[0])
	checkErr(err)
	y, err := strconv.Atoi(parts[1])
	checkErr(err)

	newDot := Dot{
		x: x,
		y: y,
	}

	return newDot
}

func (dot *Dot) FoldOnY(foldy int) {
	if dot.y > foldy {
		dot.y = foldy - (dot.y - foldy)
	}
}

func (dot *Dot) FoldOnX(foldx int) {
	if dot.x > foldx {
		dot.x = foldx - (dot.x - foldx)
	}
}

/// ---------------------------------------------------------------------------
///   Helper Functions
/// ---------------------------------------------------------------------------

// Read a list of values from the given file
func getInput(filename string) ([]string, []string) {
	coordinates := make([]string, 0)
	instructions := make([]string, 0)

	file, err := os.Open(filename)
	checkErr(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	isInstructions := false
	for scanner.Scan() {
		line := scanner.Text()

		if line == "" {
			isInstructions = true
			continue
		}
		if isInstructions {
			instructions = append(instructions, line)
		} else {
			coordinates = append(coordinates, line)
		}

	}

	return coordinates, instructions
}

// Check if an error has occured
func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
