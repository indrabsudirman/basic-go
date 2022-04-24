package main

import "fmt"

type Person struct {
	Name, Address string
	Age           uint8
}

//struct function, jika ingin buat fungsi khusus untuk struct
func (person Person) sayHello(name string) {
	fmt.Println("Hello", name, "my name is", person.Name)

}

func (p Person) sayBye() {
	fmt.Println("Bye from", p.Name)
}

func main() {
	indra := Person{
		Name:    "Indra",
		Address: "Jakarta",
		Age:     10,
	}
	// fmt.Println(indra)
	// fmt.Println(indra.Name)
	// fmt.Println(indra.Address)
	// fmt.Println(indra.Age)
	indra.sayHello("Haby")
	indra.sayBye()

	lubna := Person{}
	lubna.Name = "Lubna"
	lubna.Address = "Bandung"
	lubna.Age = 20
	fmt.Println(lubna)
}
