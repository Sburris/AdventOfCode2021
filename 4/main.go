package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

/// ---------------------------------------------------------------------------
///   Custom Types
/// ---------------------------------------------------------------------------

type bingoGame struct {
	numberDrawn int
	numberOrder []int
	isNumberDrawn []bool
	boards []*board
	boardWinningOrder []int
	boardScore []int
}

type board struct {
	squares [5][5]int
}

/// ---------------------------------------------------------------------------
///   Main Function
/// ---------------------------------------------------------------------------

func main() {
	inputFile := "input.txt"
	//inputFile := "test_data/sample_input.txt"
	game := getInput(inputFile)

	// play rounds until all boards have one
	totalBoards := len(game.boards)
	for len(game.boardWinningOrder) < totalBoards {
		game.PlayRound()
	}

	// First winner
	firstBoardIndex := game.boardWinningOrder[0]
	fmt.Print("\nWinning Board:\n")
	fmt.Printf("Id: %d\n", firstBoardIndex)
	fmt.Printf("Score: %d\n", game.boardScore[firstBoardIndex])

	// Last Inner
	lastBoardIndex := game.boardWinningOrder[totalBoards-1]
	fmt.Print("\nLosing Loser Board:\n")
	fmt.Printf("Id: %d\n", lastBoardIndex)
	fmt.Printf("Score: %d\n", game.boardScore[lastBoardIndex])
}

/// ---------------------------------------------------------------------------
///   Custom Type functions
/// ---------------------------------------------------------------------------

// Plays a round of bingo:
//  - draws next number
//  - checks all boards that have not won yet to see if they one
func (this *bingoGame) PlayRound() {
	this.drawNumber()

	for index := range this.boards {
		// Check to see if it already won
		if this.boardScore[index] > -1 {
			continue
		}

		didWin := this.didBoardWin(index)
		if didWin {
			sum := this.CalculateBoardSum(index)
			lastDrawnNumber := this.numberOrder[this.numberDrawn]
			score := sum * lastDrawnNumber
			this.boardScore[index] = score
			this.boardWinningOrder = append(this.boardWinningOrder, index)
		}
	}
}

// Pulls the next number to be drawn
func (this *bingoGame) drawNumber() {
	this.numberDrawn++
	numDrawn := this.numberOrder[this.numberDrawn]
	this.isNumberDrawn[numDrawn] = true
}

// Determines if a board has 5 numbers in a row (column or row)
func (this *bingoGame) didBoardWin(boardIndex int) bool {
	board := this.boards[boardIndex]

	for colRowIndex :=0; colRowIndex < 5; colRowIndex++ {
		// count of marked numbers
		rowCount := 0
		columnCount := 0

		for squareIndex := 0; squareIndex < 5; squareIndex++ {
			// Check Column
			columnSquareValue := board.squares[squareIndex][colRowIndex]
			if this.isNumberDrawn[columnSquareValue] {
				columnCount++ 
			}

			// Check Row
			rowSquareValue := board.squares[colRowIndex][squareIndex]
			if this.isNumberDrawn[rowSquareValue] {
				rowCount++
			}
		}

		// Check if a board has a completed row or column
		if rowCount == 5 || columnCount == 5 {
			return true // Winner
		}
	}
	return false // Loser
}

// Calculates the sum of all squares that have not been marked
func (this *bingoGame) CalculateBoardSum(boardIndex int) int {
	board := this.boards[boardIndex]

	sum := 0
	for rowIndex := 0; rowIndex < 5; rowIndex++ {
		for columnIndex := 0; columnIndex < 5; columnIndex++ {
			squareValue := board.squares[rowIndex][columnIndex]
			
			wasNotDrawn := !this.isNumberDrawn[squareValue]
			if wasNotDrawn {
				sum += squareValue
			}
		}
	}

	return sum
}

/// ---------------------------------------------------------------------------
///   General Functions
/// ---------------------------------------------------------------------------

// Creates a new bingo board.
// Expects an array of 5 strings each 5 with integers seperated by spaces
func newBoard(input []string) *board {
	board := board{ squares: [5][5]int{}}
	if len(input) != 5 {
		panic("incorrect board size")
	}

	for row := 0; row < 5; row++ {
		numbers := strings.Split(input[row], " ")
		column := 0
		for _, val := range numbers {
			
			// Handle the extra spaces between numbers 
			// that allow the input to look properly formatted
			if strings.TrimSpace(val) == "" {
				continue
			}
			
			value, err := strconv.Atoi(val)
			checkErr(err)
			board.squares[row][column] = value
			column++
		}
	}

	return &board
}

// Read in raw data so we can process it
func getInput(filename string) bingoGame {
	game := bingoGame{
		numberDrawn: -1, // because we haven't drawn the first ball yet which starts at zero
	}

	file, err := os.Open(filename)
	checkErr(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)

	// Process drawn numbers
	scanner.Scan()
	numberLine := scanner.Text()
	numbers := strings.Split(numberLine, ",")
	numberCount := len(numbers)
	game.numberOrder = make([]int, 0, numberCount)
	game.isNumberDrawn = make([]bool, 0, numberCount)
	for _, num := range numbers {
		value, err := strconv.Atoi(num)
		checkErr(err)
		game.numberOrder = append(game.numberOrder, value)
		game.isNumberDrawn = append(game.isNumberDrawn, false)
	}

	// Read all boards
	boardLines := []string{}
	for scanner.Scan() {
		boardLines = append(boardLines, scanner.Text())
	}

	boardCount := len(boardLines) / 6
	game.boardScore = make([]int, 0)
	fmt.Printf("boardCount: %d\n", boardCount)
	for count := 0; count < boardCount; count++ {
		offset := 6 * count
		newBoard := newBoard(boardLines[offset+1:offset+6])
		game.boards = append(game.boards, newBoard)
		game.boardScore = append(game.boardScore, -1)
	}

	return game
}

// Check if an error has occured
func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}