//给定一个整数数组 nums 和一个目标值 target，请你在该数组中找出和为目标值的那两个整数，并返回他们的数组下标。
//
//你可以假设每种输入只会对应一个答案。但是，数组中同一个元素不能使用两遍。
//
// 
//
//示例:
//	给定 nums = [2, 7, 11, 15], target = 9
//
//	因为 nums[0] + nums[1] = 2 + 7 = 9
//	所以返回 [0, 1]

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
