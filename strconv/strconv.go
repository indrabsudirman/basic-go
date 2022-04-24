package main

import (
	"fmt"
	"strconv"
)

func main() {
	boolean, err := strconv.ParseBool("false")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(boolean)
	}

	number, err := strconv.ParseInt("100000", 10, 32)
	if err == nil {
		fmt.Println(number)
	} else {
		fmt.Println(err)
	}

	value := strconv.FormatInt(40000, 16)
	fmt.Println(value)

	value1 := strconv.Itoa(40000)
	fmt.Println(value1)
	value2, err := strconv.Atoi("4001d") //bisa juga error di ignore dengan _. Tapi harus yakin kalo tidak ada karakter selain angka
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(value2)
	}

}
