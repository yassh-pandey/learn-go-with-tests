package main

import "fmt"

const englishHelloPrefix string = "Hello, "

//Greet greets the user
func Greet(name string) string {
	if name == "" {
		name = "World"
	}
	return fmt.Sprintf("%s%s!", englishHelloPrefix, name)
}
func main() {
	fmt.Println(Greet("Yash"))
}
