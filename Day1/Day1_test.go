package day1

import "testing"

func TestCalcFuel( t *testing.T) {
	
	testData := make(map[int]int)

	testData[12] = 2
	testData[14] = 2
	testData[1969] = 654
	testData[100756] = 33583
	testData[1] = -2
	testData[3] = -1
	testData[6] = 0
	testData[8] = 0
	testData[9] = 1
	testData[11] = 1
	testData[0] = -2
	testData[-3] = -3

	for test, expected := range testData {
		result := calcFuel(test)
		if result != expected {
			t.Errorf("Failed on: %v. Got: %v. Expected: %v", test, result, expected)
		}
	}
}

func TestCalcFuelRecursively( t *testing.T) {
	
	testData := make(map[int]int)

	testData[14] = 2
	testData[1969] = 966
	testData[100756] = 50346

	for test, expected := range testData {
		result := calcFuelRecursively(test)
		if result != expected {
			t.Errorf("Failed on: %v. Got: %v. Expected: %v", test, result, expected)
		}
	}
}
