package main

import (
	"fmt"
	"math"
)

func main() {
	temperature := map[string]int{
		"Earth": 15,
		"Mars":  -65,
	}
	planets := map[string]string{
		"Earth": "Sector ZZ9",
		"Mars":  "Sector ZZ9",
	}
	temperatures := []float64{
		-28.0, 32.0, -31.0, -29.0, -23.0, -29.0, -28.0, -33.0,
	}
	temp := temperature["Earth"]
	fmt.Printf("On average the Earth is %vº C.\n", temp)
	temperature["Earth"] = 16
	temperature["Venus"] = 464
	fmt.Println(temperature)

	if moon, ok := temperature["Moon"]; ok {
		fmt.Printf("On average the moon is %vº C.\n", moon)
	} else {
		fmt.Println("Where is the moon?")
	}

	planetsMarkII := planets
	planets["Earth"] = "whoops"
	fmt.Println(planets)
	fmt.Println(planetsMarkII)
	delete(planets, "Earth")
	fmt.Println(planetsMarkII)

	groups := make(map[float64][]float64)
	for _, t := range temperatures {
		g := math.Trunc(t/10) * 10
		groups[g] = append(groups[g], t)
	}
	for g, temperatures := range groups {
		fmt.Printf("%v: %v\n", g, temperatures)
	}
}
