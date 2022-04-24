package main

import (
	"flag"
	"fmt"
)

func main() {
	var host *string = flag.String("host", "localhost", "Put your database hostname")
	var user *string = flag.String("user", "root", "Put your database user")
	password := flag.String("password", "root", "Put your database password")
	pin := flag.Int("pin", 123456, "Put your database pin")

	flag.Parse()

	fmt.Println("Hostname :", *host)
	fmt.Println("Username :", *user)
	fmt.Println("Password :", *password)
	fmt.Println("Pin :", *pin)
}
