package app

import (
	"fmt"
	"time"
)

func (s *SidikApp) ForgotPassword() {
	ClearScreen()
	var email string
	fmt.Print("Enter your email: ")
	fmt.Scan(&email)

	for _, user := range *users {
		if user.Email == email {
			InsertNewPass(user)
			return
		} else {
			fmt.Println("Email not found.")
			time.Sleep(500 * time.Millisecond)
			fmt.Print(" .")
			time.Sleep(500 * time.Millisecond)
			fmt.Print(" .")
			time.Sleep(500 * time.Millisecond)
			fmt.Print(" .")
			time.Sleep(500 * time.Millisecond)
			s.ForgotPassword()
		}
	}
}

func InsertNewPass(user User) {
	fmt.Println("FORGOT PASSWORD")

	var password, password2 string
	fmt.Print("Insert Password :")
	fmt.Scan(&password)
	fmt.Print("Confirmation Password :")
	fmt.Scan(&password2)

	var hashed string
	if password == password2 {
		hashed = hashMD5(password)
		*users = append(*users,
			User{
				Password:  hashed,
			},
		)
	}
}
