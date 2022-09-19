package main

import (
	"fmt"
)

func numberAddition() int {
	var number1, number2, number3 int
	fmt.Println("Please enter your first number: ")
	fmt.Scanf("%d\n", &number1)
	fmt.Println("Please enter your second number: ")
	fmt.Scanf("%d\n", &number2)
	fmt.Println("Please enter your third and final number: ")
	fmt.Scanf("%d\n", &number3)
	var result int = number1 + number2 + number3
	fmt.Println("The addition of all those numbers are", result)
	return result
}

func main() {
	numberAddition()
}
