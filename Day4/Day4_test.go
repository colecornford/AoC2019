package day4

import "testing"

func TestWithinRange(t *testing.T) {
	
	beginRange := 400000
	endRange := 500000

	passRange := []int{400000, 499999, 444444}
	for _, num := range passRange {
		if !isInRange(num, beginRange, endRange) {
			t.Errorf("isInRange failed on positive test case: %v", num)
		}
	}

	failRange := []int{399999, 500000}
	for _, num := range failRange {
		if isInRange(num, beginRange, endRange) {
			t.Errorf("isInRange failed on negative test case: %v", num)
		}
	}
}

func TestLength6Digit(t *testing.T) {
	
	passLength := []int{111111, 999999}
	for _, num := range passLength {
		if !is6Digits(num) {
			t.Errorf("is6Digits failed on positive test case: %v", num)
		}
	}

	failLength := []int{1, 22, 333, 4444, 55555, 7777777, 88888888}
	for _, num := range failLength {
		if is6Digits(num) {
			t.Errorf("is6Digits failed on negative test case: %v", num)
		}
	}
}

func TestTwoAdjacent(t *testing.T) {
	
	passAdjacent := []int{112345, 123345, 123400, 112233}
	for _, num := range passAdjacent {
		if !hasTwoAdjacentNumbers(num) {
			t.Errorf("hasTwoAdjacentNumbers failed on positive test case: %v", num)
		}
	}

	failAdjacent := []int{121212, 123456, 023130}
	for _, num := range failAdjacent {
		if hasTwoAdjacentNumbers(num) {
			t.Errorf("hasTwoAdjacentNumbers failed on negative test case: %v", num)
		}
	}
}

func TestDigitsNeverDecrease(t *testing.T) {
	
	passDigits := []int{123456, 112345, 000001, 999999, 112233}
	for _, num := range passDigits {
		if !digitsNeverDecrease(num) {
			t.Errorf("digitsNeverDecrease failed on positive test: %v", num)
		}
	}

	failDigits := []int{123450, 987654, 100000, 555550}
	for _, num := range failDigits {
		if digitsNeverDecrease(num) {
			t.Errorf("digitsNeverDecrease failed on negative test case: %v", num)
		}
	}

}

func TestNotInLargerGroup(t *testing.T) {
	
	passNotInLarger := []int{112233, 111122, 111223, 144555}
	for _, num := range passNotInLarger {
		if !notInLargerGroup(num) {
			t.Errorf("notInLargerGroup failed on positive test case: %v", num)
		}
	}

	failNotInLarger := []int{123444, 111111, 111222}
	for _, num := range failNotInLarger {
		if notInLargerGroup(num) {
			t.Errorf("notInLargerGroup failed on negative test case: %v", num)
		}
	}
}