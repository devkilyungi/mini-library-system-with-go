package library

import (
	"errors"
	"testing"
)

func TestLibrary(t *testing.T) {

	t.Run("add books", func(t *testing.T) {
		library := Library{}
		library.AddBooks()
		got := library.GetAvailableBooks()

		if len(got) != 3 {
			t.Errorf("got %d books, want 3", len(got))
		}
	})

	t.Run("borrow book from empty library", func(t *testing.T) {
		bookA := Book{
			Name:       "Book A",
			Author:     "Author A",
			isBorrowed: false,
		}

		library := Library{}
		_, err := library.BorrowBook(bookA)

		if !errors.Is(err, NoBookAvailableError) {
			t.Errorf("got %v, want nil", err)
		}
	})

	t.Run("borrow unavailable book from library", func(t *testing.T) {
		bookD := Book{
			Name:       "Book D",
			Author:     "Author D",
			isBorrowed: false,
		}

		library := Library{}
		library.AddBooks()
		_, err := library.BorrowBook(bookD)

		if !errors.Is(err, NoBookFoundError) {
			t.Errorf("got %v, want nil", err)
		}
	})

	t.Run("borrow book from library with books", func(t *testing.T) {
		bookA := Book{
			Name:       "Book A",
			Author:     "Author A",
			isBorrowed: false,
		}

		library := Library{}
		library.AddBooks()
		status, err := library.BorrowBook(bookA)

		if err != nil {
			t.Fatal("got an error but didn't want one")
		}

		if !status {
			t.Errorf("Book %v not borrowed successfully. %s", bookA, err)
		}
	})

	t.Run("return book to empty library", func(t *testing.T) {
		bookA := Book{
			Name:       "Book A",
			Author:     "Author A",
			isBorrowed: true,
		}

		library := Library{}
		_, err := library.ReturnBook(bookA)

		if !errors.Is(err, NoBookFoundError) {
			t.Errorf("got %v, want nil", err)
		}
	})

	t.Run("return book to wrong library", func(t *testing.T) {
		bookD := Book{
			Name:       "Book D",
			Author:     "Author D",
			isBorrowed: true,
		}

		library := Library{}
		library.AddBooks()
		_, err := library.ReturnBook(bookD)

		if !errors.Is(err, NoBookFoundError) {
			t.Errorf("got %v, want nil", err)
		}
	})

	t.Run("return book that wasn't borrowed to library", func(t *testing.T) {
		bookA := Book{
			Name:       "Book A",
			Author:     "Author A",
			isBorrowed: true,
		}

		library := Library{}
		library.AddBooks()
		_, err := library.ReturnBook(bookA)

		if !errors.Is(err, BookShouldNotBeBorrowedError) {
			t.Errorf("got %v, want nil", err)
		}
	})

	t.Run("borrow and return book", func(t *testing.T) {
		bookA := Book{
			Name:       "Book A",
			Author:     "Author A",
			isBorrowed: true,
		}

		library := Library{}
		library.AddBooks()
		status, err := library.BorrowBook(bookA)
		if err != nil {
			t.Fatal("got an error but didn't want one")
		}

		if !status {
			t.Errorf("Book %v not borrowed successfully. %s", bookA, err)
		}

		returnStatus, returnError := library.ReturnBook(bookA)
		if returnError != nil {
			t.Fatal("got an error but didn't want one")
		}

		if !returnStatus {
			t.Errorf("Book %v not returned successfully. %s", bookA, returnError)
		}
	})
}
