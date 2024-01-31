package main

import "fmt"

func OtherAngle(a int, b int) int {

	return 180 - (a + b)
}

func main() {
	fmt.Println(OtherAngle(30, 60))
	fmt.Println(OtherAngle(60, 60))
	fmt.Println(OtherAngle(43, 78))
	fmt.Println(OtherAngle(10, 20))

}
