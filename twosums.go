package main

import "fmt"

//LEET CODE

func twoSum(nums []int, target int) []int {
	//versi brute Force
	// var i, j int
	// for i = 0; i <= len(nums)-1; i++ {
	// 	for j = 1 + i; j <= len(nums)-1; j++ {
	// 		if nums[i]+nums[j] == target {
	// 			return []int{i, j}
	// 		}
	// 	}
	// }
	// return []int{}

	//Versi Hash Table
	NM := make(map[int]int)
	for i, v := range nums {
		t := target - v
		if j, ok := NM[t]; ok {
			return []int{j, i}
		}
		NM[v] = i
	}
	return []int{}

}

func main() {
	nums := []int{
		3, 3, 6, 4, 2,
	}
	t := 6
	fmt.Println(twoSum(nums, t))
	// twoSum(nums, t)
}
