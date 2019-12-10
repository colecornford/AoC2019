package main

import (
	"io/ioutil"
	"strings"
	"fmt"
	"strconv"
	"math"
)

func main() {
	//part1()
	part2()
}

func getData()([]bearing, []bearing) {
	content, _ := ioutil.ReadFile("test.txt")
	wires := strings.Split(string(content), "\n")
	wire1Bearings := getBearings(strings.Split(string(wires[0]), ","))
	//fmt.Printf(wire1Bearings)
	wire2Bearings := getBearings(strings.Split(string(wires[1]), ","))
	//fmt.Printf(wire1Bearings)
	return wire1Bearings, wire2Bearings
}

type bearing struct {
	direction string
	distance int
}

type coordinate struct {
	x int
	y int
}

func part1() {

	wire1Bearings, wire2Bearings := getData()

	wire1Coordinates := drawPath(wire1Bearings)
	wire2Coordinates := drawPath(wire2Bearings)
	fmt.Println("Done drawingPaths")
	intersections := findIntersections(wire1Coordinates, wire2Coordinates)
	fmt.Println("Done findingIntersections")
	result := calcLowestManhattan(intersections)
	fmt.Printf("The lowest Manhattan is: %v", result)
}

func getBearings(strInput []string)([]bearing) {
	var bearings []bearing
	for _, item := range strInput {
		var newBearing bearing
		newBearing.distance, _ = strconv.Atoi(strings.Trim(item, "UDLR")) // Remove UpDownLeftRight
		newBearing.direction = strings.Trim(item, "0123456789")  // Remove Digits
		bearings = append(bearings, newBearing)
	}
	return bearings
}

func drawPath(bearings []bearing)([]coordinate) {
	currentX := 0
	currentY := 0
	var coordinates []coordinate
	for _, item := range bearings {
		for x := 1; x <= item.distance; x++ {
			if item.direction == "L"{
				if x == item.distance {
					currentX = currentX - item.distance
					continue
				}
				coord := coordinate{x: currentX - x, y: currentY}
				coordinates = append(coordinates, coord)
			}
			if item.direction == "R"{
				if x == item.distance {
					currentX = currentX + item.distance
					continue
				}
				coord := coordinate{x: currentX + x, y: currentY}
				coordinates = append(coordinates, coord)
			}
		}
		for y := 1; y <= item.distance; y++ {
			if item.direction == "U"{
				if y == item.distance {
					currentY = currentY - item.distance
					continue
				}
				coord := coordinate{x: currentX, y: currentY - y}
				coordinates = append(coordinates, coord)
				
			}
			if item.direction == "D"{
				if y == item.distance {
					currentY = currentY + item.distance
					continue
				}
				coord := coordinate{x: currentX, y: currentY + y}
				coordinates = append(coordinates, coord)
			}
		}
		coord := coordinate{x: currentX, y: currentY}
		coordinates = append(coordinates, coord)
	}
	return coordinates
}

func findIntersections(wire1 []coordinate, wire2 []coordinate)([]coordinate) {
	var intersection []coordinate
	for _, wire1Coordinate := range wire1 {
		for _, wire2Coordinate := range wire2 {
			if wire1Coordinate.x == wire2Coordinate.x && wire1Coordinate.y == wire2Coordinate.y {
				intersection = append(intersection, wire2Coordinate)	
			}
		}
	}
	return intersection
}

func calcLowestManhattan(intersections []coordinate)(int){
	max := 1000000000.0
	for _, coord := range intersections {
		newMax := math.Abs(float64(coord.x)) + math.Abs(float64(coord.y))
		if newMax < max {
			max = newMax
		}
	}
	return int(max)
}

func calculateSteps(wire1 []coordinate, wire2 []coordinate, intersection coordinate)(int) {
	steps := 0
	for index1, wire1Coordinate := range wire1 {
		if wire1Coordinate.x == intersection.x && wire1Coordinate.y == intersection.y {
			steps = steps + index1
			break
		}
	}
	for index2, wire2Coordinate := range wire2 {
		if wire2Coordinate.x == intersection.x && wire2Coordinate.y == intersection.y {
			steps = steps + index2
			break
		}
	}
	return steps
}

func part2() {

	wire1Bearings, wire2Bearings := getData()

	wire1Coordinates := drawPath(wire1Bearings)
	wire2Coordinates := drawPath(wire2Bearings)
	fmt.Println("Done drawingPaths")
	intersections := findIntersections(wire1Coordinates, wire2Coordinates)
	fmt.Println("Done findingIntersections")
	lowest := 1000000
	for _, intersection := range intersections {
		result := calculateSteps(wire1Coordinates, wire2Coordinates, intersection)
		if result < lowest {
			lowest = result
		}
	}
	lowest = lowest + 2 // account for 1 off bias on X and Y axis
	fmt.Printf("The lowest steps is: %v", lowest)
	
}
