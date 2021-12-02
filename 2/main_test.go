package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const TEST_INPUT_FILE string = "test_data/sample_input.txt"

/// ---------------------------------------------------------------------------
///   submarine Move
/// ---------------------------------------------------------------------------

func Test_submarine_move_down1(t *testing.T) {
	input := "down 1"
	sub := newDefaultSubmarine()
	expected := newSubmarine(0, 0, 1)

	sub.move(input)

	assert.Equal(t, expected, sub, "should be at the same point")
}

func Test_coordinate_move_up1(t *testing.T) {
	input := "up 1"
	sub := newDefaultSubmarine()
	expected := newSubmarine(0, 0, -1)

	sub.move(input)

	assert.Equal(t, expected, sub, "should be at the same point")
}

func Test_coordinate_move_forward1(t *testing.T) {
	input := "forward 1"
	sub := newDefaultSubmarine()
	expected := newSubmarine(1, 0, 0)

	sub.move(input)

	assert.Equal(t, expected, sub, "should be at the same point")
}

func Test_coordinate_move_forward1WithAimStartingDown5(t *testing.T) {
	input := "forward 1"
	coord := newSubmarine(0,5,5)
	expected := newSubmarine(1, 10, 5)

	coord.move(input)

	assert.Equal(t, expected, coord, "should be at the same point")
}

func Test_coordinate_move_demoCommandStep1(t *testing.T) {
	input := "forward 5"
	coord := newSubmarine(0,0,0)
	expected := newSubmarine(5, 0, 0)

	coord.move(input)

	assert.Equal(t, expected, coord, "should be at the same point")
}

func Test_coordinate_move_demoCommandStep2(t *testing.T) {
	input := "down 5"
	coord := newSubmarine(5,0,0)
	expected := newSubmarine(5, 0, 5)

	coord.move(input)

	assert.Equal(t, expected, coord, "should be at the same point")
}

func Test_coordinate_move_demoCommandStep3(t *testing.T) {
	input := "forward 8"
	coord := newSubmarine(5,0,5)
	expected := newSubmarine(13, 40, 5)

	coord.move(input)

	assert.Equal(t, expected, coord, "should be at the same point")
}

func Test_coordinate_move_demoCommandStep4(t *testing.T) {
	input := "up 3"
	coord := newSubmarine(13, 40, 5)
	expected := newSubmarine(13, 40, 2)

	coord.move(input)

	assert.Equal(t, expected, coord, "should be at the same point")
}

func Test_coordinate_move_demoCommandStep5(t *testing.T) {
	input := "down 8"
	coord := newSubmarine(13, 40, 2)
	expected := newSubmarine(13, 40, 10)

	coord.move(input)

	assert.Equal(t, expected, coord, "should be at the same point")
}

func Test_coordinate_move_demoCommandStep6(t *testing.T) {
	input := "forward 2"
	coord := newSubmarine(13, 40, 10)
	expected := newSubmarine(15, 60, 10)

	coord.move(input)

	assert.Equal(t, expected, coord, "should be at the same point")
}

/// ---------------------------------------------------------------------------
///   getInput
/// ---------------------------------------------------------------------------

func Test_getData_CanProperlyRead(t *testing.T) {
	expected := []string{
			"forward 8", 
			"forward 9", 
			"forward 9", 
			"down 3", 
			"forward 9", 
			"down 1", 
			"down 7", 
			"down 7", 
			"down 4", 
			"down 2",
		}

	output := getInput(TEST_INPUT_FILE)

	assert.Equal(t, expected, output, "input file not parsed correctly")
}










