package main

import (
	"fmt"

	"github.com/pncabido/golearning/greetings"
)

func main() {
	message := greetings.Hello("Pedro")
	fmt.Println(message)
}
