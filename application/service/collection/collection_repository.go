package collection

import (
	"github.com/pocket7878/loan-rule-go/domain/model/collection"
	"github.com/pocket7878/loan-rule-go/domain/model/collection/book"
)

type CollectionRepository interface {
	EntryList() (*collection.EntryList, error)
	Entry(bookNumber book.BookNumber) (*collection.Entry, error)
}
