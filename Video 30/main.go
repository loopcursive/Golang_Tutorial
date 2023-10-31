package main

import (
	"fmt"
)

var globalVar = 200

func number() {
	// local variabel
	var num = 100
	fmt.Println(num)
	fmt.Println(globalVar)

}

func f2() {
	// fmt.Println(num)
	fmt.Println(globalVar)
}

func main() {

	number()
	// fmt.Println(num)

	f2()
	fmt.Println(globalVar)

}
