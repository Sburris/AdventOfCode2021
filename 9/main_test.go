package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

/// ---------------------------------------------------------------------------
///   createRadar
/// ---------------------------------------------------------------------------

// func Test_createRadar(t *testing.T) {
// 	input := []string{
// 		"2199943210",
// 		"3987894921",
// 		"9856789892",
// 		"8767896789",
// 		"9899965678",
// 	}
// 	expected := &Radar{
// 		floorDepth: []string{
// 			"2199943210",
// 			"3987894921",
// 			"9856789892",
// 			"8767896789",
// 			"9899965678",
// 		},
// 		maxX: 10,
// 		maxY: 5,
// 	}
// 	output := createRadar(input)

// 	assert.Equal(t, expected, output)
// }

/// ---------------------------------------------------------------------------
///   GetAllBasinSizes
/// ---------------------------------------------------------------------------

func Test_GetAllBasinSizes(t *testing.T) {
	radar := getDefaultRadar()
	expected := []int{3, 9, 14, 9}

	output := radar.GetAllBasinSizes()

	assert.Equal(t, expected, output)
}

/// ---------------------------------------------------------------------------
///   getAllLowSpotDepths
/// ---------------------------------------------------------------------------

func Test_getAllLowSpotDepths(t *testing.T) {
	radar := getDefaultRadar()
	expected := []int{1, 0, 5, 5}

	output := radar.GetAllLowSpotDepths()

	assert.Equal(t, expected, output)
}

/// ---------------------------------------------------------------------------
///   getBasinSize
/// ---------------------------------------------------------------------------

func Test_getBasinSize_topLeft(t *testing.T) {
	radar := getDefaultRadar()
	inputX := 1
	inputY := 0
	expected := 3

	output := radar.getBasinSize(inputX, inputY)

	assert.Equal(t, expected, output)
}

func Test_getBasinSize_topRight(t *testing.T) {
	radar := getDefaultRadar()
	inputX := 9
	inputY := 0
	expected := 9

	output := radar.getBasinSize(inputX, inputY)

	assert.Equal(t, expected, output)
}

/// ---------------------------------------------------------------------------
///   getFloor
/// ---------------------------------------------------------------------------

func Test_getFloor_valid(t *testing.T) {
	radar := getDefaultRadar()
	expectedFound := true
	expectedDepth := 9

	outputDepth, outputFound := radar.GetFloor(1, 1)

	assert.Equal(t, expectedFound, outputFound)
	assert.Equal(t, expectedDepth, outputDepth)
}

func Test_getFloor_invalid(t *testing.T) {
	radar := getDefaultRadar()
	expectedFound := false
	expectedDepth := -1

	outputDepth, outputFound := radar.GetFloor(10, 10)

	assert.Equal(t, expectedFound, outputFound)
	assert.Equal(t, expectedDepth, outputDepth)
}

/// ---------------------------------------------------------------------------
///   getInput
/// ---------------------------------------------------------------------------

func Test_getData_CanProperlyRead(t *testing.T) {
	input := "test_input.txt"
	expected := []string{
		"2199943210",
		"3987894921",
		"9856789892",
		"8767896789",
		"9899965678",
	}

	output := getInput(input)

	assert.Equal(t, expected, output, "input file not parsed correctly")
}

/// ---------------------------------------------------------------------------
///   IsLowSpot
/// ---------------------------------------------------------------------------

func Test_IsLowSpot_corner_false(t *testing.T) {
	radar := getDefaultRadar()
	inputX := 0
	inputY := 0
	expectedDepth := 2
	expectedFound := false

	outputDepth, outputFound := radar.IsLowSpot(inputX, inputY)

	assert.Equal(t, expectedDepth, outputDepth)
	assert.Equal(t, expectedFound, outputFound)
}

func Test_IsLowSpot_edge_true(t *testing.T) {
	radar := getDefaultRadar()
	inputX := 1
	inputY := 0
	expectedDepth := 1
	expectedFound := true

	outputDepth, outputFound := radar.IsLowSpot(inputX, inputY)

	assert.Equal(t, expectedDepth, outputDepth)
	assert.Equal(t, expectedFound, outputFound)
}

func Test_IsLowSpot_middle_true(t *testing.T) {
	radar := getDefaultRadar()
	inputX := 2
	inputY := 2
	expectedDepth := 5
	expectedFound := true

	outputDepth, outputFound := radar.IsLowSpot(inputX, inputY)

	assert.Equal(t, expectedDepth, outputDepth)
	assert.Equal(t, expectedFound, outputFound)
}

func Test_IsLowSpot_middle_false(t *testing.T) {
	radar := getDefaultRadar()
	inputX := 3
	inputY := 3
	expectedDepth := 7
	expectedFound := false

	outputDepth, outputFound := radar.IsLowSpot(inputX, inputY)

	assert.Equal(t, expectedDepth, outputDepth)
	assert.Equal(t, expectedFound, outputFound)
}

/// ---------------------------------------------------------------------------
///   Helper Functions
/// ---------------------------------------------------------------------------

func getDefaultRadar() *Radar {
	return createRadar([]string{
		"2199943210",
		"3987894921",
		"9856789892",
		"8767896789",
		"9899965678",
	},
	)
}
