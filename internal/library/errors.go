package library

import "errors"

var (
	NoBookAvailableError         = errors.New("no book available in the library")
	NoBookFoundError             = errors.New("book not found in the library")
	BookAlreadyBorrowedError     = errors.New("book is already borrowed")
	BookShouldNotBeBorrowedError = errors.New("book is not borrowed and cannot be returned")
)
