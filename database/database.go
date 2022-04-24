package database

import "fmt"

var connection string

func init() { // func init akan dijalankan pertama kali ketika package database di import
	connection = "localhost"
	fmt.Println("database connected")
}

func GetDatabase() string {
	return connection
}
