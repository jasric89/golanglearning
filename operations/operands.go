package main

import "fmt"

/*

|| This operator returns true if the first operand is met the second operand will not be evlautated.

&& This operator returns true if both values are true. If the first value is false then the second value will not be evaluated.

! This operator is used with a single operand. It returns true if the operand is false and false if the operand is true.

Only bool values can be used with logical operators, and Go will not attempt to convert a value to get a true or false value. If its an expression then the code is evaluated to get the bool result true or false.

*/

func main() {

	maxMph := 50
	passengerCapacity := 1
	airbags := true // This is a value that go will not convert to get a true or false value. Once set that value is the value.

	// Bellow are expressions.
	familycar := passengerCapacity > 2 && airbags
	sportsCar := maxMph > 100 || passengerCapacity == 2
	canCategorize := !familycar && !sportsCar

	fmt.Println(familycar)
	fmt.Println(sportsCar)
	fmt.Println(canCategorize)

}
