package main

import (
	"fmt"

	"example.com/greetings"
)

func main() {
	message := greetings.Hello("Assane")
	fmt.Println(message)

}
