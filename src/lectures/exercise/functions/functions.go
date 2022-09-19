//--Summary:
//  Use functions to perform basic operations and print some
//  information to the terminal.
//
//--Requirements:
//* Write a function that accepts a person's name as a function
//  parameter and displays a greeting to that person.
//* Write a function that returns any message, and call it from within
//  fmt.Println()
//* Write a function to add 3 numbers together, supplied as arguments, and
//  return the answer
//* Write a function that returns any number
//* Write a function that returns any two numbers
//* Add three numbers together using any combination of the existing functions.
//  * Print the result
//* Call every function at least once

package main

import (
	"fmt"
	"math/rand"
)

var name string
var anymessage string = "I am any message"
var anynumber int

func acceptName() {
	fmt.Println("Please Enter your name: ")
	fmt.Scanf("%s\n", &name)
	fmt.Println("Hello", name)
}
func printAnymessage() {
	fmt.Println(anymessage)
}

func anyNumberfunc() int {
	anynumber = rand.Int()
	fmt.Println("AnyNumber is", anynumber)
	return anynumber
}

func twonumbers() (int, int) {
	return 2, 5
}

func addThree(a, b, c int) int {
	return a + b + c
}

func sum(lhs, rhs int) int {
	return lhs + rhs
}

func numberAddition() int {
	var number1, number2, number3 int
	fmt.Println("Please enter your first number: ")
	fmt.Scanf("%d\n", &number1) // \n adds a line break so everything in the scanf is stored on its own line
	fmt.Println("Please enter your second number: ")
	fmt.Scanf("%d\n", &number2)
	fmt.Println("Please enter your third and final number: ")
	fmt.Scanf("%d\n", &number3)
	var result int = number1 + number2 + number3
	fmt.Println("The addition of all those numbers are", result)
	return result
}

func main() {
	acceptName()
	numberAddition()
	printAnymessage()
	twonumbers()
	a, b := twonumbers()
	anwser := addThree(anyNumberfunc(), a, b)
	resulttwo := sum(2, 2)
	fmt.Println(resulttwo)
	fmt.Println(anwser)

}
