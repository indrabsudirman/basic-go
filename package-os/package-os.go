package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args
	fmt.Println("Argument :")
	fmt.Println(args)

	fmt.Println(args[0])
	fmt.Println(args[1])
	fmt.Println(args[2])

	hostName, err := os.Hostname()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Hostname :", hostName)
	}

	userName := os.Getenv("APP_USERNAME")
	password := os.Getenv("APP_PASSWORD")
	fmt.Println("Username :", userName)
	fmt.Println("Password :", password)

}
