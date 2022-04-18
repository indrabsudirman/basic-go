package main

import "fmt"

func sayHelloTo(firstName, lastName string) {
	fmt.Printf("Hello %s %s\n", firstName, lastName)
}
func main() {
	sayHelloTo("Indra", "Sudirman")
	name := "Lubna"
	sayHelloTo(name, "Sudirman")
}