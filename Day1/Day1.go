package main

import (
	"io/ioutil"
	"strings"
	"strconv"
)

func main() {
	part1()
	part2()
}

func getData()(lines []string){
	content, err := ioutil.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	return strings.Split(string(content), "\n")
}

func part1() {
	data := getData()
	fuelCount := 0
	for _, sMass := range data {
		iMass, err := strconv.Atoi(sMass)
		if err != nil {
			print(err)
		}
		fuelCount = fuelCount + calcFuel(int(iMass))
	}
	print("Part 1 Answer is: ")
	print(fuelCount)
	print("\n")
}

func calcFuel(mass int)(int){
	return (mass / 3) - 2
}

func part2() {
	data := getData()
	fuelCount := 0
	for _, sMass := range data {
		iMass, err := strconv.Atoi(sMass)
		if err != nil {
			print(err)
		}
		fuelCount = fuelCount + calcFuelRecursively(int(iMass))
	}
	print("Part 2 Answer is: ")
	print(fuelCount)
	print("\n")
}

func calcFuelRecursively(totalFuel int)(int){
	nextFuel := calcFuel(totalFuel)
	if nextFuel <= 0 {
		return 0
	} else {
		return calcFuel(totalFuel) + calcFuelRecursively(nextFuel)
	}
}