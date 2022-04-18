package main

import "fmt"

func getFullName() (firstName, middleName, lastName string) {
	firstName = "Indra"
	middleName = "Bayu"
	lastName = "Sudirman"

	return
}

func main() {
	a, b, c := getFullName()
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
}