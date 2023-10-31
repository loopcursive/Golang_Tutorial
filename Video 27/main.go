package main
import (
	"fmt"
)

func main() {
	age := []int{1,2,3,4,5}
	fmt.Println(age)
	fmt.Println(age[4])
	age[4] = 100
	fmt.Println(age)
	age = append(age, 22,33,44)
	fmt.Println(age)
}
 