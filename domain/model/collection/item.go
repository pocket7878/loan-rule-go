package collection

import "fmt"

type Item struct {
	itemNumber ItemNumber
	entry      Entry
}

func NewItem(itemNumber ItemNumber, entry Entry) *Item {
	return &Item{itemNumber: itemNumber, entry: entry}
}

func (i *Item) String() string {
	return fmt.Sprintf("Item{itemNumber=%s, entry=%s}", i.itemNumber, i.entry)
}
