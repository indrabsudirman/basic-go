package main

import "fmt"

func getGoodBye(name string) string {
	return "Good bye " + name
}

func main() {
	sayGoodBye := getGoodBye //sayGoodBye menjadi fungsi dari getGoodBye. Dalam Go function disebut first class citizen

	result := sayGoodBye("Indra")
	fmt.Println(result)
}