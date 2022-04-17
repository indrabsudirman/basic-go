package main

import "fmt"

func main() {
	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			continue //Keluar dari looping, tetap melanjutkan perulangan ke i++ (post statement)
		}
		fmt.Println("Perulangan ke : ", i)

	}
}
