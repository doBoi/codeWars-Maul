package main

import "fmt"

func Feast(beast string, dish string) bool {
	return beast[0] == dish[0] && dish[len(dish)-1] == beast[len(beast)-1]
}
func main() {
	fmt.Println(Feast("great blue heron", "garlic naan"))
	fmt.Println(Feast("chickadee", "chocolate cake"))
	fmt.Println(Feast("brown bear", "bear claw"))
}
