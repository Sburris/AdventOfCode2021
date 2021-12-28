package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	inputFile := "input_test.txt"
	//inputFile := "input_test_1.txt"
	//inputFile := "input.txt"

	chain, pairs := getInput(inputFile)
	polymer := createPolymer(chain, pairs)

	steps := 40
	polymer.ProcessSteps(steps)
	// for count := 0; count < steps; count++ {
	// 	polymer.Process()
	// 	fmt.Printf("Step: %d (%d)\n", count, len(polymer.chain))
	fmt.Println(polymer.chain)
	// }

	counts := countElements(polymer.chain)

	minCount := 99999
	minElement := "NONE"
	maxCount := 0
	maxElement := "NONE"

	for key, value := range counts {
		if value < minCount {
			minCount = value
			minElement = string(key)
		}

		if value > maxCount {
			maxCount = value
			maxElement = string(key)
		}
	}

	fmt.Printf("Min: %s - %d\n", minElement, minCount)
	fmt.Printf("Max: %s - %d\n", maxElement, maxCount)
	fmt.Printf("Diff: %d\n", maxCount-minCount)
}

func countElements(chain string) map[rune]int {
	elements := make(map[rune]int)

	for _, element := range chain {
		_, found := elements[element]
		if found {
			elements[element]++
		} else {
			elements[element] = 1
		}
	}

	return elements
}

type Polymer struct {
	chain       string
	pairInserts map[string]string
}

func createPolymer(chain string, pairInserts []string) Polymer {
	polymer := Polymer{
		chain: chain,
	}

	polymer.pairInserts = make(map[string]string)
	for _, pair := range pairInserts {
		parts := strings.Split(pair, " -> ")
		polymer.pairInserts[parts[0]] = parts[1]
	}

	return polymer
}

func (polymer *Polymer) ProcessSteps(steps int) {
	lookup := make(map[string]string)

	for key := range polymer.pairInserts {
		poly := Polymer{chain: key, pairInserts: polymer.pairInserts}
		for count := 0; count < 10; count++ {
			poly.Process()
		}

		lookup[key] = poly.chain
	}

	for count := 0; count < steps; count += 10 {
		newChain := string(polymer.chain[0])
		for index := 0; index < len(polymer.chain)-1; index++ {
			pair := polymer.chain[index : index+2]
			value, found := lookup[pair]
			if found {
				newChain += value[1:]
			} else {
				newChain += string(pair[1])
			}
		}
		polymer.chain = newChain
		fmt.Printf("len: %d\n", len(newChain))
	}
}

func (polymer *Polymer) Process() {
	newChain := string(polymer.chain[0])
	for index := 0; index < len(polymer.chain)-1; index++ {
		pair := polymer.chain[index : index+2]
		value, found := polymer.pairInserts[pair]
		if found {
			newChain += value + polymer.chain[index+1:index+2]
		} else {
			fmt.Printf("no pair found: %s\n", pair)
			newChain += polymer.chain[index+1 : index+2]
		}
	}
	polymer.chain = newChain
}

// Read a list of values from the given file
func getInput(filename string) (string, []string) {
	results := make([]string, 0)

	file, err := os.Open(filename)
	checkErr(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	chain := scanner.Text()

	scanner.Scan() // blankLine
	for scanner.Scan() {
		results = append(results, scanner.Text())
	}

	return chain, results
}

// Check if an error has occured
func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
