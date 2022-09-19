package main

import "fmt"

func main() {

	first := 100
	var second *int = &first //& this is the address operator of where the memory location lives.
	third := &first

	first++ // Go reads the memory location of the variable first, increments the value and restores that new value in the same memory location. Incrementing it by 1 each time.

	alpha := 100
	beta := &alpha

	fmt.Println(second == third)   // this is comparing the memory location of the pointer and seeing if one is equal to the other
	fmt.Println(second == beta)    // this is comparing the memory location of the pointer and seeing if one is equal to the other
	fmt.Println(*second == *third) // this is comparing the values of the pointers and seeing if they are equal NOT the memory location * does this.
	fmt.Println(*second == *beta)  // this is comparing the values of the pointers and seeing if they are equal NOT the memory location * does this.

	/*
		The value of second variable is the memory address of the first variable. Apparently this memory address restriction is important as it stops a variable from changing types.
	*/
	fmt.Println("First", first)
	fmt.Println("Second", second)
	newfirst := float64(first) // changing first variable to a float64.
	newfirst = 20.1
	fmt.Println("newfirst variable value", newfirst)
	fmt.Println("Second value, should be the original value of variable first", *second) /* this is accessing the orignal memory store from first++ and taking that value of the integer before I changed it to a float.
	   By putting * I can still access the old variable value. */
}
