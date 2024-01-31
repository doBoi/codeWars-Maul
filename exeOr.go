package main

import "fmt"

func Xor(a, b bool) bool {
	if a != b {
		return true
	}
	return false
}

func main() {
	fmt.Println(Xor(false, false)) //false
	fmt.Println(Xor(false, true))  //true
	fmt.Println(Xor(true, false))  //true
	fmt.Println(Xor(true, true))   //true
}
