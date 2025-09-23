package main

import (
	"fmt"
)

func main() {
	age := 32
	// agePointer := &age
	var agePointer *int
	agePointer = &age
	fmt.Println("age: ", *agePointer)
	adultAge := getAdultyYears(agePointer)
	fmt.Println(adultAge)
}

func getAdultyYears(age *int) int {
	return *age - 18
}
