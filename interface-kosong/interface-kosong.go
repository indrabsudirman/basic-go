package main

import "fmt"

func Ups(i int) interface{} {
	if i == 1 {
		return 1
	} else if i == 2 {
		return "dua"
	} else {
		return true
	}
}

func main() {
	data := Ups(1)
	fmt.Println(data)

	var data1 interface{} = Ups(2)
	fmt.Println(data1)
}
