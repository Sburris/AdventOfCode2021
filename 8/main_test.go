package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

/// ---------------------------------------------------------------------------
///   countSegments
/// ---------------------------------------------------------------------------

// func Test_countSegments(t *testing.T) {
// 	input := []string{
// 		"abcdefg",
// 		"bcdef",
// 		"acdfg",
// 		"abcdf",
// 		"abd",
// 		"abcdef",
// 		"bcdefg",
// 		"abef",
// 		"abcdeg",
// 		"ab",
// 	}
// 	expected := map[int][]rune{
// 		4: {'g'},
// 		6: {'e'},
// 		7: {'f', 'c'},
// 		8: {'a', 'd'},
// 		9: {'b'},
// 	}

// 	output := countSegments(input)

// 	assert.Equal(t, expected, output)
// }

/// ---------------------------------------------------------------------------
///   createPatternLookup
/// ---------------------------------------------------------------------------

func Test_createPatternLookup(t *testing.T) {
	input := []string{
		"abcdefg",
		"bcdef",
		"acdfg",
		"abcdf",
		"abd",
		"abcdef",
		"bcdefg",
		"abef",
		"abcdeg",
		"ab",
	}
	expected := map[string]rune{
		"abcdeg":  '0',
		"ab":      '1',
		"acdfg":   '2',
		"abcdf":   '3',
		"abef":    '4',
		"bcdef":   '5',
		"bcdefg":  '6',
		"abd":     '7',
		"abcdefg": '8',
		"abcdef":  '9',
	}

	output := createPatternLookup(input)

	assert.Equal(t, expected, output)
}

/// ---------------------------------------------------------------------------
///   getNumber
/// ---------------------------------------------------------------------------

func Test_getNumber(t *testing.T) {
	input := "acedgfb cdfbe gcdfa fbcad dab cefabd cdfgeb eafb cagedb ab | cdfeb fcadb cdfeb cdbaf"
	expected := 5353

	output := getNumber(input)

	assert.Equal(t, expected, output)
}

/// ---------------------------------------------------------------------------
///   sortString
/// ---------------------------------------------------------------------------

func Test_sortString(t *testing.T) {
	input := "dca"
	expected := "acd"
	output := sortString(input)

	assert.Equal(t, expected, output)
}

/// ---------------------------------------------------------------------------
///   processInput
/// ---------------------------------------------------------------------------

func Test_processInput(t *testing.T) {
	input := "acedgfb cdfbe gcdfa fbcad dab cefabd cdfgeb eafb cagedb ab | cdfeb fcadb cdfeb cdbaf"
	expectedPatterns := []string{
		"abcdefg",
		"bcdef",
		"acdfg",
		"abcdf",
		"abd",
		"abcdef",
		"bcdefg",
		"abef",
		"abcdeg",
		"ab",
	}
	expectedDigits := []string{
		"bcdef",
		"abcdf",
		"bcdef",
		"abcdf",
	}

	outputPatterns, outputDigits := processInput(input)

	assert.Equal(t, expectedPatterns, outputPatterns)
	assert.Equal(t, expectedDigits, outputDigits)
}
