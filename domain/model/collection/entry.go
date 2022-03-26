package collection

import (
	"fmt"

	"github.com/pocket7878/loan-rule-go/domain/model/collection/book"
)

type Entry struct {
	book book.Book
}

func NewEntry(book book.Book) *Entry {
	return &Entry{book: book}
}

func (e *Entry) Description() string {
	return e.book.Description()
}

func (e *Entry) Number() string {
	return e.book.Number()
}

func (e *Entry) String() string {
	return fmt.Sprintf("Entry{book=%s}", e.book)
}
