//converts from fahrenheit to celsius

package main

import "fmt"

func main() {
	var inputTemperature float64
	var outputTemperature float64
	fmt.Print("Enter temperature in Fahrenheit: ")
	fmt.Scanf("%f", &inputTemperature)
	outputTemperature = (inputTemperature - 32) * 5 / 9
	fmt.Println("Converted to Celsius: ", outputTemperature)
}
