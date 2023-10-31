package main

import (
	"fmt"
)

func main() {

	user_name := "Abhishek@123"
	user_password := "1234"

	var input_user_name string
	var input_user_password string

	fmt.Println("Enter your user name")
	fmt.Scan(&input_user_name)
	fmt.Println("Enter your password")
	fmt.Scan(&input_user_password)

	if user_name == input_user_name && user_password == input_user_password {
		fmt.Println("Authentication successful")
	} else {
		fmt.Println("Authentication fail")
	}
}
