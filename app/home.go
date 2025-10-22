package app

import (
	"fmt"
	"time"
)

func (s *SidikApp) Home() {
	for {
		ClearScreen()
		fmt.Println("========================================")
		fmt.Println("WELCOME TO HOME")
		fmt.Println("========================================")
		fmt.Println("\n1. Show all users\n2. Logout\n\n0. Exit")
		var choosed int
		fmt.Print("\nChoose a menu: ")
		fmt.Scan(&choosed)
		switch choosed {
		case 1:
			s.ShowUserList()
		case 2:
			for i, user := range *users {
				if user == loggedInUser {
					*users = append((*users)[:i], (*users)[i+1:]...)
				}
			}
			fmt.Print("Loggout")
			time.Sleep(500 * time.Millisecond)
			fmt.Print(" .")
			time.Sleep(500 * time.Millisecond)
			fmt.Print(" .")
			time.Sleep(500 * time.Millisecond)
			fmt.Print(" .")
			time.Sleep(500 * time.Millisecond)
			return
		case 0:
			fmt.Print("Exitting")
			time.Sleep(500 * time.Millisecond)
			fmt.Print(" .")
			time.Sleep(500 * time.Millisecond)
			fmt.Print(" .")
			time.Sleep(500 * time.Millisecond)
			fmt.Print(" .")
			time.Sleep(500 * time.Millisecond)
			return
		}
	}
}

func (s *SidikApp) ShowUserList() {
	ClearScreen()
	fmt.Println("\n--- Show All Users ---")
	for i, user := range *users {
		fmt.Println(i)
		fmt.Printf("Full Name: %s %s\n", user.FirstName, user.LastName)
		fmt.Printf("Email    : %s\n", user.Email)
		fmt.Printf("Password : %s\n\n", user.Password)
	}
	fmt.Println("------------------------")
	fmt.Print("\nInput anything to exit: ")
	var i int
	fmt.Scan(&i)
	if i == 0 {
		fmt.Print("\nExiting")
		time.Sleep(500 * time.Millisecond)
		fmt.Print(" .")
		time.Sleep(500 * time.Millisecond)
		fmt.Print(" .")
		time.Sleep(500 * time.Millisecond)
		fmt.Print(" .")
		time.Sleep(500 * time.Millisecond)
		return
	}
}
