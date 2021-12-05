package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const TEST_INPUT_FILE string = "test_data/sample_input.txt"

/// ---------------------------------------------------------------------------
///   calculateDangerousPoints
/// ---------------------------------------------------------------------------

func Test_calculateDangerousPoints_Simple(t *testing.T) {
	rawInput := getInput(TEST_INPUT_FILE)
	ventsInput := processVentInput(rawInput)
	ventMap := calculateVentMap(ventsInput, true)
	expected := 5

	output := calculateDangerousPoints(ventMap)

	assert.Equal(t, expected, output, "should match")
}

func Test_calculateDangerousPoints_Complex(t *testing.T) {
	rawInput := getInput(TEST_INPUT_FILE)
	ventsInput := processVentInput(rawInput)
	ventMap := calculateVentMap(ventsInput, false)
	expected := 12

	output := calculateDangerousPoints(ventMap)

	assert.Equal(t, expected, output, "should match")
}

/// ---------------------------------------------------------------------------
///   calculateSlope
/// ---------------------------------------------------------------------------

func Test_calculateSlope_horizontalVent(t *testing.T) {
	input := newVent("0,9 -> 5,9")
	expected := 0

	output, err := calculateSlope(input)

	assert.Nil(t, err, "no error should be thrown")
	assert.Equal(t, expected, output, "should match")
}

func Test_calculateSlope_positiveSlope(t *testing.T) {
	input := newVent("0,0 -> 8,8")
	expected := 1

	output, err := calculateSlope(input)

	assert.Nil(t, err, "no error should be thrown")
	assert.Equal(t, expected, output, "should match")
}

func Test_calculateSlope_negativeSlope(t *testing.T) {
	input := newVent("8,0 -> 0,8")
	expected := -1

	output, err := calculateSlope(input)

	assert.Nil(t, err, "no error should be thrown")
	assert.Equal(t, expected, output, "should match")
}

func Test_calculateSlope_verticalVent(t *testing.T) {
	input := newVent("2,2 -> 2,1")

	_, err := calculateSlope(input)

	assert.Equal(t, ErrVerticalLine, err, "should match")
}

func Test_calculateSlope_singlePoint(t *testing.T) {
	input := newVent("2,2 -> 2,2")
	expected := 0

	output, err := calculateSlope(input)

	assert.Nil(t, err, "no error should be thrown")
	assert.Equal(t, expected, output, "should match")
}

/// ---------------------------------------------------------------------------
///   calculateVentMap
/// ---------------------------------------------------------------------------

func Test_calculateVentMap_SampleInput(t *testing.T) {
	rawInput := getInput(TEST_INPUT_FILE)
	ventsInput := processVentInput(rawInput)
	expected := [][]int {
		{0,0,0,0,0,0,0,1,0,0},
		{0,0,1,0,0,0,0,1,0,0},
		{0,0,1,0,0,0,0,1,0,0},
		{0,0,0,0,0,0,0,1,0,0},
		{0,1,1,2,1,1,1,2,1,1},
		{0,0,0,0,0,0,0,0,0,0},
		{0,0,0,0,0,0,0,0,0,0},
		{0,0,0,0,0,0,0,0,0,0},
		{0,0,0,0,0,0,0,0,0,0},
		{2,2,2,1,1,1,0,0,0,0},
	}

	output := calculateVentMap(ventsInput, true)

	assert.Equal(t, expected, output, "should match")
}

func Test_calculateVentMap_ComplexInput(t *testing.T) {
	rawInput := getInput(TEST_INPUT_FILE)
	ventsInput := processVentInput(rawInput)
	expected := [][]int {
		{1,0,1,0,0,0,0,1,1,0},
		{0,1,1,1,0,0,0,2,0,0},
		{0,0,2,0,1,0,1,1,1,0},
		{0,0,0,1,0,2,0,2,0,0},
		{0,1,1,2,3,1,3,2,1,1},
		{0,0,0,1,0,2,0,0,0,0},
		{0,0,1,0,0,0,1,0,0,0},
		{0,1,0,0,0,0,0,1,0,0},
		{1,0,0,0,0,0,0,0,1,0},
		{2,2,2,1,1,1,0,0,0,0},
	}

	output := calculateVentMap(ventsInput, false)

	assert.Equal(t, expected, output, "should match")
}

/// ---------------------------------------------------------------------------
///   getInput
/// ---------------------------------------------------------------------------

func Test_getInput_CanProperlyRead(t *testing.T) {
	expected := []string {
		"0,9 -> 5,9",
		"8,0 -> 0,8",
		"9,4 -> 3,4",
		"2,2 -> 2,1",
		"7,0 -> 7,4",
		"6,4 -> 2,0",
		"0,9 -> 2,9",
		"3,4 -> 1,4",
		"0,0 -> 8,8",
		"5,5 -> 8,2",
	}

	output := getInput(TEST_INPUT_FILE)

	assert.Equal(t, expected, output, "should match")
}

/// ---------------------------------------------------------------------------
///   newVent
/// ---------------------------------------------------------------------------

func Test_newVent_sampleInputOne(t *testing.T) {
	input := "0,9 -> 5,9"

	expected := &vent{
		x1: 0,
		y1: 9,
		x2: 5,
		y2: 9,
	}

	output := newVent(input)

	assert.Equal(t, expected, output, "should match")
}

func Test_newVent_sampleInputTwo(t *testing.T) {
	input := "0,0 -> 8,8"

	expected := &vent{
		x1: 0,
		y1: 0,
		x2: 8,
		y2: 8,
	}

	output := newVent(input)

	assert.Equal(t, expected, output, "should match")
}

/// ---------------------------------------------------------------------------
///   printVentMap
/// ---------------------------------------------------------------------------

func Test_printVentMap_sampleInput(t *testing.T) {
	rawData := getInput(TEST_INPUT_FILE)
	vents := processVentInput(rawData)
	input := calculateVentMap(vents, true)
	expect := []string {
		".......1..",
		"..1....1..",
		"..1....1..",
		".......1..",
		".112111211",
		"..........",
		"..........",
		"..........",
		"..........",
		"222111....",
	}

	output := printVentMap(input)

	assert.Equal(t, expect, output, "should match")
}

/// ---------------------------------------------------------------------------
///   processVentInput
/// ---------------------------------------------------------------------------

func Test_processVentInput_sampleInput(t *testing.T) {
	input := []string {
		"0,9 -> 5,9",
		"8,0 -> 0,8",
		"9,4 -> 3,4",
	}
	expected := []*vent {
		{x1: 0, y1: 9, x2: 5, y2: 9},
		{x1: 8, y1: 0, x2: 0, y2: 8},
		{x1: 9, y1: 4, x2: 3, y2: 4},
	}

	output := processVentInput(input)

	assert.Equal(t, expected, output, "should match")
}

/// ---------------------------------------------------------------------------
///   processVentCoordinate
/// ---------------------------------------------------------------------------

func Test_processVentCoordinate_sampleInputOne(t *testing.T) {
	input := "0,0"
	expectedX1 := 0
	expectedY1 := 0

	outputX1, outputY1 := processVentCoordinate(input)

	assert.Equal(t, expectedX1, outputX1, "should match")
	assert.Equal(t, expectedY1, outputY1, "should match")
}