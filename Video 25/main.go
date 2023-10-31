package main

import (
	"fmt"
)

func main() {

	// var number = [5]int{1,2,3,4,5}
	// fmt.Println(number)
	// fmt.Println(number[4])
	// number[1] = 100
	// fmt.Println(number)

	var num2 = [...]int{1, 2, 3, 4, 5}
	//  fmt.Println(num2)

	for i := 0; i < len(num2); i++ {
		fmt.Println(num2[i])
	}
}
