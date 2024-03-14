package main

import "fmt"

func GetCount(str string) (count int) {
	vowels := [...]string{
		"a",
		"i",
		"u",
		"e",
		"o",
	}
	for _, st := range str {
		s := string(st)
		for _, vowel := range vowels {
			if s == vowel {
				count += 1
			}
		}
	}
	return count
}
func main() {
	fmt.Println(GetCount("Himada"))
}
