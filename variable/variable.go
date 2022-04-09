package main

import "fmt"

func main() {
	var name string

	name = "Indra"
	fmt.Println(name)

	name = "Indra Bayu"
	fmt.Println(name)

	var childName = "Lubna"
	fmt.Println(childName)

	var age uint8 = 41
	fmt.Println(age)

	village := "Tangerang Selatan"
	fmt.Println(village)

	var (
		name1 = "Indra"
		age1  = 41
	)
	fmt.Println(name1)
	fmt.Println(age1)
}