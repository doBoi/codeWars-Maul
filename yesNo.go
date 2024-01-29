package main

import "fmt"

func BoolToWord(word bool) string {
	// if word {
	// 	return "Yes"
	// } else {
	// 	return "No"
	// }
	switch word {
	case true:
		return "Yes"
	default:
		return "No"
	}

}
func main() {
	fmt.Println(BoolToWord(true))
	fmt.Println(BoolToWord(false))
}
