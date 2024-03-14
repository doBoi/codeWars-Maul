package main

import "fmt"

func CountSheeps(numbers []bool) int {
	sheeps := 0
	if numbers != nil {
		for _, number := range numbers {
			if number == true {
				sheeps += 1
			}
		}
	}
	return sheeps // your code here
}

func main() {
	arr := []bool{
		true, true, true, false,
		true, true, true, true,
		true, false, true, false,
		true, false, false, true,
		true, true, true, true,
		false, false, true, true,
	}
	fmt.Println(CountSheeps(arr))

}
