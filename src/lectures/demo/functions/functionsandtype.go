package main

import (
	"fmt"
)

type celsius float64
type kelvin float64

//Kelvin to Celsius converts K to C
func kelvinToCelsius(k kelvin) celsius {
	return celsius(k - 273.15)
}

func celsiusToKelvin(c celsius) kelvin {
	return kelvin(c + 273.15 )
}

func main() {
var k kelvin = 294.0
var b celsius = 127.0
c := kelvinToCelsius(k)
d := celsiusToKelvin(b)
fmt.Print(k, "K is ", c, "C")
fmt.Print(b, "C is ", d, "C")
}