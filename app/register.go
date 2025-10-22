package app

import (
	"fmt"
	"time"
)

func (s *SidikApp) Register() {
	ClearScreen()
	var firstName, lastName, email, password, password2 string
	fmt.Println("--- REGISTER ---")
	fmt.Print("Enter your first name: ")
	fmt.Scan(&firstName)
	fmt.Print("Enter your last name: ")
	fmt.Scan(&lastName)
	fmt.Print("Enter your email: ")
	fmt.Scan(&email)
	fmt.Print("Enter password: ")
	fmt.Scan(&password)
	fmt.Print("Confirmation password: ")
	fmt.Scan(&password2)

	var hashed string
	if password == password2 {
		hashed = hashMD5(password)
		*users = append(*users,
			User{
				FirstName: firstName,
				LastName:  lastName,
				Email:     email,
				Password:  hashed,
			},
		)

		fmt.Print("\nUser registered successfully!")
		time.Sleep(500 * time.Millisecond)
		fmt.Print(" .")
		time.Sleep(500 * time.Millisecond)
		fmt.Print(" .")
		time.Sleep(500 * time.Millisecond)
		fmt.Print(" .")
		time.Sleep(500 * time.Millisecond)
	}
	fmt.Print("\nYour Password is difference!")
	time.Sleep(500 * time.Millisecond)
	fmt.Print(" .")
	time.Sleep(500 * time.Millisecond)
	fmt.Print(" .")
	time.Sleep(500 * time.Millisecond)
	fmt.Print(" .")
	time.Sleep(500 * time.Millisecond)
}
