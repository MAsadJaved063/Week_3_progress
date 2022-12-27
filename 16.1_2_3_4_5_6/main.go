package main

import "fmt"

func terraform(planets [8]string) {
	for i := range planets {
		planets[i] = "New " + planets[i]
	}
}
func main() {
	var planets [8]string
	planets[0] = "Mercury"
	planets[1] = "Venus"
	planets[2] = "Earth"
	earth := planets[2]
	fmt.Println(earth)
	fmt.Println(len(planets))
	fmt.Println(planets[3] == "")

	dwarfs := [5]string{"Ceres", "Pluto", "Haumea", "Makemake", "Eris"}
	for i := 0; i < len(dwarfs); i++ {
		dwarf := dwarfs[i]
		fmt.Println(i, dwarf)
	}

	planets_1 := [...]string{
		"Mercury",
		"Venus",
		"Earth",
		"Mars",
		"Jupiter",
		"Saturn",
		"Uranus",
		"Neptune",
	}
	terraform(planets_1)
	fmt.Println(planets_1)
	//Arrays of arrays
	var board [8][8]string
	board[0][0] = "r"
	board[0][7] = "r"
	for column := range board[1] {
		board[1][column] = "p"
	}
	fmt.Print(board)
}
