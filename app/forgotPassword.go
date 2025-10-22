package app

import (
	"fmt"
)

func (s *SidikApp) ForgotPassword(i int) {
	Title("FORGOT PASSWORD")

	fmt.Print("Enter your email: ")
	var email string
	fmt.Scanln(&email)

	for i, user := range *users {
		if user.Email == email {
			s.InsertNewPass(i)
			return
		}
	}
	if i == 3{
		Alert("\nUser not found!\nExitting . . .")
		return
	}
	Alert("\nUser not found!\nTry Again . . .")
	s.ForgotPassword(i+1)
}

func (s *SidikApp) InsertNewPass(i int) {
	Title("FORGOT PASSWORD")

	var password, password2 string
	fmt.Print("Insert Password :")
	fmt.Scanln(&password)
	fmt.Print("Confirmation Password :")
	fmt.Scanln(&password2)

	var hashed string
	if password == password2 {
		hashed = hashMD5(password)
		(*users)[i].Password = hashed
		Alert("\nChange Password Successfully!")
		return
	} else {
		Alert("\nYour Password not match!\nTry again. . . .")
		s.InsertNewPass(i)
	}
}
