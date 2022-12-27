package main

import (
	"fmt"
	"sort"
	"strings"
)

func hyperspace(worlds []string) {
	for i := range worlds {
		worlds[i] = strings.TrimSpace(worlds[i])
	}
}
func main() {
	planets := []string{
		"Mercury",
		"Venus",
		"Earth",
		"Mars",
		"Jupiter",
		"Saturn",
		"Uranus",
		"Neptune",
	}
	terrestrial := planets[0:4]
	gasGiants := planets[4:6]
	iceGiants := planets[6:8]
	fmt.Println(terrestrial, gasGiants, iceGiants)

	fmt.Println(gasGiants[0]) //Prints Jupiter only

	// hyperspace removes the space surrounding worlds

	planets_1 := []string{" Venus ", "Earth ", " Mars"}
	hyperspace(planets_1)
	fmt.Println(strings.Join(planets_1, " "))

	sort.StringSlice(planets).Sort()
	fmt.Println(planets)
}
