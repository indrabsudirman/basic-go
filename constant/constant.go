package main

import "fmt"

func main() {
	const firstName string = "Indra" //Tipe data constant, jika tidak dipakai tidak akan error
	const lastName = "Bayu"
	const age = 41
	const isMarried = true
	const address = "Tangerang Selatan"

	//Bisa juga deklarasi multiple constant seperti ini
	const (
		name1 = "Indra"
		age1  = 41
		
	)

	fmt.Println("Nama saya adalah ", firstName, lastName)
	fmt.Println("Umur saya adalah ", age)
	fmt.Println("Status saya adalah ", isMarried)
	fmt.Println("Alamat saya adalah ", address)

	fmt.Println("Nama dan usia dari multiple constant : ", name1, age1)


}