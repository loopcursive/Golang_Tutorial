package main

import (
	"fmt"
)

func main() {
	fmt.Println("enter number")
	var input int
	fmt.Scan(&input)
	var sum = 0
	for i := 1; i <= input; i++ {
		if i%2 == 0 {
			fmt.Println(i)
			sum = sum + i
		}

	}
	fmt.Println(sum)

}
