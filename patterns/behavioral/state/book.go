package state

import "fmt"

// Book is the context that maintains current state
type Book struct {
	title string
	isbn  string
	state BookState
}

// NewBook creates a new book with available state
func NewBook(title, isbn string) *Book {
	return &Book{
		title: title,
		isbn:  isbn,
		state: NewAvailableState(),
	}
}

// GetTitle returns book title
func (b *Book) GetTitle() string {
	return b.title
}

// GetISBN returns book ISBN
func (b *Book) GetISBN() string {
	return b.isbn
}

// SetState changes the current state
func (b *Book) SetState(state BookState) {
	b.state = state
}

// GetState returns the current state
func (b *Book) GetState() BookState {
	return b.state
}

// Borrow attempts to borrow the book
func (b *Book) Borrow() error {
	return b.state.Borrow(b)
}

// Return attempts to return the book
func (b *Book) Return() error {
	return b.state.Return(b)
}

// MarkOverdue attempts to mark the book as overdue
func (b *Book) MarkOverdue() error {
	return b.state.MarkOverdue(b)
}

// GetStateName returns the current state name
func (b *Book) GetStateName() string {
	return b.state.GetStateName()
}

// Display displays book information with current state
func (b *Book) Display() {
	fmt.Printf("Book: %s (ISBN: %s)\n", b.title, b.isbn)
	fmt.Printf("  Current State: %s\n", b.GetStateName())
}
