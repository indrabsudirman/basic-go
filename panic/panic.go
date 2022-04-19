package main

import "fmt"

func endApp() {
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
	runApp(true)
}
