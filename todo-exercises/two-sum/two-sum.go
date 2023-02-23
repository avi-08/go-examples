package main

import "fmt"

func twoSum(arr *[]int, target int) []int {
	seen := make(map[int]int)
	for i, v := range *arr {
		if j, ok := seen[target-v]; ok {
			return []int{j, i}
		}
		seen[v] = i
	}
	return []int{}
}

func main() {
	testArr := []int{100, 3, 5, -1, 2, 0, 10, 39, 4}
	testVals := []int{110, 90, 8, 7, 98, 37}
	for _, v := range testVals {
		fmt.Println("searched", testArr, "for", v, "Indices whose values sum upto target", v, "=>", twoSum(&testArr, v))
	}
}
