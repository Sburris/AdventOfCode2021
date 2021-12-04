package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const TEST_INPUT_FILE string = "test_data/sample_input.txt"

/// ---------------------------------------------------------------------------
///   bingoGame - drawNumber
/// ---------------------------------------------------------------------------

func Test_bingoGame_drawNumber_once(t *testing.T) {
	game := bingoGame {
		numberDrawn: -1,
		numberOrder: []int {1, 2, 3, 0},
		isNumberDrawn: []bool {false, false, false, false},
	}

	expected := bingoGame {
		numberDrawn: 0,
		numberOrder: []int {1, 2, 3, 0},
		isNumberDrawn: []bool {false, true, false, false},
	}

	game.drawNumber()

	assert.Equal(t, expected, game, "should match")
}

func Test_bingoGame_drawNumber_threeDraws(t *testing.T) {
	game := bingoGame {
		numberDrawn: -1,
		numberOrder: []int {1, 2, 3, 0},
		isNumberDrawn: []bool {false, false, false, false},
	}

	expected := bingoGame {
		numberDrawn: 2,
		numberOrder: []int {1, 2, 3, 0},
		isNumberDrawn: []bool {false, true, true, true},
	}

	game.drawNumber()
	game.drawNumber()
	game.drawNumber()

	assert.Equal(t, expected, game, "should match")
}

/// ---------------------------------------------------------------------------
///   getInput
/// ---------------------------------------------------------------------------

func Test_getData_CanProperlyRead(t *testing.T) {
	board1 := board{ squares: [5][5]int {
		{22, 13, 17, 11,  0},
		{ 8,  2, 23,  4, 24},
		{21,  9, 14, 16,  7},
		{ 6, 10,  3, 18,  5},
		{ 1, 12, 20, 15, 19},
	}}

	board2 := board{ squares: [5][5]int {
		{ 3, 15,  0,  2, 22},
		{ 9, 18, 13, 17,  5},
		{19,  8,  7, 25, 23},
		{20, 11, 10, 24,  4},
		{14, 21, 16, 12,  6},
	}}

	board3 := board{ squares: [5][5]int {
		{14, 21, 17, 24,  4},
		{10, 16, 15,  9, 19},
		{18,  8, 23, 26, 20},
		{22, 11, 13,  6,  5},
		{ 2,  0, 12,  3,  7},
	}}

	expected := bingoGame{
		numberDrawn: -1,
		numberOrder: []int{7,4,9,5,11,17,23,2,0,14,21,24,10,16,13,6,15,25,12,22,18,20,8,19,3,26,1},
		isNumberDrawn: []bool{false,false,false,false,false,false,false,false,false,false,false,false,false,false,false,false,false,false,false,false,false,false,false,false,false,false,false,},
		boardScore: []int {-1,-1,-1},
		boards: []*board{ &board1, &board2, &board3},
	}

	output := getInput(TEST_INPUT_FILE)

	assert.Equal(t, expected, output, "input file not parsed correctly")
}

