package main

import "fmt"

func main() {
	//type declaration membuat tipe baru dari tipe data yang sudah ada
	type NoKTP string
	type Married bool

	var noktp NoKTP = "123456789"
	var married Married = true

	fmt.Println("noktp = ", noktp)
	fmt.Println("married = ", married)

	
}