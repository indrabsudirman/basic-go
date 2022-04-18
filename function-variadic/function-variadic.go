package main

import "fmt"

func sumAll(numbers ...int) int {
	total := 0
	for _, number := range numbers {
		total += number
	}
	return total
}

func main() {
	total := sumAll(10, 20, 30, 40, 50)
	fmt.Println(total)

	//bisa juga jika ingin passing slice, ke variabel argument
	slice := []int{5, 10, 15, 20, 25}
	total = sumAll(slice...)
	fmt.Println(total)
}