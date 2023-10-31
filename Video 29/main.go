package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello")
	fmt.Println(add(7))

	var value = add(10)
	fmt.Println(value+100)
	
}

func add(x int)int{
	square := x*x
	return square
// 
}
