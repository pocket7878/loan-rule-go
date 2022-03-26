package book

import "fmt"

type Book struct {
	number BookNumber
	title  string
	author string
}

func NewBook(number BookNumber, title, author string) *Book {
	return &Book{
		number: number,
		title:  title,
	}
}

func (b *Book) Description() string {
	return fmt.Sprintf("%s (%s)", b.title, b.author)
}

func (b *Book) Number() string {
	return b.number.String()
}

func (b *Book) String() string {
	return fmt.Sprintf("Book{number=%s, title=%s, author=%s}", b.number, b.title, b.author)
}
