package main

import (
	"fmt"
	"hello_world_go/pkg/greetings"
	"hello_world_go/pkg/hello"
)

func main() {
	message := hello.SayHello()
	fmt.Println(message)

	greeting := greetings.Hi("Alice")
	fmt.Println(greeting)
}
