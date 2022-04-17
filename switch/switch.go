package main

import "fmt"

func main() {

	name := "Indra"

	switch name {
	case "Indra":
		fmt.Println("name is Indra")
	case "Budi":
		fmt.Println("name is Budi")
	default:
		fmt.Println("name is not Indra")
	}

	switch length := len(name); length {
	case 1:
		fmt.Println("name is short")
	case 2:
		fmt.Println("name is medium")
	case 3:
		fmt.Println("name is long")
	default:
		fmt.Println("name is very long")

	}

	length := len(name)
	//switch tanpa kondisi
	switch {
	case length > 5:
		fmt.Printf("name %v is long\n", name)
	case length > 3:
		fmt.Printf("name %v is short\n", name)
	default:
		fmt.Printf("name %v is short\n", name)
	}
}
