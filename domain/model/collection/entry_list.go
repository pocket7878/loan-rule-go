package collection

import "fmt"

type EntryList struct {
	list []Entry
}

func NewEntryListOf(list []Entry) *EntryList {
	return &EntryList{list: list}
}

func (l *EntryList) List() []Entry {
	return l.list
}

func (l *EntryList) String() string {
	return fmt.Sprintf("EntryList{list=%v}", l.list)
}
