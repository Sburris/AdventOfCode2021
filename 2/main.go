package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type submarine struct {
	x int
	y int
	aim int
}

func main() {
	inputFile := "input.txt"
	input := getInput(inputFile)
	sub := newDefaultSubmarine()
	for _, command := range input {
		sub.move(command)
	}
	fmt.Printf("Sub position: %v\n", sub)
	fmt.Printf("Answer: %d\n", sub.x * sub.y)
}

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

func newDefaultSubmarine() *submarine {
	return &submarine{x:0, y: 0, aim: 0}
}

func newSubmarine(x int, y int, aim int) *submarine {
	return &submarine{x: x, y: y, aim: aim}
}

func (this *submarine) move(command string) {
	values := strings.Split(command, " ")
	direction := values[0]
	distance, err := strconv.Atoi(values[1])
	checkErr(err)
	switch direction {
	case "up":
		this.aim -= distance
	case "down":
		this.aim += distance
	case "forward":
		this.x += distance
		this.y += this.aim * distance
	default:
		panic(fmt.Sprintf("Unknown command: %s\n", command))
	}
}
