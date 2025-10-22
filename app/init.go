package app

import (
	"fmt"
	"os"
)

type SidikApp struct{}

type Auth interface {
	Register()
	Login()
	ForgotPassword(i int)
}

type User struct {
	FirstName string
	LastName  string
	Email     string
	Password  string
}

var loggedInUser User

var users *[]User = &[]User{}

var auth Auth = &SidikApp{}

func Init() {
	defer func() {
		if r := recover(); r != nil {
			msg := "\n" + fmt.Sprint(r) + "\nClose Program. . ."
			Alert(msg)
		}
	}()

	for {
		Title("Welcome to My System")
		fmt.Println("   1. Register\n   2. Login\n   3. Forgot Password")
		fmt.Println("\n   0. Exit")
		fmt.Printf("\nChoose a menu: ")
		var input int
		fmt.Scanln(&input)

		switch input {
		case 0:
			Alert("Exiting...")
			os.Exit(0)
		case 1:
			auth.Register()
		case 2:
			auth.Login()
		case 3:
			auth.ForgotPassword(1)
		default:
			panic("Invalid input")
		}
	}
}
