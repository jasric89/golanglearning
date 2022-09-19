package main

import (
	"fmt"
	"strings"
)

func WelcomeMessage(customer string) string {
	return "Welcome to the Tech Palace, " + "" + customer
	panic("Please implement the WelcomeMessage() function")
}

func AddBorder(welcomeMsg string, numStarsPerLine int) string {
	numStarsPerLine = "" + strings.Repeat("*", 10)
	return numStarsPerLine
	panic("Please implement the AddBorder() function")
}

func main() {
	const name string = "Judy"
	fmt.Println(WelcomeMessage(name))
	fmt.Println("Welcome" + strings.Repeat("*", 10))
	//AddBorder("Welcome!", 10)
}
