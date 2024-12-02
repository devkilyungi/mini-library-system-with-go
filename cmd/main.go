package main

import (
	"fmt"
	"github.com/devkilyungi/mini-library-system-with-go/internal/handlers"
	"github.com/devkilyungi/mini-library-system-with-go/internal/library"
)

func main() {
	lib := library.NewLibrary()
	lib.AddBooks(
		library.Book{Name: "Book A", Author: "Author A"},
		library.Book{Name: "Book B", Author: "Author B"},
		library.Book{Name: "Book C", Author: "Author C"},
	)

	fmt.Println("Welcome to the Library Management System!")
	fmt.Println("\nChoose an option:")
	fmt.Println("1. View available books")
	fmt.Println("2. Borrow a book")
	fmt.Println("3. Return a book")
	fmt.Println("4. Exit")

	for {
		choice, err := handlers.GetUserChoice()
		if err != nil {
			fmt.Println("Invalid input. Please enter a valid number.")
			continue
		}

		switch choice {
		case 1:
			fmt.Println("Available books:")
			availableBooks := lib.GetAvailableBooks()
			if len(availableBooks) == 0 {
				fmt.Println("- No books available.")
				break
			}

			for _, book := range availableBooks {
				fmt.Printf("- %s by %s\n", book.Name, book.Author)
			}
		case 2:
			title := handlers.GetBookTitle()
			book, err := lib.FindBookByTitle(title)
			if err != nil {
				fmt.Printf("Error: %s\n", err)
				break
			}

			if err := book.Borrow(); err != nil {
				fmt.Printf("Error: %s\n", err)
				break
			}
			fmt.Printf("You have borrowed: %s\n", book.Name)
		case 3:
			title := handlers.GetBookTitle()
			book, err := lib.FindBookByTitle(title)
			if err != nil {
				fmt.Printf("Error: %s\n", err)
				break
			}

			if err := book.Return(); err != nil {
				fmt.Printf("Error: %s\n", err)
				break
			}
			fmt.Printf("You have returned: %s\n", book.Name)
		case 4:
			fmt.Println("Goodbye!")
			return
		default:
			fmt.Println("Invalid option. Please try again.")
		}
	}
}
