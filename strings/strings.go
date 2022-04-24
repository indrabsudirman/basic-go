package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(strings.Contains("Indra Bayu Sudirman Banjarnegara", "Indra"))
	fmt.Println(strings.Split("Indra Bayu Sudirman Banjarnegara", " "))
	fmt.Println(strings.ToLower("Indra Bayu Sudirman Banjarnegara"))
	fmt.Println(strings.ToUpper("Indra Bayu Sudirman Banjarnegara"))
	fmt.Println(strings.Trim("        Indra Bayu Sudirman Banjarnegara       ", " "))
	fmt.Println(strings.ReplaceAll("Indra Bayu Sudirman Banjarnegara", "Indra", "Jabs"))
	fmt.Println(strings.ToLower(strings.ReplaceAll("Indra Bayu Sudirman Banjarnegara", " ", "")))
}
