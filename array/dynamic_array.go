package array

import (
	"go-algorithm/utils"
)

const defaultDynamicCap = 5

type DynamicArray struct {
	item   []interface{}
	cap    int
	length int
}

func NewDynamicArray() DynamicArray {
	return DynamicArray{
		item:   make([]interface{}, defaultDynamicCap),
		cap:    defaultDynamicCap,
		length: 0,
	}
}

func (d *DynamicArray) Item() []interface{} {
	return d.item
}

func (d *DynamicArray) Cap() int {
	return d.cap
}

func (d *DynamicArray) Len() int {
	return d.length
}

func (d *DynamicArray) Push(e interface{}) {
	d.expandCap()
	d.length++
	d.item[d.Len()-1] = e

}

func (d *DynamicArray) Pop() {
	if d.Len() > 0 {
		d.item[d.Len()-1] = nil
		d.reduceCap()
		d.length--
	}
}

func (d *DynamicArray) expandCap() {
	if d.Len() >= d.Cap() {
		newCap := d.Cap() * 2
		newArray := make([]interface{}, newCap)
		for i := 0; i < d.Len(); i++ {
			newArray[i] = d.item[i]
		}
		d.item = newArray
		d.cap = newCap
	}
}

func (d *DynamicArray) reduceCap() {
	if d.Len() <= d.Cap()/4 {
		newCap := d.Cap() / 2
		newArray := make([]interface{}, newCap)
		for i := 0; i < d.Len(); i++ {
			newArray[i] = d.item[i]
		}
		d.item = newArray
		d.cap = newCap
	}
}

func (d *DynamicArray) Index(i int) (interface{}, error) {
	if i >= 0 && i <= d.Len()-1 {
		return d.item[i], nil
	}
	return 0, utils.NewError("error index", nil)
}
