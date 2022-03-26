package collection

import (
	"github.com/pocket7878/loan-rule-go/domain/model/collection"
	"github.com/pocket7878/loan-rule-go/domain/model/collection/book"
)

type CollectionService struct {
	collectionRepository CollectionRepository
}

func NewCollectionService(collectionRepository CollectionRepository) *CollectionService {
	return &CollectionService{collectionRepository: collectionRepository}
}

func (s *CollectionService) EntryList() (*collection.EntryList, error) {
	return s.collectionRepository.EntryList()
}

func (s *CollectionService) Entry(bookNumber book.BookNumber) (*collection.Entry, error) {
	return s.collectionRepository.Entry(bookNumber)
}
