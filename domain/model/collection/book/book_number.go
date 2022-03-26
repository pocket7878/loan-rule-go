package book

type BookNumber struct {
	value string
}

func NewBookNumber(value string) *BookNumber {
	return &BookNumber{value: value}
}

func (b *BookNumber) String() string {
	return b.value
}
