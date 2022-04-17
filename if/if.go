package main

import "fmt"

func main() {
	name := "Indra"

	if name == "Indra" {
		fmt.Println("name is Indra")
	} else {
		fmt.Println("name is not Indra")
	}
	if length := len(name); length > 5 {
		fmt.Println("name is long")
	} else {
		fmt.Println("name is short")
	}

}
