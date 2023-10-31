package main

import (
	"fmt"
)

func main() {
	slice := []int{1, 2, 3, 4, 5}
	fmt.Println(slice)
	fmt.Println(len(slice))

	array := [5]int{1, 2, 3, 4, 5}
	slice_2 := array[0:5]
	fmt.Println(slice_2)
}
