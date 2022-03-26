package collection

import "fmt"

type ItemNumber struct {
	value int
}

func NewItemNumber(value int) *ItemNumber {
	return &ItemNumber{value: value}
}

func (i *ItemNumber) String() string {
	return fmt.Sprintf("ItemNumber{value=%d}", i.value)
}
