package main
 
import (
	"fmt"
)

func main() {

	var number = 1

	switch number {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	default:
		fmt.Println("No value")
	}

}
