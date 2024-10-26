package greetings

import "fmt"

// type comes in after the name of the parameter. Return type is also after
func Hello(name string) string {
	message := fmt.Sprintf("Hi , %v. Welcome", name) // %v is format verb. like format string in C#
	return message
}
