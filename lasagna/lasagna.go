package main

import "fmt"

func RemainingOvenTime(OvenTime, remaining int) (Remainingtime int) {
	Remainingtime = OvenTime - remaining
	return Remainingtime
}

func NumberOfLayers(NumberOflasLayers int) int {
	return NumberOflasLayers * 2
}

func ElapsedTime(NumberOfLayers, RemainingOvenTime int) int {
	return NumberOfLayers*2 + RemainingOvenTime
}

func main() {
	const OvenTime = 40
	const Minutes = 2
	NumberOflasLayers := 2
	remaining := 30
	fmt.Println("The remaining time left to cook the lasagna is,", RemainingOvenTime(OvenTime, remaining), "Minutes")
	fmt.Println("The time to prepare was,", NumberOfLayers(NumberOflasLayers), "Minutes")
	fmt.Println("The Total time to prepare the lasgne today took,", ElapsedTime(NumberOfLayers(OvenTime), RemainingOvenTime(OvenTime, remaining)), "Minutes")
}
