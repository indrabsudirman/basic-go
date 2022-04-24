package main

import "fmt"

func random() interface{} {
	return false
}

func main() {
	var result = random()
	// resultString := result.(string) // type assertion, tapi harus yakin interface yang di-return dari random() adalah string
	// fmt.Println(resultString)

	//Agar lebih aman, bisa gunakan switch case daripada direct type assertion langsung
	switch value := result.(type) {
	case string:
		fmt.Println("Value", value, "is a string")
	case int:
		fmt.Println("Value", value, "is an int")
	default:
		fmt.Println("Unknown type")
	}
}
