package main

import "fmt"

func main() {

	result := 10 + 10
	fmt.Println(result)

	var a = 10
	var b = 20
	var c = a * b
	fmt.Println(c)

	var i = 10
	i += 10 //i = i + 10 ini namanya augmented assignment
	i *= 10 //i = i * 10 = 200
	fmt.Println("hasil i : ",i)

	i++
	fmt.Println("hasil i++ : ",i)

	var negatif = -10
	var positif = +10 //ini namanya unary operator, secara default nilainya sudah positif tanpa tanda +
	fmt.Println("hasil negatif : ",negatif)
	fmt.Println("hasil positif : ",positif)

}