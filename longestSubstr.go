package main

import (
	"fmt"
)

func substr(str string) int {
	var max int
	sm := make(map[string]int) // membuat hashmap
	start := 0                 //initial start
	for end := 0; end < len(str); end++ {
		sm[string(str[end])]++ //menjadikan string index ke end sebuah map dan memasukan value +1

		for sm[string(str[end])] > 1 { //akan terus loop sampai value end adalah 1
			sm[string(str[start])]-- //mengurangi value key start, dgn harap key start ==  key end. jika tidak, akan terjadi infiniteLoop sampai key start ==  key end.
			start++                  // memindahkkan titik start
		}
		if end-start+1 > max {
			max = end - start + 1 //  update max
		}
	}

	return max
}

func main() {
	fmt.Println(substr("abcabcbb"))
	fmt.Println(substr("bbbbbbb"))
	fmt.Println(substr("pwwkew"))
	fmt.Println(substr("dvdf"))
}
