package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const TEST_INPUT_FILE string = "test_data/sample_input.txt"

/// ---------------------------------------------------------------------------
///   countIncreases
/// ---------------------------------------------------------------------------

func Test_countIncreases_emptyInput(t *testing.T) {
	input := []int{}
	expected := 0

	output := countIncreases(input)

	assert.Equal(t, expected, output, "should match")
}

func Test_countIncreases_allTheSame(t *testing.T) {
	input := []int{1, 1, 1, 1}
	expected := 0

	output := countIncreases(input)

	assert.Equal(t, expected, output, "should match")
}

func Test_countIncreases_onlyDecreases(t *testing.T) {
	input := []int{5, 4, 3, 2, 1}
	expected := 0

	output := countIncreases(input)

	assert.Equal(t, expected, output, "should match")
}

func Test_countIncreases_onlyIncreases(t *testing.T) {
	input := []int{1, 2, 3, 4, 5}
	expected := 4

	output := countIncreases(input)

	assert.Equal(t, expected, output, "should match")
}

func Test_countIncreases_allCases(t *testing.T) {
	input := []int{1, 2, 3, 3, 3, 2, 1}
	expected := 2

	output := countIncreases(input)

	assert.Equal(t, expected, output, "should match")
}

/// ---------------------------------------------------------------------------
///   createWindows
/// ---------------------------------------------------------------------------

func Test_createWindows_emptyList(t *testing.T) {
	input := []int{}
	expected := []int{}

	output := createWindows(input, 3)

	assert.Equal(t, expected, output, "should match")
}

func Test_createWindows_negativeSize(t *testing.T) {
	input := []int{1, 2, 3, 3, 3, 2, 1}
	expected := []int{}

	output := createWindows(input, -1)

	assert.Equal(t, expected, output, "should match")
}

func Test_createWindows_sizeLargerThanInput(t *testing.T) {
	input := []int{1, 2, 3, 3, 3, 2, 1}
	expected := []int{}

	output := createWindows(input, 10)

	assert.Equal(t, expected, output, "should match")
}

func Test_createWindows_sizeMatchesInputCount(t *testing.T) {
	input := []int{1, 2, 3}
	expected := []int{6}

	output := createWindows(input, 3)

	assert.Equal(t, expected, output, "should match")
}

func Test_createWindows_size0(t *testing.T) {
	input := []int{1, 2, 3, 3, 3, 2, 1}
	expected := []int{}

	output := createWindows(input, 0)

	assert.Equal(t, expected, output, "should match")
}

func Test_createWindows_size1(t *testing.T) {
	input := []int{1, 2, 3, 3, 3, 2, 1}
	expected := []int{1, 2, 3, 3, 3, 2, 1}

	output := createWindows(input, 1)

	assert.Equal(t, expected, output, "should match")
}

func Test_createWindows_size3(t *testing.T) {
	input := []int{1, 2, 3, 3, 3, 2, 1}
	expected := []int{6, 8, 9, 8, 6}

	output := createWindows(input, 3)

	assert.Equal(t, expected, output, "should match")
}

func Test_createWindows_size2(t *testing.T) {
	input := []int{1, 2, 3, 3, 3, 2, 1}
	expected := []int{3, 5, 6, 6, 5, 3}

	output := createWindows(input, 2)

	assert.Equal(t, expected, output, "should match")
}

/// ---------------------------------------------------------------------------
///   getInput
/// ---------------------------------------------------------------------------

func Test_getData_CanProperlyRead(t *testing.T) {
	expected := []int{198, 208, 209, 212, 213, 217, 218, 223, 222, 224}

	output := getInput(TEST_INPUT_FILE)

	assert.Equal(t, expected, output, "input file not parsed correctly")
}
