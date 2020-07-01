package array

import "fmt"

func TestDynamicArray() {
	a := NewDynamicArray()
	b, e := a.Index(199)
	fmt.Println(a.Item(), b, e)
}

func TestSparseArray() {
	v := [][]interface{}{{1, 0, 0, 0}, {0, 0, 0, 0}, {0, 0, 0, 0}, {0, 0, 0, 0}}
	a := NewSparseArray(v)
	b := a.TransToArray()
	fmt.Println(*a, b)
}
