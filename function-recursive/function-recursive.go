package main

import "fmt"

//Factorial dengan looping
func factorialLoop(value int) int {
	result := 1
	for i := value; i > 0; i-- {
		result *= i
	}
	return result
}

//Factorial dengan recursion
func factorialRecursion(value int) int {
	if value == 1 {
		return 1
	} else {
		return value * factorialRecursion(value-1)
	}
}

func main() {
	fmt.Println(factorialLoop(10))
	fmt.Println(factorialRecursion(10))
}
