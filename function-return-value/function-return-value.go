package main

import "fmt"

func getHello(name string) string {
	if name == "" {
		return "Hello Bro, what's your name?"
	} else {
		return "Hello " + name
	}

}

func main() {
	result := getHello("")
	fmt.Println(result)

	fmt.Println(getHello("Indra"))
}