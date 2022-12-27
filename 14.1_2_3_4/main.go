package main

import (
	"fmt"
	"math/rand"
	"time"
)

type kelvin float64

func fakeSensor() kelvin {
	return kelvin(rand.Intn(151) + 150)
}
func realSensor() kelvin {
	return 0
}

type kelvin_2 float64

func measureTemperature(samples int, sensor func() kelvin) {
	for i := 0; i < samples; i++ {
		k := sensor()
		fmt.Printf("%vº K\n", k)
		time.Sleep(time.Second)
	}
}
func fakeSensor_2() kelvin {
	return kelvin(rand.Intn(151) + 150)
}

func main() {
	sensor := fakeSensor
	fmt.Println(sensor())
	sensor = realSensor
	fmt.Println(sensor())
	measureTemperature(3, fakeSensor_2)

}
