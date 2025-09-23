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

func (u *User) outputUserDetails() { //method của User struct
	fmt.Println(u.firstName, u.lastName, u.birthdate)
}

func (u *User) clearUserName() {
	u.firstName = ""
	u.lastName = ""
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
		birthdate: ubirthdate,
		createdAt: time.Now(),
	}
	// fmt.Println(firstName, lastName, birthdate)
	// outputUserDetails(&appUser) //truyen con tro
	appUser.outputUserDetails() //method struct
	appUser.clearUserName()
	appUser.outputUserDetails()
}

// func outputUserDetails(firstName, lastName, birthdate string) {
// 	// ...
// 	fmt.Println(firstName, lastName, birthdate)
// }

// func outputUserDetails(u *User) { //truyen con tro
// 	// ...
// 	//(*u).birthdate dereference --
// 	fmt.Println((*u).firstName, u.lastName, u.birthdate) //dang viet tat duoc go chap nhan
// }

func getUserData(promptText string) string {
	fmt.Print(promptText)
	var value string
	fmt.Scan(&value)
	return value
}
