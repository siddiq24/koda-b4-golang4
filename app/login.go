package app

import (
	"fmt"
	"time"
)

func (s *SidikApp) Login() {
	ClearScreen()
	var email, password string
	fmt.Print("Enter email: ")
	fmt.Scan(&email)
	fmt.Print("Enter password: ")
	fmt.Scan(&password)

	hashed := hashMD5(password)
	for i, user := range *users {
		if user.Email == email && user.Password == hashed {
			loggedInUser = (*users)[i]
			fmt.Println("\nLogin successful!")
			time.Sleep(500 * time.Millisecond)
			fmt.Print(" .")
			time.Sleep(500 * time.Millisecond)
			fmt.Print(" .")
			time.Sleep(500 * time.Millisecond)
			fmt.Print(" .")
			time.Sleep(500 * time.Millisecond)
			s.Home()
			return
		}
	}
	fmt.Println("Login failed. Invalid credentials.")
	time.Sleep(500 * time.Millisecond)
	fmt.Print(" .")
	time.Sleep(500 * time.Millisecond)
	fmt.Print(" .")
	time.Sleep(500 * time.Millisecond)
	fmt.Print(" .")
	time.Sleep(500 * time.Millisecond)
}
