package main

import "fmt"

func main() {

	var planets [8]string // This is a basic Array. 

	planets[0] = "Mercury"
	planets[1] = "Venus"
	planets[2] = "Earth"

	earth := planets[2]
	fmt.Println(len(planets))
	fmt.Println(earth)

	planetstwo := [...]string { //This is a Composite Array meaning each value is initialised at compile time all together instead of individulally.
	"Mercury",
	"Venus",
	"Earth",
	"Mars",
	"Jupiter",
	"Saturn",
	"Uranus",
	"Neptune",
	}
	fmt.Println(len(planetstwo))

	dwarfs := [5]string{"Ceres", "Pluto", "Haumea", "Makemake", "Eris"} // This is setting up the dwarfs array
	for i := 0; i < len(dwarfs); i++ { // This is checking the lengh of the array so that the for loop can loop through away and print all values of the dwarf array. 
		dwarf := dwarfs[i] //This makes a copy of the dwarfs array
		fmt.Println(i, dwarf) // This is asking the compiler to now print each value of the array from 0 to 4. As all arrays start at 0. 
	} // this is closing the for loop. 

    for i, item := range dwarfs { // This does the same thing as the above example except you can use range to work with Runes. 
    fmt.Println(i, item)
    }
    
	planetsthree := [...]string{ // This is about assigning an Array to a variable or passing it to a function what happens is it makes a complete copy of that Array.  
	"Mercury",                   // In this example we are assigning it to a variable not a function.
	"Venus",
	"Earth",
	"Mars",
	"Jupiter",
	"Saturn",
	"Uranus",
	"Neptune",
	}
	planetsMarkII := planetsthree // Here is the assign where a copy is made. 
	planetsthree[2] = "whoops"
	fmt.Println(planetsthree)
	fmt.Println(planetsMarkII)
}
