package main

import "fmt"

func main() {
	var months = [...]string{
		"Januari",
		"Februari",
		"Maret",
		"April",
		"Mei",
		"Juni",
		"Juli",
		"Agustus",
		"September",
		"Oktober",
		"November",
		"Desember",
	}

	var slice = months[1:4]
	fmt.Println("slice = ", slice)
	fmt.Println("length of slice = ", len(slice))
	fmt.Println("capacity of slice = ", cap(slice))

	months[3] = "dirubah"
	fmt.Println("slice month yang dirubah = ", slice)

	slice[0] = "dirubahTrus"
	fmt.Println("month yang dirubah = ", months)

	//membuat slice baru dari months
	var slice2 = months[10:]
	fmt.Println("slice2 = ", slice2)

	var slice3 = append(slice2, "bulanBaru")
	fmt.Println("slice3 = ", slice3)

	//membuat slice baru
	slice4 := make([]string, 2, 5)
	slice4[0] = "Januari"
	slice4[1] = "Februari"
	fmt.Println("slice4 = ", slice4)
	fmt.Println("length of slice4 = ", len(slice4))
	fmt.Println("capacity of slice4 = ", cap(slice4))

	//copy slice ke slice baru
	slice5 := make([]string, len(slice4), cap(slice4))
	copy(slice5, slice4)
	fmt.Println("slice5 hasil copy dari slice4 = ", slice5)
	fmt.Println("length of slice5 = ", len(slice5))
	fmt.Println("capacity of slice5 = ", cap(slice5))

}
