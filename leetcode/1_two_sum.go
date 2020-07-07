package main

import "fmt"

func twoSum(nums []int, target int) []int {
	var m = map[int]int{}
	for i, v := range nums {
		find := target - v
		if index, s := m[find]; s {
			return []int{i, index}
		}
		m[v] = i
	}
	return []int{}
}

func main() {
	a := []int{1, 2, 2, 2, 9}
	b := twoSum(a, 10)
	fmt.Println(b)
}
