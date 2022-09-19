package main

import "fmt"

// This function will convert a integer to a floating point number
func CalculateWorkingCarsPerHour(productionRate int, successRate float64) float64 {
	productionRate = 10
	successRate = 91.2
	result := float64(productionRate) * (successRate / 100) // Here you have to put the type you want on the oposite type number so I have put float64 on an int because I want a float64 number.
	return result
	//panic("CalculateWorkingCarsPerHour not implemented")
}

// CalculateWorkingCarsPerMinute calculates how many working cars are
// produced by the assembly line every minute.
// This function will convert a floatingpoint of 64 to an integer

func CalculateWorkingCarsPerMinute(productionRate int, successRate float64) int {
	perHour := float64(productionRate) * (successRate / 100)
	return int(perHour) / 60 // Here I am converting a float 64 number to an int, because I am returning an int out of the function. So I have put int on the outgoing float64 value
	//panic("CalculateWorkingCarsPerMinute not implemented")
}

func CalculateCost(carsCount int) uint {
	const tenCarcost = 95000
	const oneCar = 10000
	resultOfCarsCount := carsCount / 10
	groupOfCars := carsCount % 10
	result := ((uint64)(resultOfCarsCount * tenCarcost)) + ((uint64)(groupOfCars * oneCar))
	return uint(result)
	panic("CalculateCost not implemented")
}

func main() {
	fmt.Println("Cars made per hour is,", CalculateWorkingCarsPerHour(10, 91.2))
	fmt.Println("Cars made per minute is,", CalculateWorkingCarsPerMinute(1105, 90.0))
	fmt.Println("Cost to Makes Cars,", CalculateCost(37))
}
