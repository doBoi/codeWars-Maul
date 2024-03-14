package main

import (
	"fmt"
)

func PowersOfTwo(n int) []uint64 {
	if n == 0 {
		return []uint64{1}
	} else {
		results := make([]uint64, n+1, n+1)
		results[0] = 1

		for i := 1; i <= n; i++ {
			results[i] = uint64(2 * results[i-1])
		}
		return results
	}
}

func main() {
	fmt.Println(PowersOfTwo(0))
	fmt.Println(PowersOfTwo(1))
	fmt.Println(PowersOfTwo(2))
	fmt.Println(PowersOfTwo(4))
	fmt.Println(PowersOfTwo(5))
	fmt.Println(PowersOfTwo(6))
}
