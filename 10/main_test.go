package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

/// ---------------------------------------------------------------------------
///   processline
/// ---------------------------------------------------------------------------

func Test_processLine_one(t *testing.T) {
	input := "{([(<{}[<>[]}>{[]{[(<()>"
	expected := 1197

	output, _ := processLine(input)

	assert.Equal(t, expected, output)
}

func Test_processLine_two(t *testing.T) {
	input := "<{([{{}}[<[[[<>{}]]]>[]]"
	expected := 294

	_, output := processLine(input)

	assert.Equal(t, expected, output)
}

/// ---------------------------------------------------------------------------
///   Stack - Create
/// ---------------------------------------------------------------------------

func Test_Stack_createStack(t *testing.T) {
	expected := &Stack{
		buffer: []rune{},
	}

	output := createStack()

	assert.Equal(t, expected, output)
}

/// ---------------------------------------------------------------------------
///   Stack - head
/// ---------------------------------------------------------------------------

func Test_Stack_head(t *testing.T) {
	stack := createStack()
	expected := 'c'

	stack.push('a')
	stack.push('b')
	stack.push('c')

	output, _ := stack.head()

	assert.Equal(t, expected, output)
}

/// ---------------------------------------------------------------------------
///   Stack - push
/// ---------------------------------------------------------------------------

func Test_Stack_pop(t *testing.T) {
	stack := createStack()
	expected := &Stack{
		buffer: []rune{'a', 'b'},
	}

	stack.push('a')
	stack.push('b')
	stack.push('c')

	stack.pop()

	assert.Equal(t, expected, stack)
}

/// ---------------------------------------------------------------------------
///   Stack - push
/// ---------------------------------------------------------------------------

func Test_Stack_push(t *testing.T) {
	stack := createStack()
	expected := &Stack{
		buffer: []rune{'a', 'b', 'c'},
	}

	stack.push('a')
	stack.push('b')
	stack.push('c')

	assert.Equal(t, expected, stack)
}
