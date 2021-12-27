package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

/// ---------------------------------------------------------------------------
///   createCaveSystem
/// ---------------------------------------------------------------------------

func Test_createCaveSystem(t *testing.T) {
	input := []string{
		"start-A",
		"start-b",
		"A-c",
		"A-b",
		"b-d",
		"A-end",
		"b-end",
	}

	expected := &CaveSystem{
		segments: map[string][]string{
			"start": {"A", "b"},
			"A":     {"c", "b", "end"},
			"b":     {"A", "d", "end"},
			"c":     {"A"},
			"d":     {"b"},
		},
	}

	output := createCaveSystem(input)

	assert.Equal(t, expected, output)
}
