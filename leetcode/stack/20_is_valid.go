package main

import "fmt"

func isValid(s string) bool {
	if l := len(s); l%2 != 0 {
		// 如果不是2的倍数，直接返回false
		return false
	}
	// 创建一个int32的map和切片
	symbolMap := map[int32]int32{'{': '}', '[': ']', '(': ')'}
	stackSlice := make([]int32, 0, len(s)/2)
	for _, v := range s {
		if _, ok := symbolMap[v]; ok {
			// 当s中的字符可以在map中匹配到时，将其添加到切片末尾
			stackSlice = append(stackSlice, v)
		} else {
			if len(stackSlice) > 0 && symbolMap[stackSlice[len(stackSlice)-1]] == v {
				// 匹配不到时，查找当前字符是否与切片中已存在的字符匹配，如果匹配，移除最后一位匹配到的字符
				stackSlice = stackSlice[:len(stackSlice)-1]
			} else {
				// 不匹配则直接返回false
				return false
			}
		}
	}
	if len(stackSlice) == 0 {
		// 如果最后切片长度为0，则说明，全部符合规则，返回true
		return true
	}
	return false
}

func main() {
	res := isValid(`{[]}`)
	fmt.Println(res)
}
