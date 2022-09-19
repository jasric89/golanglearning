package main

import (
	"fmt"
     "time"
)

func main() {
   var count = 10 // define a variable and assign a value in this instance our variable is count and we have made it an interger with a value of 10. 
   for count > 0 { // count is more than 0 run the loop. 
	fmt.Println(count)
	time.Sleep(time.Second) // pauses the program for a second to make the countdown look realistic.(Kinda cool)
	count-- //Reduces count to 0 
   } // Everything before this parenthesis is in the for loop. 
   fmt.Println("Liftoff!!") // This is outside of the for loop and runs when the loop is finished so when Count is = to 0 
}