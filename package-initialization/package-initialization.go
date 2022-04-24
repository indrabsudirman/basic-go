package main

import (
	_ "basic-go/database" //tambahkan _ blank identifier untuk menghindari error
	_ "fmt"
)

func main() {
	// result := database.GetDatabase() // import dari package database, ketika dipanggil pertama func init akan dijalankan
	// fmt.Println(result)
}
