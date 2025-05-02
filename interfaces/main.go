package main

import (
	"fmt"
)

func main() {
	aName := GetAdog("hero")
	fmt.Println(aName.SpeakName())
	// fmt.Println(aName.SayDogName()) cannot access because  aName is of type Name

	// using specialized interface
	specialDog := GetSpecialDog("Hisenberg")
	fmt.Println(specialDog.SayDogName())
}
