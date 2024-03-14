package main

import "fmt"

func HowMuchILoveYou(i int) string {
	petals := ""
	switch i {
	case 1:
		petals = "I love you"
	case 2:
		petals = "a little"
	case 3:
		petals = "a lot"
	case 4:
		petals = "passionately"
	case 5:
		petals = "madly"
	case 6:
		petals = "not at all"
	default:
		return HowMuchILoveYou(i - 6)
	}
	return petals
}

func main() {
	fmt.Println(HowMuchILoveYou(7))
	fmt.Println(HowMuchILoveYou(4))
	fmt.Println(HowMuchILoveYou(2))
	fmt.Println(HowMuchILoveYou(9))
}
