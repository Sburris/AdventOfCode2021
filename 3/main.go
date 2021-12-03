package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	inputFile := "input.txt"
	input := getInput(inputFile)

	data := processAllData(input)
	
	epsilon := data.calculateEpsilonRate()
	gamma := data.calculateGammaRate()	
	fmt.Printf("Epsilon: %d\n", epsilon)
	fmt.Printf("Gamma: %d\n", gamma)
	fmt.Printf("Power Consumption: %d\n", epsilon * gamma)

	oxygen := data.calculateOxygenRate()
	co2 := data.calculateCo2Rate()
	fmt.Printf("Oxygen Rate: %d\n", oxygen)
	fmt.Printf("Co2 Rate: %d\n", co2)
	fmt.Printf("Life Support Rate: %d\n", co2 * oxygen)
}

/// ---------------------------------------------------------------------------
///   Data Types rateData
/// ---------------------------------------------------------------------------

type rateData struct {
	rawData []string
	processedData []bitData
}

func newRateData(length int) *rateData {
	d := rateData{}
	d.processedData = make([]bitData, length)
	return &d
}

func (this *rateData) processValue(data string) {
	for index, value := range data {
		this.processedData[index].processBit(value)
	}
}

func (this *rateData) calculateEpsilonRate() int64 {
	rawData := ""
	for _, value := range this.processedData {
		rawData += value.getEpsilonValue()
	}

	epsilonValue, err := strconv.ParseInt(rawData, 2, 64)
	checkErr(err)

	return epsilonValue
}

func (this *rateData) calculateGammaRate() int64 {
	rawData := ""
	for _, value := range this.processedData {
		rawData += value.getGammaValue()
	}

	gammaValue, err := strconv.ParseInt(rawData, 2, 64)
	checkErr(err)

	return gammaValue
}

func (this *rateData) calculateOxygenRate() int64 {
	length := len(this.rawData[0])
	data := this.rawData

	for index:=0; index < length; index++ {
		processedData := processAllData(data)

		mostCommon := []rune(processedData.processedData[index].getMostCommon())
		data = filterData(data, index, mostCommon[0])
		
		if len(data) == 1 {
			value, err := strconv.ParseInt(data[0], 2, 64)
			checkErr(err)
			return value
		}
	}

	panic("could not calculate CO2 rate")
}

func (this *rateData) calculateCo2Rate() int64 {
	length := len(this.rawData[0])
	data := this.rawData

	for index:=0; index < length; index++ {
		processedData := processAllData(data)

		leastCommon := []rune(processedData.processedData[index].getLeastCommon())
		data = filterData(data, index, leastCommon[0])
		
		if len(data) == 1 {
			value, err := strconv.ParseInt(data[0], 2, 64)
			checkErr(err)
			return value
		}
	}

	panic("could not calculate oxygen rate")
}

/// ---------------------------------------------------------------------------
///   Data Types: bitData
/// ---------------------------------------------------------------------------


type bitData struct {
	zero int
	one int
}

func newBitData() *bitData {
	return &bitData{zero: 0, one: 0}
}

func newBitDataWithValues(zeros int, ones int) *bitData {
	return &bitData{zero: zeros, one: ones}
}

/// ---------------------------------------------------------------------------
///   
/// --------Data Types rateData-------------------------------------------------------------------

func filterData(data []string, index int, targetValue rune) []string {
	result := make([]string, 0)
	
	for _, value := range data {
		runes := []rune(value)
		if runes[index] == targetValue {
			result = append(result, value)
		}
	}

	return result
}

/// ---------------------------------------------------------------------------
///   General functions
/// ---------------------------------------------------------------------------

func (this *bitData) processBit(value rune) {
	switch value {
	case '0': this.zero++
	case '1': this.one++
	default: panic(fmt.Sprintf("unknown bit value: '%c'\n", value))
	}
}

func (this *bitData) getMostCommon() string {
	if this.zero > this.one {
		return "0"
	}
	return "1"
}

func (this *bitData) getLeastCommon() string {
	if this.zero > this.one {
		return "1"
	}
	return "0"
}

func (this *bitData) getGammaValue() string {
	if this.one > this.zero {
		return "1"
	} else if this.zero > this.one {
		return "0"
	}
	panic("values are the same, reboot universe")
}

func (this *bitData) getEpsilonValue() string {
	if this.one > this.zero {
		return "0"
	} else if this.zero > this.one {
		return "1"
	}
	panic("values are the same, reboot universe")
}

/// ---------------------------------------------------------------------------
///   Main Program functions
/// ---------------------------------------------------------------------------

// Take the raw input and put it into the proper data data structure
func processAllData(data []string) *rateData {
	if len(data) == 0 {
		return nil
	}
	dataValueLenght := len(data[0])
	results := newRateData(dataValueLenght)

	results.rawData = data

	for _, value := range data{
		results.processValue(value)
	}

	return results
}

// Read in raw data so we can process it
func getInput(filename string) []string {
	results := make([]string, 0)

	file, err := os.Open(filename)
	checkErr(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		results = append(results, scanner.Text())
	}

	return results
}

// Check if an error has occured
func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}