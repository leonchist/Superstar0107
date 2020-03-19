package main

import "fmt"

func main() {
	fmt.Println(Hello("world"))
}

func Hello(name string) string {
	return "Hello, " + name
}
