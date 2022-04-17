package main

import "fmt"

func main() {
	counter := 1

	for counter <= 10 {
		fmt.Println("For model 1, Perulangan ke : ", counter)
		counter++
	}

	//Bisa disederhanakan menjadi, dengan init statement dan post statement
	for counter := 1; counter <= 10; counter++ {
		fmt.Println("For model 2, Perulangan ke : ", counter)
	}

	slice := []string{"Indra", "Bayu", "Sudirman", "Budi"}
	//for juga bisa digunakan untuk iterasi slice atau array
	for i := 0; i < len(slice); i++ {
		fmt.Println("For model 3, Perulangan ke : ", i, " dengan nilai ", slice[i])
	}

	//Bisa juga pakai for range
	for i, v := range slice {
		fmt.Println("For model 4, Perulangan ke : ", i, " dengan nilai ", v)
	}
	//Sederhananya jika indexnya tidak diperlukan
	for _, v := range slice {
		fmt.Println(v)
	}

	//Bisa juga pakai for range untuk map
	person := map[string]string{
		"name":    "Indra",
		"address": "Jl. Kebon Kacang",
		"age":     "20",
	}
	for key, value := range person {
		fmt.Println("For model 5 (map), Perulangan ke : ", key, " dengan nilai ", value)
	}

}
