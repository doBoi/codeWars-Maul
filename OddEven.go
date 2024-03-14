package main

import "fmt"

func EvenOrOdd(number int) string {
	if number%2 == 0 {
		return "Even"
	}
	return "Odd" // your code here
}

func main() {
	fmt.Println(EvenOrOdd(5))
	fmt.Println(EvenOrOdd(6))
	fmt.Println(EvenOrOdd(90))
	fmt.Println(EvenOrOdd(31))
}
