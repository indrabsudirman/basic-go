package main

import "fmt"

func endApp() {
	//recover() diletakkan di defer function
	message := recover() // recover() akan mengembalikan nilai yang terjadi pada panic()
	if message != nil {
		fmt.Println("error message = ", message)
	}
	fmt.Println("Aplikasi Selesai")
}

func runApp(error bool) {
	defer endApp()
	if error {
		panic("Terjadi kesalahan")
	}
	fmt.Println("Aplikasi berjalan")
}

func main() {
	runApp(false)
}
