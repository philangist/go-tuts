package main

import "fmt"

const feetToMeter float64 = 0.3048

func main() {
	var inputHeight, outputHeight float64
	fmt.Print("Enter your height in feet: ")
	fmt.Scanf("%f", &inputHeight)
	outputHeight = inputHeight * feetToMeter
	fmt.Println("Your height has been converted to meters!")
	fmt.Println("Your height in meters is: ", outputHeight)
}
