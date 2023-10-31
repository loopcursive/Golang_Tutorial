package main

import (
	"fmt"
)

func main() {
	// var num = 1

	// for num <= 10 {
	// 	fmt.Println(num)
	// 	num++
	// }

	var input_number int
	fmt.Println("Enter number here")
	fmt.Scan(&input_number)
	for i := 1; i <= 10; i++ {
		fmt.Printf("%v X %v = %v\n", input_number , i , (input_number*i))
	}

}