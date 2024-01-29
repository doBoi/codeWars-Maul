package main

import "fmt"

func Opposite(value int) int {
	return -value
}

func main() {
	fmt.Println(Opposite(1))
	fmt.Println(Opposite(-1))
}
