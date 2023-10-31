package main
import("fmt")
func main(){

	var number_1 = 10
	var number_2 = 20

	var sum = number_1+number_2
	var sub = number_1-number_2
	var multi = number_1*number_2
	var div = number_2/number_1

	fmt.Println(div)
	fmt.Println(sub)
	fmt.Println(sum)
	fmt.Println(multi)

	fmt.Println(number_1+number_2)

	fmt.Println("Assign")
	var name = "Abhishek"
	fmt.Println(name)

	var num_1 = 10
	// num_1 -= 20
	fmt.Println(num_1)

	var num_2 = 10
	fmt.Println(num_1 == num_2)
	fmt.Println(num_1 > num_2)
	fmt.Println(num_1 < num_2)
	fmt.Println(num_1 <= num_2)
	fmt.Println(num_1 >= num_2)
	
}