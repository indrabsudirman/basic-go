package main

import (
	"container/list"
	"fmt"
)

func main() {
	data := list.New()

	data.PushBack("Indra")
	data.PushBack("Bayu")
	data.PushBack("Sudirman")
	data.PushFront("Ridwan")

	fmt.Println(data.Len())
	fmt.Println(data.Front().Value)
	fmt.Println(data.Back().Value)

	//Jika ingin iterasi list dari depan
	for element := data.Front(); element != nil; element = element.Next() {
		fmt.Print(element.Value, " ")
	}
	fmt.Println()
	//Jika ingin iterasi list dari belakang
	for element := data.Back(); element != nil; element = element.Prev() {
		fmt.Print(element.Value, " ")
	}
}
