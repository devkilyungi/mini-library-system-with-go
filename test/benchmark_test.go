package test

import (
	"github.com/devkilyungi/mini-library-system-with-go/internal/library"
	"testing"
)

func BenchmarkAddBooks(b *testing.B) {
	for i := 0; i < b.N; i++ {
		lib := library.NewLibrary() // Reset library for each iteration
		lib.AddBooks(library.Book{Name: "Book A", Author: "Author A"})
	}
}

func BenchmarkBorrowBook(b *testing.B) {
	for i := 0; i < b.N; i++ {
		lib := library.NewLibrary()
		lib.AddBooks(library.Book{Name: "Book A", Author: "Author A"})

		book, err := lib.FindBookByTitle("Book A")
		if err != nil {
			b.Fatalf("unexpected error: %s", err)
		}

		err = book.Borrow()
		if err != nil {
			b.Fatalf("unexpected error: %s", err)
		}
	}
}

func BenchmarkBorrowAndReturnBook(b *testing.B) {
	for i := 0; i < b.N; i++ {
		lib := library.NewLibrary()
		lib.AddBooks(library.Book{Name: "Book A", Author: "Author A"})

		book, err := lib.FindBookByTitle("Book A")
		if err != nil {
			b.Fatalf("unexpected error: %s", err)
		}

		borrowErr := book.Borrow()
		if borrowErr != nil {
			b.Fatalf("unexpected error during borrow: %s", borrowErr)
		}

		returnErr := book.Return()
		if returnErr != nil {
			b.Fatalf("unexpected error during return: %s", returnErr)
		}
	}
}
