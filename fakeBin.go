package main

import "fmt"

func FakeBin(x string) string {
	bin := ""
	// for i := 0; i < len(x); i++ {
	// 	if x[i] < 53 { // solosi 1 cheating ascii
	// 		bin += "0"
	// 	} else {
	// 		bin += "1"
	// 	}
	// }

	for _, num := range x {
		if num < '5' {
			bin += "0"
		} else {
			bin += "1"
		}
	}
	return bin
}

func main() {
	fmt.Println(FakeBin("45385593107843568"))
}
