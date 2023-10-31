package main

import (
	"fmt"
	"strings"
)

func main() {
	var name = "Abhi@sAAAAAAhek"
	var name_1 = "KUMAR"
	// fmt.Printf("%c" , name[7])
	// fmt.Println(len(name))
	fmt.Println(name)
	fmt.Println(len(name))
	fmt.Println(strings.Contains(name, "kum"))
	fmt.Println(strings.ToLower(name_1))
	fmt.Println(strings.ToUpper(name))
	fmt.Println(strings.Count(name, "A"))
	fmt.Println(strings.Index(name, "k"))
	fmt.Println(strings.Split(name, "@"))

}
