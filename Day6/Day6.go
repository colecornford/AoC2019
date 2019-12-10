package main

import (
	"fmt"
	"strings"
	"io/ioutil"
)

func main() {
	part1()
	part2()
}

type space map[string]planet

type planet struct {
	parent string
	count int
}

func getData()(space) {
	content, _ := ioutil.ReadFile("input.txt")
	space := space{}
	for _, item := range strings.Split(string(content), "\n") {
		newPlanet := strings.Split(item, ")")
		parent := newPlanet[0]
		child := newPlanet[1]
		// hashmap, child is the key and planet (which contains parent) is the value.
		space[child] = planet{parent, 0}
	}
	return space
}

func part1() {
	space := getData()
	fmt.Printf("Part1 Answer: %v", countOrbits(space))
}

func part2() {
	space := getData()
	fmt.Printf("Part2 Answer: %v", YOUtoSAN(space))
}

func countOrbits(s space) (checksum int) {
	// For every item in the space hashmap, run getCount recursively and add the result to the checksum
	for name := range s {
		checksum = checksum + getCount(s, name)
	}
	return checksum
}

func getCount(s space, child string) int {
	planet, found := s[child]
	if !found {
		return 0 // error if no parent of child.
	}
	if planet.count == 0 {
		planet.count = 1 + getCount(s, planet.parent) // use the parent as next child
		s[child] = planet // update the count of this node once recursion is done.
	}

	return planet.count
}

func YOUtoSAN(s space) (checksum int) {
	// Start at the end of the route from YOU/SAN to COM.
	// Keep going backwards until the paths don't match.
	// When that happens, you've found the distance.

	YOU := getRoute(s, "YOU")
	SAN := getRoute(s, "SAN")

	fmt.Printf("YOU: %v\n SAN: %v", YOU, SAN)
	incrementYOU := len(YOU) - 1 
	incrementSAN := len(SAN) - 1 // Magic one because can't count reaching santa, only orbitting him.
	for (0 <= incrementYOU) && 0 <= incrementSAN { // Shutup, this is a While loop.
		if YOU[incrementYOU] != SAN[incrementSAN] {
			return incrementYOU + incrementSAN + 2 //  
		}
		incrementYOU = incrementYOU - 1
		incrementSAN = incrementSAN - 1
	}
	return 0
}

func getRoute(s space, from string) []string {
	route := []string{}
	planet, ok := s[from]
	for ok {
		route = append(route, planet.parent)
		planet, ok = s[planet.parent] // loop until you can't find the item in the hashmap.
	}
	return route
}