package main

import "fmt"

// helloWorld returns a greeting for the given name.
func helloWorld(name string) string {
	return "Hello, " + name + "!"
}

func main() {
	fmt.Println(helloWorld("World"))
}
