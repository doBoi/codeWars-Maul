package main

import "fmt"

func Past(h, m, s int) int {
	ms := 0
	if h != 0 {
		ms += (h * 3600000)
	}

	if m != 0 {
		ms += (m * 60000)
	}

	if s != 0 {
		ms += (s * 1000)
	}

	return ms
	return (h*3600000 + m*60000 + s*1000) // opsi one line
}

func main() {
	fmt.Println(Past(0, 1, 1))
	fmt.Println(Past(1, 0, 1))
	fmt.Println(Past(1, 1, 1))

}
