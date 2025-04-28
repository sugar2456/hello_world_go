package greetings

import "fmt"

func Hi(name string) string {
	message := fmt.Sprintf("Hi, %v. Welcome!", name)
	return message
}
