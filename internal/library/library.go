package library

type Library struct {
	books []Book
}

func NewLibrary() *Library {
	return &Library{books: []Book{}}
}

func (l *Library) AddBooks(books ...Book) {
	l.books = append(l.books, books...)
}

func (l *Library) FindBookByTitle(title string) (*Book, error) {
	for i := range l.books {
		if l.books[i].Name == title {
			return &l.books[i], nil
		}
	}
	return nil, NoBookFoundError
}

func (l *Library) GetAvailableBooks() []Book {
	var available []Book
	for _, book := range l.books {
		if !book.GetBorrowedStatus() {
			available = append(available, book)
		}
	}
	return available
}
