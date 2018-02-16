package main

import (
"fmt"
)

type (
	Celsius    float64
	Fahrenheit float64
	Kelvin 	   float64
)

func main() {
	var todayF Fahrenheit = 48
	var todayC Celsius
	var todayK Kelvin
	fmt.Printf("%6.2f\n", CtoF(0))   // should be 32
	fmt.Printf("%6.2f\n", CtoF(100)) // should be 212
	todayC = FtoC(todayF)
	fmt.Printf("%6.2f\n", todayC)

	fmt.Printf("%6.2f\n", FtoK(32))
	fmt.Printf("%6.2f\n", FtoK(212))
	todayK = FtoK(todayF)
	fmt.Printf("%6.2f\n", todayK)

	fmt.Printf("%6.2f\n", CtoK(0))
	fmt.Printf("%6.2f\n", CtoK(100))
	todayK = CtoK(todayC)
	fmt.Printf("%6.2f\n", todayK)
}

// CelsiusToFahrenheit converts celsius to fahrenheit.
func CtoF(c Celsius) Fahrenheit {
	return Fahrenheit(float64((c * 9 / 5) + 32))
}

// FahrenheitToCelsius converts fahrenheit to celsius.
func FtoC(f Fahrenheit) Celsius {
	return Celsius(float64((f - 32) * 5 / 9))
}

func FtoK(f Fahrenheit) Kelvin {
	return CtoK(FtoC(f))
}

func CtoK(c Celsius) Kelvin {
	return Kelvin(c + 273.15)
}