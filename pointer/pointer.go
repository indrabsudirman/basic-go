package main

import "fmt"

type Address struct {
	City, Province, Country string
}

//Pointer function
func changeCountryToIndonesia(address *Address) { //*Address merupakan pointer
	address.Country = "Indonesia"
}

func main() {
	address1 := Address{"Tangerang Selatan", "Banten", "Indonesia"}
	//bisa juga seperti ini
	// var address1  Address = Address{"Tangerang Selatan", "Banten", "Indonesia"}
	// var address2 *Address = &address1 //tanda * dan & digunakan untuk pointer dan reference

	address2 := &address1 //& merupakan operator pemanggilan pointer
	address3 := address1  //menyalin nilai address1 ke address3 (pass by value)) tidak menggunakan pointer yang pass by reference
	address4 := &address1 //& merupakan operator pemanggilan pointer

	address2.City = "Kota Tangerang"

	*address4 = Address{"Jakarta", "DKI Jakarta", "Indonesia"}

	fmt.Println(address1)
	fmt.Println(address2)
	fmt.Println(address3)
	fmt.Println(address4)

	var address5 *Address = new(Address) //new membuat sebuah pointer baru(kosong)
	address5.City = "Banjarnegara"
	fmt.Println(address5)

	var alamat = Address{"Jakarta", "DKI Jakarta", ""}

	//Bisa juga buat pointer seperti ini
	alamatPointer := &alamat //bisa juga : var alamatPointer *Address = &alamat

	changeCountryToIndonesia(alamatPointer) //&alamat merupakan pointer atau panggil alamatPointer
	fmt.Println(alamat)                     //Country tidak berubah karena tidak menggunakan pointer
}
