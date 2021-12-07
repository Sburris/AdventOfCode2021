package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const TEST_INPUT_FILE string = "test_data/sample_input.txt"

/// ---------------------------------------------------------------------------
///   addNewFish
/// ---------------------------------------------------------------------------

func Test_addNewFish_OneFish(t *testing.T) {
	inputSchool := make([]*schoolOfFish, 0, 1)
	inputFish := 1
	expected := []schoolOfFish{
		{timer: 1, count: 1},
	}

	output := addNewFish(inputSchool, inputFish)

	assert.Equal(t, len(expected), len(output), "they should match")
}

func Test_addNewFish_TwoFishSameTimer(t *testing.T) {
	inputSchool := make([]*schoolOfFish, 0, 1)
	inputFish := 1
	expected := []schoolOfFish{
		{timer: 1, count: 2},
	}

	output := addNewFish(inputSchool, inputFish)
	output = addNewFish(output, inputFish)

	assert.Equal(t, len(expected), len(output), "they should match")
}
/// ---------------------------------------------------------------------------
///   processInput
/// ---------------------------------------------------------------------------

// func Test_process_SampleInput(t *testing.T) {
// 	input := "3,4,3,1,2"
// 	expected := []*schoolOfFish {
// 		&schoolOfFish{timer: 3, count: 2},
// 		&schoolOfFish{timer: 4, count: 1},
// 		&schoolOfFish{timer: 1, count: 1},
// 		&schoolOfFish{timer: 2, count: 1},
// 	}

// 	output := processInput(input)

// 	assert.Equal(t, expected, output, "input file not parsed correctly")
// }


/// ---------------------------------------------------------------------------
///   getInput
/// ---------------------------------------------------------------------------

func Test_getInput_CanProperlyRead(t *testing.T) {
	expected := "3,4,3,1,2"

	output := getInput(TEST_INPUT_FILE)

	assert.Equal(t, expected, output, "input file not parsed correctly")
}
