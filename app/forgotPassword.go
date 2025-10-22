package app

import (
	"fmt"
	"time"
)

func (s *SidikApp) ForgotPassword() {
	ClearScreen()
	fmt.Println("--- FORGOT PASSWORD ---")
	fmt.Print("\nEnter your email: ")

	var email string
	fmt.Scan(&email)

	for i, user := range *users {
		if user.Email == email {
			s.InsertNewPass(i)
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

func (s *SidikApp) InsertNewPass(i int) {
	ClearScreen()
	fmt.Println("--- FORGOT PASSWORD ---")

	var password, password2 string
	fmt.Print("\nInsert Password :")
	fmt.Scan(&password)
	fmt.Print("Confirmation Password :")
	fmt.Scan(&password2)

	var hashed string
	if password == password2 {
		hashed = hashMD5(password)
		(*users)[i].Password = hashed
		fmt.Print("\n Change Password Successfully!")
		time.Sleep(500 * time.Millisecond)
		fmt.Print(" .")
		time.Sleep(500 * time.Millisecond)
		fmt.Print(" .")
		time.Sleep(500 * time.Millisecond)
		fmt.Print(" .")
		time.Sleep(500 * time.Millisecond)
		return
	} else {
		fmt.Print("\nYour Password not match!")
		time.Sleep(500 * time.Millisecond)
		fmt.Print(" .")
		time.Sleep(500 * time.Millisecond)
		fmt.Print(" .")
		time.Sleep(500 * time.Millisecond)
		fmt.Print(" .")
		time.Sleep(500 * time.Millisecond)
		s.InsertNewPass(i)
	}
}
