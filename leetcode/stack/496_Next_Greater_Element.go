//给定两个 没有重复元素 的数组 nums1 和 nums2 ，其中nums1 是 nums2 的子集。找到 nums1 中每个元素在 nums2 中的下一个比其大的值。
//
//nums1 中数字 x 的下一个更大元素是指 x 在 nums2 中对应位置的右边的第一个比 x 大的元素。如果不存在，对应位置输出 -1 。

package main

import "fmt"

//暴力破解
//func findIndex(slice []int, x int) int {
//	for i, v := range slice {
//		if v == x {
//			return i
//		}
//	}
//	return -1
//}
//
//func nextGreaterElement(nums1 []int, nums2 []int) []int {
//	var (
//		result []int
//		loop   bool
//	)
//	for _, v := range nums1 {
//		i := findIndex(nums2, v)
//		for ; i < len(nums2); i++ {
//			if nums2[i] > v {
//				result = append(result, nums2[i])
//				loop = true
//				break
//			} else {
//				loop = false
//			}
//		}
//		if !loop {
//			result = append(result, -1)
//		}
//	}
//	return result
//}

func nextGreaterElement(nums1 []int, nums2 []int) []int {

}

func main() {
	r := nextGreaterElement([]int{4, 1, 2}, []int{1, 3, 4, 2})
	fmt.Println(r)
}
