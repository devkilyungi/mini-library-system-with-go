package library

import "errors"

type Book struct {
	Name       string
	Author     string
	isBorrowed bool
}

func (b *Book) GetBorrowedStatus() bool {
	return b.isBorrowed
}

type Library struct {
	availableBooks []Book
}

func (l *Library) AddBooks() {
	bookA := Book{
		Name:       "Book A",
		Author:     "Author A",
		isBorrowed: false,
	}
	bookB := Book{
		Name:       "Book B",
		Author:     "Author B",
		isBorrowed: false,
	}
	bookC := Book{
		Name:       "Book C",
		Author:     "Author C",
		isBorrowed: false,
	}

	l.availableBooks = append(l.availableBooks, bookA, bookB, bookC)
}

func (l *Library) GetAvailableBooks() []Book {
	return l.availableBooks
}

func (l *Library) BorrowBook(book Book) (bool, error) {
	bookFound := false

	if len(l.availableBooks) == 0 {
		return false, NoBookAvailableError
	}

	for i := 0; i < len(l.availableBooks); i++ {
		if book.Name == l.availableBooks[i].Name {
			if l.availableBooks[i].isBorrowed {
				return false, BookShouldNotBeBorrowedError
			} else {
				l.availableBooks[i].isBorrowed = true
				bookFound = true
				break
			}
		}
	}

	if !bookFound {
		return false, NoBookFoundError
	} else {
		return true, nil
	}
}

func (l *Library) ReturnBook(book Book) (bool, error) {
	bookFound := false

	for i := 0; i < len(l.availableBooks); i++ {
		if book.Name == l.availableBooks[i].Name {
			if l.availableBooks[i].isBorrowed {
				l.availableBooks[i].isBorrowed = false
				bookFound = true
				break
			} else {
				return false, BookShouldNotBeBorrowedError
			}
		}
	}

	if !bookFound {
		return false, NoBookFoundError
	} else {
		return true, nil
	}
}

var NoBookAvailableError = errors.New("no book available")
var NoBookFoundError = errors.New("no book found")
var BookShouldNotBeBorrowedError = errors.New("book shouldn't be borrowed")
