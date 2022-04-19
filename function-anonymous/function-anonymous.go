package main

import "fmt"

type Blacklist func(string) bool

func registerUser(name string, blacklist Blacklist) {
	if blacklist(name) {
		fmt.Println("User", name, "is blacklisted")
	} else {
		fmt.Println("Welcome", name)
	}

}
func main() {
	blacklist := func(name string) bool { // bisa buat anonymous function, kemudian dimasukkan ke variable
		return name == "indra"
	}

	registerUser("admin", blacklist)
	registerUser("indra", blacklist)
}
