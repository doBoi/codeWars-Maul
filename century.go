package main

import "fmt"

func century(year int) int {
	if year%100 == 0 {
		return year / 100
	} else {
		return year/100 + 1
	}
}

func main() {
	fmt.Println(century(99))
	fmt.Println(century(2000))
	fmt.Println(century(1705))
}
