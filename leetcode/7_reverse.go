//给出一个 32 位的有符号整数，你需要将这个整数中每位上的数字进行反转。
//
//示例 1:
//	输入: 123
//	输出: 321
//
//示例 2:
//	输入: -123
//	输出: -321
//
//示例 3:
//	输入: 120
//	输出: 21
//
//注意:
//假设我们的环境只能存储得下 32 位的有符号整数，则其数值范围为 [−2^31,  2^(31−1)]。请根据这个假设，如果反转后整数溢出那么就返回 0。

package main

import (
	"fmt"
	"math"
)

// 使用递归也可以，但是leetcode声明全局变量会出错
//var num int
//
//func reverse(x int) int {
//	y := x % 10
//	x = (x - y) / 10
//	if x != 0 {
//		num = (num + y) * 10
//		reverse(x)
//	} else {
//		num += y
//	}
//	if num < math.MinInt32 || num > math.MaxInt32 {
//		return 0
//	}
//	return num
//}

func reverse(x int) int {
	var num int
	for x != 0 {
		// 取余拿到最后一位数字
		y := x % 10
		// 取商拿到还未取余部分的数字
		x /= 10
		// 对位数进行乘10运算
		num = num*10 + y
	}
	if num < math.MinInt32 || num > math.MaxInt32 {
		return 0
	}
	return num
}

func main() {
	a := reverse(1534236469)
	fmt.Println(a)
}
