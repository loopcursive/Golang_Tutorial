package main

import (
	"fmt"
)

func ret(x int) (z int) {
	z = x + 100
	return
	// return x + 100
}

func two(age int, name string) (a int, b string) {
	a = age
	b = name
	return
}

func main() {

	fmt.Println(ret(100))
	fmt.Println(two(2, "Abhi"))

	val_1, val_2 := two(12, "Abhishek")
	fmt.Println(val_1)
	fmt.Println(val_2)

	_, val_4 := two(120, "Abhishek kumar")
	fmt.Println(val_4)

}
