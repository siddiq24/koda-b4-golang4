package app

import (
	"fmt"
)

func (s *SidikApp) Login() {
	var email, password string
	Title("LOGIN")
	fmt.Print("Enter email: ")
	fmt.Scanln(&email)
	fmt.Print("Enter password: ")
	fmt.Scanln(&password)

	hashed := hashMD5(password)
	for i, user := range *users {
		if user.Email == email {
			if user.Password == hashed {
				loggedInUser = (*users)[i]
				Alert("\nLogin successful . . .")
				s.Home()
				return
			} else {
				Alert("\nEmail or password invalid. . .")
				s.Login()
				return
			}
		}
	}
	Alert("\nUser Not Found . . . .")
	s.Register()
}
