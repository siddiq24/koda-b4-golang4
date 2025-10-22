package app

import (
	"fmt"
	"os"
)

type SidikApp struct{}

type UserAction interface {
	Register()
	Login()
	ForgotPassword()
}

type User struct {
	FirstName string
	LastName string
	Email string
	Password string
}

var loggedInUser User

var users *[]User = &[]User{}

func Init() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println(r)
		}
	}()

	system := &SidikApp{}

	for {
		ClearScreen()
		fmt.Println("========================================")
		fmt.Println("Welcome to My System")
		fmt.Println("\n1. Register\n2. Login\n3. Forgot Password")
		fmt.Println("\n0. Exit")
		fmt.Printf("\nChoose a menu: ")
		var choosed int
		fmt.Scan(&choosed)

		switch choosed {
		case 0:
			fmt.Println("Exiting...")
			os.Exit(0)
		case 1:
			system.Register()
		case 2:
			system.Login()
		case 3:
			system.ForgotPassword()
		default:
			panic("Invalid menu choice")
		}
	}
}
