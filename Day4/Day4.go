package main

import (
	"strconv"
	"fmt"
	"strings"
)

func main() {
	part1()
	part2()
}

func part1() {
	numberPasswords := 0
	startRange := 271973
	endRange := 785961
	for i := startRange; i < endRange; i++ {
		if isInRange(i, startRange, endRange) && is6Digits(i) && hasTwoAdjacentNumbers(i) && digitsNeverDecrease(i) {
			numberPasswords += 1
		}
	}
	
	fmt.Printf("Part1: Total passwords: %v", numberPasswords)
}

func part2() {
	numberPasswords := 0
	startRange := 271973
	endRange := 785961
	for i := startRange; i < endRange; i++ {
		if isInRange(i, startRange, endRange) && is6Digits(i) && hasTwoAdjacentNumbers(i) && digitsNeverDecrease(i) && notInLargerGroup(i){
			numberPasswords += 1
		}
	}
	
	fmt.Printf("Part2: Total passwords: %v", numberPasswords)
}

func isInRange(number int, lower int, upper int)(bool){
	return number >= lower && number < upper 
}

func is6Digits(number int)(bool){
	return isInRange(number, 100000, 1000000)
}

func hasTwoAdjacentNumbers(number int)(bool){
	strNumber := strconv.Itoa(number)
	strFineSubstrings := []string{"00","11","22","33","44","55","66","77","88","99"}
	for _, input := range strFineSubstrings {
		if input != "" {
			if strings.Contains(strNumber, input){
				return true
			}
		}
	}
	return false
}

func digitsNeverDecrease(number int)(bool){
	maxNum := 0
	strDigit := 0
	strNumber := strconv.Itoa(number)
	for _, char := range strNumber {
		strDigit, _ = strconv.Atoi(string(char))
		if strDigit < maxNum {
			return false
		}
		maxNum = strDigit
	}
	return true
}

func notInLargerGroup(number int)(bool){
	strNumber := strconv.Itoa(number)
	strEvilSubstrings := []string{"000","111","222","333","444","555","666","777","888","999"}
	strFineSubstrings := []string{"00","11","22","33","44","55","66","77","88","99"}
	for index, input := range strEvilSubstrings {
		if strings.Contains(strNumber, input){
			strFineSubstrings[index] = ""
		}
	}
	for _, input := range strFineSubstrings {
		if input != "" {
			if strings.Contains(strNumber, input){
				return true
			}
		}
	}
	return false	
}