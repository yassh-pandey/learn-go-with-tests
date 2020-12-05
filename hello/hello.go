package main

import "fmt"

//Greet greets the user
func Greet(name string) string {
	return fmt.Sprintf("Hello, %s!", name)
}
func main() {
	fmt.Println(Greet("Yash"))
}
