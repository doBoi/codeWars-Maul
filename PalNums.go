package main

//LEET CODE
import (
	"fmt"
	"strconv"
)

func IsPalindrome(x int) bool {
	var r string
	st := strconv.Itoa(x)
	for i := len(st) - 1; i >= 0; i-- {
		h := string(st[i])
		r += h
	}
	R, _ := strconv.Atoi(r)
	return R == x
}

func main() {
	fmt.Println(IsPalindrome(563))
	fmt.Println(IsPalindrome(636))
}
