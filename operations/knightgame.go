package main

import "fmt"

func CanFastAttack(knightIsAwake bool) bool {
	canfastattack := knightIsAwake != true
	return canfastattack
	panic("Please implement the CanFastAttack() function")
}

// CanSpy can be executed if at least one of the characters is awake.
func CanSpy(knightIsAwake, archerIsAwake, prisonerIsAwake bool) bool {
	canspyhappen := knightIsAwake || archerIsAwake || prisonerIsAwake == true
	return canspyhappen
	panic("Please implement the CanSpy() function")
}

// CanSignalPrisoner can be executed if the prisoner is awake and the archer is sleeping.
func CanSignalPrisoner(archerIsAwake, prisonerIsAwake bool) bool {
	canSignalPrisonerhappen := prisonerIsAwake == true && archerIsAwake == false
	return canSignalPrisonerhappen
	panic("Please implement the CanSignalPrisoner() function")
}

// CanFreePrisoner can be executed if the prisoner is awake and the other 2 characters are asleep
// or if Annalyn's pet dog is with her and the archer is sleeping.
func CanFreePrisoner(knightIsAwake, archerIsAwake, prisonerIsAwake, petDogIsPresent bool) bool {
	//var canfreeprisonerhappen bool
	//if petDogIsPresent == true && archerIsAwake == false {
	//	canfreeprisonerhappen = true
	//} else if prisonerIsAwake == true && knightIsAwake == false && archerIsAwake == false {
	//	canfreeprisonerhappen = true
	//}
	canfreeprisonerhappen := (petDogIsPresent && !archerIsAwake) || (prisonerIsAwake && !knightIsAwake && !archerIsAwake)
	return canfreeprisonerhappen

	panic("Please implement the CanFreePrisoner() function")
}

func main() {
	var knightIsAwake = false
	var archerIsAwake = false
	var prisonerIsAwake = true
	var petDogIsPresent = false

	//fmt.Println(CanFastAttack(knightIsAwake))
	//fmt.Println(CanSpy(knightIsAwake, archerIsAwake, prisonerIsAwake))
	//fmt.Println(CanSignalPrisoner(archerIsAwake, prisonerIsAwake))
	fmt.Println(CanFreePrisoner(knightIsAwake, archerIsAwake, prisonerIsAwake, petDogIsPresent))
}
