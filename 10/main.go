package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"sort"
)

func main() {
	lines := getInput("input.txt")
	sum := 0
	incompletes := make([]int, 0)
	for _, line := range lines {
		corruptScore, incompleteScore := processLine(line)
		if incompleteScore > 0 {
			incompletes = append(incompletes, incompleteScore)
		}
		sum += corruptScore
	}

	sort.Ints(incompletes)

	length := len(incompletes)

	middle := incompletes[(length-1)/2]

	fmt.Printf("Corrupt Sum: %d\n", sum)
	fmt.Printf("Incomplete Score: %d\n", middle)
}

func processLine(line string) (int, int) {
	stack := createStack()
	for _, symbol := range line {
		switch symbol {
		case '(', '[', '{', '<':
			stack.push(symbol)
		case ')':
			head, _ := stack.head()
			if head != '(' {
				return getSymbolPoints(symbol), 0
			}
			stack.pop()
		case ']':
			head, _ := stack.head()
			if head != '[' {
				return getSymbolPoints(symbol), 0
			}
			stack.pop()
		case '}':
			head, _ := stack.head()
			if head != '{' {
				return getSymbolPoints(symbol), 0
			}
			stack.pop()
		case '>':
			head, _ := stack.head()
			if head != '<' {
				return getSymbolPoints(symbol), 0
			}
			stack.pop()
		}
	}

	symbol, err := stack.pop()
	sum := 0
	for err == nil {
		value := getClosingPoints(symbol)
		sum *= 5
		sum += value

		symbol, err = stack.pop()
	}

	return 0, sum
}

func getClosingPoints(symbol rune) int {
	switch symbol {
	case '(':
		return 1
	case '[':
		return 2
	case '{':
		return 3
	case '<':
		return 4
	}
	return 0
}

func getSymbolPoints(symbol rune) int {
	switch symbol {
	case ')':
		return 3
	case ']':
		return 57
	case '}':
		return 1197
	case '>':
		return 25137
	}
	return 0
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

type Stack struct {
	buffer []rune
}

func (stack *Stack) push(symbol rune) {
	stack.buffer = append(stack.buffer, symbol)
}

func (stack *Stack) pop() (rune, error) {
	if len(stack.buffer) == 0 {
		return ' ', errors.New("stack empty")
	}

	element := stack.buffer[len(stack.buffer)-1]
	stack.buffer = stack.buffer[:len(stack.buffer)-1]
	return element, nil
}

func (stack *Stack) head() (rune, error) {
	if len(stack.buffer) == 0 {
		return ' ', errors.New("stack empty")
	}

	return stack.buffer[len(stack.buffer)-1], nil
}

func createStack() *Stack {
	return &Stack{
		buffer: make([]rune, 0, 100),
	}
}
