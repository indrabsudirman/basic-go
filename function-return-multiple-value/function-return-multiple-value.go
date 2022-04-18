package main

import "fmt"

func getFullName() (string, string, uint8) {
	return "Indra", "Sudirman", 49
}

func main() {
	firstName, lastName, age := getFullName()
	fmt.Println(firstName)
	fmt.Println(lastName)
	fmt.Println(age)


	//Jika ingin diabaikan return value, maka dapat menggunakan _
	firstName, _, _ = getFullName()
	fmt.Println(firstName)
}