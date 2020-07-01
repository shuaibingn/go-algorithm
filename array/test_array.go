package array

import "fmt"

func TestArray() {
	a := NewDynamicArray()
	b, e := a.Index(199)
	fmt.Println(a.Item(), b, e)
}
