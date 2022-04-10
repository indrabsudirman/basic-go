package main

import "fmt"

func main() {

	var names [4]string

	names[0] = "Indra"
	names[1] = "Budi"
	names[2] = "Pajri"
	names[3] = "Sri"

	fmt.Println("names are : ", names)
	fmt.Println("names[0] = ", names[0])
	fmt.Println("names[1] = ", names[1])
	fmt.Println("names[2] = ", names[2])
	fmt.Println("names[3] = ", names[3])
	fmt.Println("lenght of names = ", len(names))

	var age = [4]uint8{10, 20, 30, 40}

	fmt.Println("age are = ", age)
	fmt.Println("length of age = ", len(age))

}
