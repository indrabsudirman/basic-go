package main

import "fmt"

func main() {

	var name1 = "Indra"
	var name2 = "Bayu"
	// Operator Perbandingan
	result := name1 > name2
	fmt.Println("result = ", result)

	var value1 = 10
	var value2 = 20

	fmt.Println("Value 1 = ", value1)
	fmt.Println("Value 2 = ", value2)

	fmt.Println("value1 > value2 = ", value1 > value2)
	fmt.Println("value1 < value2 = ", value1 < value2)
	fmt.Println("value1 >= value2 = ", value1 >= value2)
	fmt.Println("value1 <= value2 = ", value1 <= value2)
	fmt.Println("value1 == value2 = ", value1 == value2)
	fmt.Println("value1 != value2 = ", value1 != value2)

}