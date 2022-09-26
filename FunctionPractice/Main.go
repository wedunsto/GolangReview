// Objective: Print "Hello world" using a function
package main

import "fmt"

func printHelloWorld() { //Helper function
	fmt.Println("Hello world")
}
func main() {
	printHelloWorld()
	output := printHelloWorldAgain()
	fmt.Println(output)
}
