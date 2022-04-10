package main

import "fmt"

func main() {
	var nilai32 int32 = 129
	var nilai64 int64 = int64(nilai32)
	var nilai8 int8 = int8(nilai32)

	fmt.Println("nilai32 = ", nilai32)
	fmt.Println("nilai64 = ", nilai64)
	fmt.Println("nilai8 = ", nilai8) //melebihi kapasitas nilai int8

	var name = "Indra"
	var i = name[0]
	var iString = string(i)

	fmt.Println("name = ", name)
	fmt.Println("i = ", iString)
}