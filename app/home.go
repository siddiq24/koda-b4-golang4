package app

import (
	"fmt"
)

func (s *SidikApp) Home() {
	for {
		Title("  HOME  ")

		fmt.Println("\n   1. Show all users\n   2. Logout\n\n   0. Exit")
		var input int
		fmt.Print("\nChoose a menu: ")
		fmt.Scanln(&input)
		switch input {
		case 1:
			s.ShowUserList()
		case 2:
			for i, user := range *users {
				if user == loggedInUser {
					*users = append((*users)[:i], (*users)[i+1:]...)
				}
			}
			Alert("Loggout Successfull. . . .")
			return
		case 0:
			Alert("Exitting. . . .")
			return
		default:
			s.Home()
		}
	}
}

func (s *SidikApp) ShowUserList() {
	Title("  HOME  ")

	fmt.Println("")
	for i, user := range *users {
		fmt.Println(i + 1)
		fmt.Printf("Full Name: %s %s\n", user.FirstName, user.LastName)
		fmt.Printf("Email    : %s\n", user.Email)
		fmt.Printf("Password : %s\n\n", user.Password)
	}
	fmt.Println("------------------------")
	Alert("\nPress enter to exit: ")
	var i string
	fmt.Scanln(&i)
}
