package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const TEST_INPUT_FILE string = "test_data/sample_input.txt"

/// ---------------------------------------------------------------------------
///   bitData - getEpsilonValue
/// ---------------------------------------------------------------------------

func Test_bitData_getEpsilonValue_zero(t *testing.T) {
	input := newBitDataWithValues(1, 0)
	expected := "1"

	gammaValue := input.getEpsilonValue()

	assert.Equal(t, expected, gammaValue, "should be at the same")
}

func Test_bitData_getEpsilonValue_one(t *testing.T) {
	input := newBitDataWithValues(0, 1)
	expected := "0"

	gammaValue := input.getEpsilonValue()

	assert.Equal(t, expected, gammaValue, "should be at the same")
}

/// ---------------------------------------------------------------------------
///   bitData - getGammaValue
/// ---------------------------------------------------------------------------

func Test_bitData_getGammaValue_zero(t *testing.T) {
	input := newBitDataWithValues(1, 0)
	expected := "0"

	gammaValue := input.getGammaValue()

	assert.Equal(t, expected, gammaValue, "should be at the same")
}

func Test_bitData_getGammaValue_one(t *testing.T) {
	input := newBitDataWithValues(0, 1)
	expected := "1"

	gammaValue := input.getGammaValue()

	assert.Equal(t, expected, gammaValue, "should be at the same")
}

/// ---------------------------------------------------------------------------
///   bitData - processBit
/// ---------------------------------------------------------------------------

func Test_bitData_processBit_zero(t *testing.T) {
	input := '0'
	bitData := newBitData()
	expected := newBitDataWithValues(1, 0)

	bitData.processBit(input)

	assert.Equal(t, expected, bitData, "should be at the same")
}

func Test_bitData_processBit_one(t *testing.T) {
	input := '1'
	bitData := newBitData()
	expected := newBitDataWithValues(0, 1)

	bitData.processBit(input)

	assert.Equal(t, expected, bitData, "should be at the same")
}

func Test_bitData_processBit_invalidInput(t *testing.T) {
	input := '2'
	bitData := newBitData()

	assert.Panics(t, assert.PanicTestFunc(func() {
		bitData.processBit(input)
	}))
}

/// ---------------------------------------------------------------------------
///   dataRate - calculateEpsilonRate
/// ---------------------------------------------------------------------------

func Test_bitData_calculateEpsilonRate(t *testing.T) {
	input := []string{
		"00100",
		"11110",
		"10110",
		"10111",
		"10101",
		"01111",
		"00111",
		"11100",
		"10000",
		"11001",
		"00010",
		"01010",
		}
	data := processAllData(input)
	expected := int64(9)

	output := data.calculateEpsilonRate()

	assert.Equal(t, expected, output, "should be the same")
}

/// ---------------------------------------------------------------------------
///   dataRate - calculateGammaRate
/// ---------------------------------------------------------------------------

func Test_bitData_calculateGammaRate(t *testing.T) {
	input := []string{
		"00100",
		"11110",
		"10110",
		"10111",
		"10101",
		"01111",
		"00111",
		"11100",
		"10000",
		"11001",
		"00010",
		"01010",
		}
	data := processAllData(input)
	expected := int64(22)

	output := data.calculateGammaRate()

	assert.Equal(t, expected, output, "should be the same")
}

/// ---------------------------------------------------------------------------
///   dataRate - calculateCo2Rate
/// ---------------------------------------------------------------------------

func Test_bitData_calculateCo2Rate(t *testing.T) {
	input := []string{
		"00100",
		"11110",
		"10110",
		"10111",
		"10101",
		"01111",
		"00111",
		"11100",
		"10000",
		"11001",
		"00010",
		"01010",
		}
	data := processAllData(input)
	expected := int64(10)

	output := data.calculateCo2Rate()

	assert.Equal(t, expected, output, "should be the same")
}

/// ---------------------------------------------------------------------------
///   dataRate - calculateOxygenRate
/// ---------------------------------------------------------------------------

func Test_bitData_calculateOxygenRate(t *testing.T) {
	input := []string{
		"00100",
		"11110",
		"10110",
		"10111",
		"10101",
		"01111",
		"00111",
		"11100",
		"10000",
		"11001",
		"00010",
		"01010",
		}
	data := processAllData(input)
	expected := int64(23)

	output := data.calculateOxygenRate()

	assert.Equal(t, expected, output, "should be the same")
}

/// ---------------------------------------------------------------------------
///   dataRate - processValue
/// ---------------------------------------------------------------------------

func Test_bitData_processValue_zero(t *testing.T) {
	input := "00100"
	rateData := newRateData(5)
	expected := newRateData(5)
	expected.processedData[0].zero++
	expected.processedData[1].zero++
	expected.processedData[2].one++
	expected.processedData[3].zero++
	expected.processedData[4].zero++

	rateData.processValue(input)

	assert.Equal(t, expected, rateData, "should be at the same")
}

/// ---------------------------------------------------------------------------
///   filterData
/// ---------------------------------------------------------------------------

func Test_filterData_zeroInFirstPosition(t *testing.T) {
	input := []string{
		"00100",
		"11110",
		"10110",
	}
	expected := []string{
		"00100",
		}

	output := filterData(input, 0, '0')

	assert.Equal(t, expected, output, "should match")
}

func Test_filterData_oneInFirstPosition(t *testing.T) {
	input := []string{
		"00100",
		"11110",
		"10110",
	}
	expected := []string{
		"11110",
		"10110",
	}

	output := filterData(input, 0, '1')

	assert.Equal(t, expected, output, "should match")
}

/// ---------------------------------------------------------------------------
///   getInput
/// ---------------------------------------------------------------------------

func Test_getData_CanProperlyRead(t *testing.T) {
	expected := []string{
		"00100",
		"11110",
		"10110",
		"10111",
		"10101",
		"01111",
		"00111",
		"11100",
		"10000",
		"11001",
		"00010",
		"01010",
		}

	output := getInput(TEST_INPUT_FILE)

	assert.Equal(t, expected, output, "input file not parsed correctly")
}

/// ---------------------------------------------------------------------------
///   processAllData
/// ---------------------------------------------------------------------------

func Test_getData_processAllData(t *testing.T) {
	input := []string{
		"00100",
		"11110",
		"10110",
		"10111",
		"10101",
		"01111",
		"00111",
		"11100",
		"10000",
		"11001",
		"00010",
		"01010",
		}
	expected := newRateData(5)
	expected.rawData = input
	expected.processedData[0].zero = 5
	expected.processedData[0].one = 7
	expected.processedData[1].zero = 7
	expected.processedData[1].one = 5
	expected.processedData[2].zero = 4
	expected.processedData[2].one = 8
	expected.processedData[3].zero = 5
	expected.processedData[3].one = 7
	expected.processedData[4].zero = 7
	expected.processedData[4].one = 5
	
	output := processAllData(input)

	assert.Equal(t, expected, output, "they should match")
}