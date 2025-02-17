package main

import (
	"fmt"
	"laboratory/internal/patterns"
)

func main() {

	var weatherstation patterns.Weatherstation

	var t tmp

	display := patterns.NewTemperatureDisplay(&weatherstation, t)
	weatherstation.SetTemperature(40)
}

type tmp struct{}

func (t tmp) Update(temperature float64) {
	fmt.Println("TemperatureDisplay: I need to update my display...")
}
