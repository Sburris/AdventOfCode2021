package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

/// ---------------------------------------------------------------------------
///   createDot
/// ---------------------------------------------------------------------------

func Test_createDot(t *testing.T) {
	input := "1,1"

	expected := Dot{x: 1, y: 1}

	output := createDot(input)

	assert.Equal(t, expected, output)
}

/// ---------------------------------------------------------------------------
///   createPaper
/// ---------------------------------------------------------------------------

func Test_createPaper(t *testing.T) {
	input := []string{
		"6,10",
		"0,14",
		"9,10",
		"0,3",
		"10,4",
		"4,11",
		"6,0",
		"6,12",
		"4,1",
		"0,13",
		"10,12",
		"3,4",
		"3,0",
		"8,4",
		"1,10",
		"2,14",
		"8,10",
		"9,0",
	}

	expected := &Paper{
		dots: []Dot{
			createDot("6,10"),
			createDot("0,14"),
			createDot("9,10"),
			createDot("0,3"),
			createDot("10,4"),
			createDot("4,11"),
			createDot("6,0"),
			createDot("6,12"),
			createDot("4,1"),
			createDot("0,13"),
			createDot("10,12"),
			createDot("3,4"),
			createDot("3,0"),
			createDot("8,4"),
			createDot("1,10"),
			createDot("2,14"),
			createDot("8,10"),
			createDot("9,0"),
		},
	}

	output := createPaper(input)

	assert.ElementsMatch(t, expected.dots, output.dots)
}

/// ---------------------------------------------------------------------------
///   Dot FoldOnX
/// ---------------------------------------------------------------------------

func Test_Dot_FoldOnX_NoChange(t *testing.T) {
	dot := createDot("0,0")
	input := 7
	expected := createDot("0,0")

	dot.FoldOnX(input)

	assert.Equal(t, expected, dot)
}

func Test_Dot_FoldOnX_Change(t *testing.T) {
	dot := createDot("10,2")
	input := 5
	expected := createDot("0,2")

	dot.FoldOnX(input)

	assert.Equal(t, expected, dot)
}

func Test_Dot_FoldOnX_Change2(t *testing.T) {
	dot := createDot("6,2")
	input := 5
	expected := createDot("4,2")

	dot.FoldOnX(input)

	assert.Equal(t, expected, dot)
}

/// ---------------------------------------------------------------------------
///   Dot FoldOnY
/// ---------------------------------------------------------------------------

func Test_Dot_FoldOnY_NoChange(t *testing.T) {
	dot := createDot("0,0")
	input := 7
	expected := createDot("0,0")

	dot.FoldOnY(input)

	assert.Equal(t, expected, dot)
}

func Test_Dot_FoldOnY_Change(t *testing.T) {
	dot := createDot("0,14")
	input := 7
	expected := createDot("0,0")

	dot.FoldOnY(input)

	assert.Equal(t, expected, dot)
}

/// ---------------------------------------------------------------------------
///   Paper FoldOnY
/// ---------------------------------------------------------------------------

func Test_Paper_FoldOnY(t *testing.T) {
	inputCoords := []string{
		"6,10",
		"0,14",
		"9,10",
		"0,3",
		"10,4",
		"4,11",
		"6,0",
		"6,12",
		"4,1",
		"0,13",
		"10,12",
		"3,4",
		"3,0",
		"8,4",
		"1,10",
		"2,14",
		"8,10",
		"9,0",
	}

	expectedCoords := []string{
		"6,4",
		"0,0",
		"9,4",
		"0,3",
		"10,4",
		"4,3",
		"6,0",
		"6,2",
		"4,1",
		"0,1",
		"10,2",
		"3,4",
		"3,0",
		"1,4",
		"2,0",
		"8,4",
		"9,0",
	}

	paper := createPaper(inputCoords)
	expectedPaper := createPaper(expectedCoords)

	paper.FoldOnY(7)

	assert.ElementsMatch(t, expectedPaper.dots, paper.dots)
}

/// ---------------------------------------------------------------------------
///   getInput
/// ---------------------------------------------------------------------------

func Test_getInput(t *testing.T) {
	filename := "input_test.txt"

	expectCoords := []string{
		"6,10",
		"0,14",
		"9,10",
		"0,3",
		"10,4",
		"4,11",
		"6,0",
		"6,12",
		"4,1",
		"0,13",
		"10,12",
		"3,4",
		"3,0",
		"8,4",
		"1,10",
		"2,14",
		"8,10",
		"9,0",
	}

	expectedInstruction := []string{
		"fold along y=7",
	}

	outputCoords, outputInstruction := getInput(filename)

	assert.ElementsMatch(t, expectCoords, outputCoords)
	assert.ElementsMatch(t, expectedInstruction, outputInstruction)
}
