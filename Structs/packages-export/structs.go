package main

import (
	"fmt"

	"example.com/structs/user"
)

func main() {
	ufirstName := getUserData("Please enter your first name: ")
	ulastName := getUserData("Please enter your last name: ")
	ubirthdate := getUserData("Please enter your birthdate (MM/DD/YYYY): ")

	// ... do something awesome with that gathered data!
	var appUser *user.User

	appUser, err := user.New(ufirstName, ulastName, ubirthdate)
	if err != nil {
		fmt.Println("Error!")
		return
	}

	appUser.OutputUserDetails() //method struct
	appUser.ClearUserName()
	appUser.OutputUserDetails()
}

func getUserData(promptText string) string {
	fmt.Print(promptText)
	var value string
	fmt.Scanln(&value)
	return value
}
