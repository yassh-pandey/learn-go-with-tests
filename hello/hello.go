package main

import (
	"fmt"
)

// Greeting prefix constants
const englishHelloPrefix = "Hello, "
const spanishHelloPrefix = "Hola, "
const frenchHelloPrefix = "Bonjour, "

// Language constants
const spanish = "spanish"
const english = "english"
const french = "french"

// Returns greeting prefix based on the language
func greetingPrefix(language string) (prefix string) {
	switch language {
	case spanish:
		prefix = spanishHelloPrefix
	case english:
		prefix = englishHelloPrefix
	case french:
		prefix = frenchHelloPrefix
	default:
		prefix = englishHelloPrefix
	}
	return
}

//Greet greets the user
func Greet(name, language string) string {
	if name == "" {
		name = "World"
	}
	return fmt.Sprintf("%s%s!", greetingPrefix(language), name)
}

func main() {
	fmt.Println(Greet("Yash", "english"))
}
