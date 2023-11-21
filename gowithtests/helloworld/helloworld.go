package main

import "fmt"

func Hello(s string) string {
	if s == "" {
		s = "World"
	}
	return "Hello " + s
}

func main() {
	fmt.Println(Hello("Chris"))
}
