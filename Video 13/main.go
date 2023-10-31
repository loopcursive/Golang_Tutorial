package main

import (
	"fmt"
)

func main() {
	// var age = 10
	// fmt.Println(age)
	var age int

	fmt.Println("Enter Your age")
	fmt.Scan(&age)
	fmt.Println(age)

	var age_2 int
	fmt.Println("Enter Your age 2")
	fmt.Scan(&age_2)

	fmt.Println(age + age_2)

}
