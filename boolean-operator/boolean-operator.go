package main

import "fmt"

func main() {
	var ujian = 80
	var absensi = 75

	var lulusUjian = ujian >= 80
	var lulusAbsensi = absensi >= 80
	fmt.Println("Lulus Ujian = ", lulusUjian)
	fmt.Println("Lulus Absensi = ", lulusAbsensi)

	var lulus = lulusUjian && lulusAbsensi
	fmt.Println("lulus = ", lulus)
}
