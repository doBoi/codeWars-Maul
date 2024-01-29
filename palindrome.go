package main

import (
	"fmt"
	"strings"
)

func IsPalindrome(str string) bool {
	result := ""
	strLow := strings.ToLower(str)
	for i := len(strLow) - 1; i >= 0; i-- {
		result += string(strLow[i])
	}
	if strLow == result {
		return true
	}
	return false
}
func main() {
	fmt.Println(IsPalindrome("macam"))
	fmt.Println(IsPalindrome("Aa"))
}
