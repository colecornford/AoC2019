package main

import (
	"io/ioutil"
	"strings"
	"strconv"
	"fmt"
	"os"
)

func main() {
	part1()
	part2()
}

func getData()(lines []int) {
	content, err := ioutil.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}

	strSlice := strings.Split(string(content), ",")
	var intSlice []int
	for _, y := range strSlice {
		num, err := strconv.Atoi(y)
		if err != nil {
			panic(err)
		}
		intSlice = append(intSlice, num)
	}
	return intSlice
}

func part1() {
	input := getData()
	fmt.Println(restoreState(input))
}

func restoreState(intcode []int)([]int) {
	for index, element := range intcode {
		if index % 4 == 0 && index + 3 < len(intcode) { // Added so we don't encounter OOB issues.
			storelocation := intcode[index + 3]
			num1location := intcode[index + 2]
			num2location := intcode[index + 1]
			switch element {
			case 1:
				intcode[storelocation] = intcode[num1location] + intcode[num2location]
			case 2:
				intcode[storelocation] = intcode[num1location] * intcode[num2location]
			case 99:
				return intcode
			default:
				fmt.Println("Invalid intcodes programmed, exitting")
				os.Exit(2)
			}
		}
	}
	fmt.Println("Fully processed intcode, returning now.")
	return intcode
}

func part2() {
	for i := 0; i < 100; i++ {
		for j := 0; j < 100; j++ {
			input := getData()
			input[1] = i
			input[2] = j
			newState := restoreState(input)
			if newState[0] == 19690720 {
				fmt.Println(newState)
			}
		}
	}
}