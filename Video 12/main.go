package main

import (
	"fmt"
	"strings"
)

func main() {
	var name = "AbhishekA"
	// fmt.Println(len(name))
	// fmt.Println(strings.Contains(name , "t" ))
	// fmt.Println(strings.ToLower(name))
	// fmt.Println(strings.ToUpper(name))
	// fmt.Println(strings.Index(name , "b"))
	// fmt.Println(strings.Count(name , "b"))
	// fmt.Println(strings.Split(name , "i"))
	fmt.Println(strings.Replace(name, "A", "Kumar", 2))

}
