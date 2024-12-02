package test

import (
	"errors"
	"github.com/devkilyungi/mini-library-system-with-go/internal/library"
	"testing"
)

func TestFindBookByTitle(t *testing.T) {
	lib := library.NewLibrary()
	lib.AddBooks(library.Book{Name: "Book A", Author: "Author A"})

	t.Run("find existing book", func(t *testing.T) {
		book, err := lib.FindBookByTitle("Book A")
		if err != nil {
			t.Fatalf("unexpected error: %s", err)
		}
		if book.Name != "Book A" {
			t.Errorf("got book %s, want Book A", book.Name)
		}
	})

	t.Run("find non-existent book", func(t *testing.T) {
		book, err := lib.FindBookByTitle("Book Z")
		if err == nil {
			t.Fatalf("expected an error but got none")
		}
		if !errors.Is(err, library.NoBookFoundError) {
			t.Errorf("got error %s, want NoBookFoundError", err)
		}
		if book != nil {
			t.Errorf("expected nil book, but got %v", book)
		}
	})
}

func TestBorrowBook(t *testing.T) {
	lib := library.NewLibrary()
	lib.AddBooks(library.Book{Name: "Book A", Author: "Author A"})

	t.Run("borrow existing book", func(t *testing.T) {
		book, err := lib.FindBookByTitle("Book A")
		if err != nil {
			t.Fatalf("unexpected error: %s", err)
		}

		if err := book.Borrow(); err != nil {
			t.Fatalf("unexpected error: %s", err)
		}
		if !book.GetBorrowedStatus() {
			t.Errorf("expected book to be borrowed, but it was not")
		}
	})

	t.Run("borrow non-existent book", func(t *testing.T) {
		book, err := lib.FindBookByTitle("Book Z")
		if err == nil {
			t.Fatalf("expected an error but got none")
		}
		if !errors.Is(err, library.NoBookFoundError) {
			t.Errorf("got error %s, want NoBookFoundError", err)
		}
		if book != nil {
			t.Errorf("expected nil book, but got %v", book)
		}
	})
}

func TestReturnBook(t *testing.T) {
	lib := library.NewLibrary()
	lib.AddBooks(library.Book{Name: "Book A", Author: "Author A"})

	t.Run("return borrowed book", func(t *testing.T) {
		book, err := lib.FindBookByTitle("Book A")
		if err != nil {
			t.Fatalf("unexpected error: %s", err)
		}

		if err := book.Borrow(); err != nil {
			t.Fatalf("unexpected error: %s", err)
		}

		if err := book.Return(); err != nil {
			t.Fatalf("unexpected error: %s", err)
		}
		if book.GetBorrowedStatus() {
			t.Errorf("expected book to be returned, but it was not")
		}
	})

	t.Run("return non-existent book", func(t *testing.T) {
		book, err := lib.FindBookByTitle("Book Z")
		if err == nil {
			t.Fatalf("expected an error but got none")
		}
		if !errors.Is(err, library.NoBookFoundError) {
			t.Errorf("got error %s, want NoBookFoundError", err)
		}
		if book != nil {
			t.Errorf("expected nil book, but got %v", book)
		}
	})
}
