package main

import (
	"fmt"
	"time"
)

type User struct {
	firstName string
	lastName  string
	birthdate string
	createdAt time.Time
}

func main() {
	ufirstName := getUserData("Please enter your first name: ")
	ulastName := getUserData("Please enter your last name: ")
	ubirthdate := getUserData("Please enter your birthdate (MM/DD/YYYY): ")

	// ... do something awesome with that gathered data!
	var appUser User
	appUser = User{
		// ufirstName,
		// ulastName,
		// ubirthdate,
		// time.Now(),
		firstName: ufirstName,
		lastName:  ulastName,
		// birthdate: ubirthdate,
		createdAt: time.Now(),
	}
	// fmt.Println(firstName, lastName, birthdate)
	outputUserDetails(&appUser)
}

// func outputUserDetails(firstName, lastName, birthdate string) {
// 	// ...
// 	fmt.Println(firstName, lastName, birthdate)
// }

func outputUserDetails(u *User) {
	// ...
	//(*u).birthdate dereference --
	fmt.Println((*u).firstName, u.lastName, u.birthdate) //dang viet tat duoc go chap nhan
}

func getUserData(promptText string) string {
	fmt.Print(promptText)
	var value string
	fmt.Scan(&value)
	return value
}
