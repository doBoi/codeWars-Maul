package main

// To convert a Roman numeral string to an integer, you can use the following approach:

// 1.Initialize a variable result to store the final integer value.
// 2.Create a map that associates each Roman numeral symbol with its corresponding value.
// 3.Iterate through the Roman numeral string from left to right.
// 4.If the current Roman numeral symbol is less than the next symbol, subtract its value from the result.
// 5.Otherwise, add its value to the result.
// 6.Return the final value stored in result.

func roman(str string) int {
	var result int
	rm := map[string]int{
		"I": 1,
		"V": 5,
		"X": 10,
		"L": 50,
		"C": 100,
		"D": 500,
		"M": 1000,
	}
	for i, s := range str {
		if i == len(str)-1 {
			result += rm[string(s)]
		} else if rm[string(s)] < rm[string(str[i+1])] {
			result = result - rm[string(s)]
		} else {
			result += rm[string(s)]
		}

	}
	return result
}

func main() {
	roman("XIV")
	roman("III")
	roman("LVIII")
}
