package main

import (
	"fmt"
	"reflect"
)

type Sample struct {
	Name string
}

func main() {
	sample := Sample{Name: "Indra"}

	var sampleType reflect.Type = reflect.TypeOf(sample)

	fmt.Println(sampleType.NumField())
	fmt.Println(sampleType.Field(0).Name)

}
