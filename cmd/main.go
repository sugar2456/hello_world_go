package main

import (
	"fmt"
	"hello_world_go/pkg/greetings"
	"hello_world_go/pkg/hello"
	"log"
)

func main() {
	message := hello.SayHello()
	fmt.Println(message)

	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	greeting, err := greetings.Hi("Gladys")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(greeting)
}
