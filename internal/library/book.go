package library

type Book struct {
	Name       string
	Author     string
	isBorrowed bool
}

func (b *Book) Borrow() error {
	if b.isBorrowed {
		return BookAlreadyBorrowedError
	}

	b.isBorrowed = true
	return nil
}

func (b *Book) Return() error {
	if !b.isBorrowed {
		return BookShouldNotBeBorrowedError
	}

	b.isBorrowed = false
	return nil
}

func (b *Book) GetBorrowedStatus() bool {
	return b.isBorrowed
}
