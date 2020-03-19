package main

import "fmt"

func main() {
	fmt.Println(Hello("world"))
}

const englishHelloPrefix = "Hello, "

func Hello(name string) string {
	if name == "" {
		name = "World"
	}
	return englishHelloPrefix + name
}
