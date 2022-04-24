package main

import "fmt"

type HasName interface {
	GetName() string
}

func sayHello(hasName HasName) {
	fmt.Println("Hello", hasName.GetName())

}

type Person struct {
	name string
}

func (p Person) GetName() string {
	return p.name

}

type Employee struct {
	Name string
}

func (e Employee) GetName() string {
	return e.Name

}

func main() {
	indra := Person{
		name: "Indra",
	}
	sayHello(indra)

	e := Employee{
		Name: "Haby",
	}

	sayHello(e)
}
