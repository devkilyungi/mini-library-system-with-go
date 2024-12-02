package main

import (
	"bufio"
	"fmt"
	"github.com/devkilyungi/mini-library-system-with-go/internal/library"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	newLibrary := library.Library{}
	newLibrary.AddBooks()

	fmt.Println("Welcome to the Library Management System!")
	fmt.Println("Choose option:")
	fmt.Println("1. View available books")
	fmt.Println("2. Borrow a book")
	fmt.Println("3. Return a book")
	fmt.Println("4. Exit")

	for {
		fmt.Println("Enter your choice:")
		availableBooks := newLibrary.GetAvailableBooks()
		var borrowableBooks []library.Book

		userOption, _ := reader.ReadString('\n')
		userOption = strings.TrimSpace(userOption)
		choice, err := strconv.Atoi(userOption)
		if err != nil {
			fmt.Println("Invalid choice. Please enter a valid number.")
			return
		}

		switch choice {
		case 1:
			fmt.Println("Available books:")

			for i := 0; i < len(availableBooks); i++ {
				if !availableBooks[i].GetBorrowedStatus() {
					borrowableBooks = append(borrowableBooks, availableBooks[i])
				}
			}

			for i := 0; i < len(borrowableBooks); i++ {
				fmt.Printf("- %s\n", borrowableBooks[i].Name)
			}
		case 2:
			// TODO: Correct borrowing logic
			fmt.Printf("Enter the book title to borrow:")
			title, _ := reader.ReadString('\n')
			title = strings.TrimSpace(title)
			title = strings.ToTitle(title)

			for i := 0; i < len(borrowableBooks); i++ {
				if borrowableBooks[i].Name == title {
					success, err := newLibrary.BorrowBook(borrowableBooks[i])

					if success {
						fmt.Printf("Book %s borrowed!\n", borrowableBooks[i].Name)
					}

					if err != nil {
						fmt.Printf("Error borrowing book, %s\n", err)
					}
				}
			}
		case 3:
			// TODO: Correct returning logic
			fmt.Printf("Enter the book title to return:")
			title, _ := reader.ReadString('\n')
			title = strings.TrimSpace(title)
			title = strings.ToTitle(title)

			for i := 0; i < len(availableBooks); i++ {
				if availableBooks[i].Name == title {
					success, err := newLibrary.ReturnBook(availableBooks[i])

					if success {
						fmt.Printf("Book %s returned!\n", availableBooks[i].Name)
					}

					if err != nil {
						fmt.Printf("Error returning book, %s\n", err)
					}
				}
			}
		case 4:
			fmt.Printf("Goodbye!")
			return
		}
	}
}
