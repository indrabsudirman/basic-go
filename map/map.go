package main

import "fmt"

func main() {
	var person = map[string]string{
		"firstName": "Indra",
		"lastName":  "Bayu",
		"age":       "41",
	}

	//menambahkan data baru di map
	person["address"] = "Tangerang Selatan"

	fmt.Println(person)
	fmt.Println(person["firstName"])
	fmt.Println(person["lastName"])
	fmt.Println(person["age"])
	fmt.Println(person["address"])

	book := make(map[string]string)
	book["title"] = "Belajar Go basic"
	book["author"] = "Indra Sudirman"
	book["publisher"] = "Jabrikos Ltd"
	book["year"] = "2022"

	fmt.Println(book)
	fmt.Println(len(book))
	delete(book, "year")
	fmt.Println(book)
	fmt.Println(len(book))
}