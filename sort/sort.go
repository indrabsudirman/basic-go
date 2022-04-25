package main

import (
	"fmt"
	"sort"
)

type User struct {
	name string
	age  int
}

type UsersSlices []User

func (value UsersSlices) Len() int {
	return len(value)
}

func (value UsersSlices) Less(i, j int) bool {
	return value[i].age < value[j].age

}

func (value UsersSlices) Swap(i, j int) {
	value[i], value[j] = value[j], value[i]
}

func main() {
	users := []User{
		{"Indra", 40},
		{"Bayu", 11},
		{"Sudirman", 19},
	}

	fmt.Println(users)
	sort.Sort(UsersSlices(users))
	fmt.Println(users)
}
