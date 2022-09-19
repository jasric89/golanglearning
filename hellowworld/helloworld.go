package main

import (
	"fmt"
)

// For future ref as I will prob come back to this issue again: https://www.geeksforgeeks.org/named-return-parameters-in-golang/
func printHelloWorld(h string, w string) (combinedstring string) { //Because I am using a named return value I have to parenthisis the paremeters.
	combinedstring = h + "," + w
	return combinedstring
}

func main() {
	fmt.Println(printHelloWorld("hello", "world"))
}
