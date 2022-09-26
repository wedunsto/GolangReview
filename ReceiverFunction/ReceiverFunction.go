//Objective: User a receiver function to print a person's name and age

package main

import "fmt"

// Will be discussed later
type person struct {
	name string
	age  int
}

//Receiver function; Objects of type person have access to this function
func (p person) print() {
	fmt.Printf("%s is %d years old \n", p.name, p.age)
}

func main() {
	//Create an object of type person
	william := person{
		name: "William Dunston",
		age:  30,
	}

	//Object has access to print function
	william.print()
}
