package login

import (
	"fmt"
	"simple_selling/admin"
	"simple_selling/user"
)

func takeInput() (string, string) {
	fmt.Print("Enter your name: ")
	var input string
	fmt.Scan(&input)
	fmt.Print("Enter your Password: ")
	var pass string
	fmt.Scan(&pass)
	return input, pass
}
func LoginMain() {
	fmt.Println("Please login first")
	fmt.Println("1.Login as user")
	fmt.Println("2.Login as Admin")
	fmt.Println("Select your choice")
	var choice int
	fmt.Scanf("%d", &choice)
	switch choice {
	case 1:
		name, pass := takeInput()
		exist := auth(name, pass)
		if exist {
			fmt.Println("Login success")
			user.UserMain()
		} else {
			fmt.Println("Login fail")
		}
	case 2:
		name, pass := takeInput()
		exist := auth(name, pass)
		if exist {
			fmt.Println("Login success")
			admin.AdminMain()
		} else {
			fmt.Println("Login fail")
		}
	default:
		fmt.Println("Invalid choice")
	}
}
