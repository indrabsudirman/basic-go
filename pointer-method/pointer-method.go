package main

import "fmt"

type Man struct {
	Name string
}

func (man *Man) Married() { //Jika buat struct func atau method menggunakan pointer * agar tidak bengkak dimemory
	man.Name = "Mr. " + man.Name
	fmt.Println("Di method", man.Name)
}

func main() {
	indra := Man{"Indra"}
	indra.Married()
	fmt.Println(indra.Name)
}
