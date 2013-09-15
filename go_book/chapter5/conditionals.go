package main

import "fmt"

func main() {
	for j := 1; j <= 10; j++ {
		if j%2 == 0 {
			fmt.Println(j, "Even")
		} else {
			fmt.Println(j, "Odd")
		}
	}
	i := 0
	for i <= 6 {
		switch i {
		case 0:
			fmt.Println("Zero")
		case 1:
			fmt.Println("One")
		case 2:
			fmt.Println("Two")
		case 3:
			fmt.Println("Three")
		case 4:
			fmt.Println("Four")
		case 5:
			fmt.Println("Five")
		default:
			fmt.Println("Unknown Number")
		}
		i++
	}
}
