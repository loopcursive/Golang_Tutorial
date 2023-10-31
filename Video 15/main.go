package main

import (
	"fmt"
)

func main() {

	// area of rectangle

	var length int
	fmt.Println("Enter legth")
	fmt.Scan(&length)
	var width int
	fmt.Println("Enter width")
	fmt.Scan(&width)
	var answer = length * width
	fmt.Println(answer)

	var name string
	fmt.Println("Enter name")
	fmt.Scan(&name)
	var age int
	fmt.Println("Enter age")
	fmt.Scan(&age)
	fmt.Printf("Your name is %v and age is %v", name, age)

}
