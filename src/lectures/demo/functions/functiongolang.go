package main

import (
	"fmt"
)
//The kelvinToCelsius function in listing 12.1 is isolated from other functions. Its only input is the parameter it accepts, and its only output is the result it returns. 
//It makes no modifications to external state. Such functions are side-effect-free and are the the easiest to understand, test, and reuse.
//The kelvinToCelsius function does modify the variable k, but k and kelvin are completely independent variables, so assigning a new value to k inside the function has no impact 
//on the kelvin variable in main. This behavior is called pass by value, because the k parameter is initialized with the value of the kelvin argument. 
//Pass by value facilitates the boundary between functions, helping to isolate one function from another.
//We’ve given the variables different names, but pass by value applies even if arguments and parameters have the same names.
//Additionally, the variable named k in kelvinToCelsius is completely independent from any variable named k in other functions, thanks to variable scope. 
//Scope is covered in lesson 4, but to reiterate, the parameters in a function declaration and the variables declared within a function body have function scope. 
//Variables declared in different functions are completely independent, even if they have the same name.

//kelvin to celsius converts k to c 

func kelvinToCelsius(float64) float64 { // This function returns one value of type64 the result is delivered back to the user with the return keyword. 
	k := 273.15
	return k 
} 

func main(){
   
	kelvin := 294.0
	celsius := kelvinToCelsius(kelvin)
	fmt.Println(kelvin, "k is ", celsius, "c")


}