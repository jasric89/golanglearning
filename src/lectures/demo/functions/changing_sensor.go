package main
// annonymous function or AKA literal functions 
//func main() {
// func() { // This is an Annonymous Function 
//  fmt.Println("Functions Annonymous") 
// }() // Calls Function Annonymous 
//}
//
//
//
//
import (
	"fmt"
	"math/rand"
	"time"
)

type kelvin float64 // The Type declared. 
type sensor func() kelvin // This tells go that the variable sensor has a type to it which is a function 

func measureTemperature(samples int, sensor func() kelvin) { // Measure temp accepts a function as a the second parameter. 
	for i := 0; i < samples; i++ {
	k := sensor()
	fmt.Printf("%v K\n", k)
	time.Sleep(time.Second)
 }
}

func fakeSensor() kelvin {
	return kelvin(rand.Intn(151)+150)
}

func realSensor() kelvin {
	return 0
}

func calibrate(s sensor, offset kelvin) sensor {
	return func() kelvin {
		return s() + offset 
	}
}

func main(){
	measureTemperature(3, fakeSensor)
	//sensor := fakeSensor // if you want the same variable to show two results from to functions, you assign it to the first
	//fmt.Println(sensor())
	var offset kelvin = 5
	sensor := calibrate(realSensor, offset) // then = it to the second and you will get both functions outputs. But the variables that are calling it must be of the same type, 
	                    // in this instance Kelvin 
	for count := 0; count < 10; count ++ {
        fmt.Println(sensor())
	}	
}