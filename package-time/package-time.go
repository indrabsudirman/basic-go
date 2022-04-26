package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()

	fmt.Println(now)
	fmt.Println(now.Year())
	fmt.Println(now.Month())
	fmt.Println(now.Day())
	fmt.Println(now.Hour())
	fmt.Println(now.Minute())
	fmt.Println(now.Second())
	fmt.Println(now.Nanosecond())

	utc := time.Date(2022, 04, 25, 22, 02, 00, 00, time.UTC)

	fmt.Println(utc)

	// layout := time.RubyDate //"2006-01-02"
	parse, _ := time.Parse("2006-01-02", "1984-01-02")
	fmt.Println(parse)
}
