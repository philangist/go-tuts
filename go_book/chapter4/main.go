package main

import "fmt"

const (
	firstName = "John"
	lastName  = "Smith"
)

var x string = "Hello World!"

func f() {
	fmt.Println(x)
	fmt.Println("Full name: ", firstName+" "+lastName)
}

func main() {
	fmt.Println(x)
	var y string
	y = "I like Ice Cream! :)"
	fmt.Println(y)
	var z string = "first"
	fmt.Println(z)
	z = "second"
	fmt.Println(z)
	x = "first"
	fmt.Println(x)
	x += " second"
	fmt.Println(x)
	z = "first second"
	fmt.Println(x == z)
	dogsName := "Max"
	fmt.Println("My dog's name is ", dogsName)
	f()
}
