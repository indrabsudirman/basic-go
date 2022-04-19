package main

import "fmt"

func main() {
	name := "Indra"
	counter := 0

	increment := func() {
		name := "Bayu" //closure variable, agar tidak merubah nilai name di luar fungsi, tambahkan var/:=
		counter++
		fmt.Println(name, counter)
	}

	increment()
	increment()

	fmt.Println(counter)
	fmt.Println(name)
}
