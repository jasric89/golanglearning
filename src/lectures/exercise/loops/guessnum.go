package main 

import (
	"fmt"
	"math/rand"
)

func main(){
    // I have a number and its less than a 100 the loop will exicute, If I put the for num > 100 this will mean that the compiler will say
	// My number is lower than a 100 im going to ignore the loop and go to the end, because the condition is not met. 
	// num++ will increase the value of num by 1 and num-- will decrease the value of num by 1. 
	var num = rand.Intn(1)+1//This makes the random genertor go through from number 1 up to 100, or the value thats in the for loop. 
	var mynum = 20
	for num < 100 { // When num equals a 100 exit. 
		fmt.Println("My Number is", num)
		if num == mynum {
		fmt.Println("you guessed my number!!")	
			break
		}
		num++ // This increases the Num variable 
	}
    fmt.Println("End of number game")
}