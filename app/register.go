package app

import (
	"fmt"
)

func (s *SidikApp) Register() {
	var newUser User
	var password1, password2 string

	Title("REGISTER")
	fmt.Print("Enter your first name: ")
	fmt.Scanln(&newUser.FirstName)
	fmt.Print("Enter your last name : ")
	fmt.Scanln(&newUser.LastName)
	fmt.Print("Enter your email     : ")
	fmt.Scanln(&newUser.Email)
	fmt.Print("Enter password       : ")
	fmt.Scanln(&password1)
	fmt.Print("Confirmation password: ")
	fmt.Scanln(&password2)

	if password1 == password2 {
		for _, user := range *users {
			if user.Email == newUser.Email {
				Alert("\nYour Email Alredy Exist. . .")
				s.Login()
				return
			}
		}
		newUser.Password = hashMD5(password1)
		s.Validate(newUser)
		Alert("\nUser registered successfully! . . . .")

	} else {
		Alert("\nYour Password not match!\n\nTry Again . . . .")
		s.Register()
	}
}

func (s *SidikApp) Validate(usr User) {
	Title("REGISTER")

	var input string

	fmt.Println("Is it true?")
	fmt.Println("\n   First Name :", usr.FirstName)
	fmt.Println("   Last Name  :", usr.LastName)
	fmt.Println("   Email      :", usr.Email)
	fmt.Print("\nContinue (y/n): ")
	fmt.Scanln(&input)
	switch input {
	case "y":
		*users = append(*users, usr)
		return
	case "n":
		s.Register()
	default:
		s.Validate(usr)
	}

}
