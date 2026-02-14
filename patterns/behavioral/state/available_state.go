package state

import "fmt"

// AvailableState represents when a book is available for borrowing
type AvailableState struct{}

// NewAvailableState creates a new available state
func NewAvailableState() *AvailableState {
	return &AvailableState{}
}

// Borrow allows borrowing book
func (as *AvailableState) Borrow(book *Book) error {
	if book != nil {
		book.SetState(NewBorrowedState())
		fmt.Printf("Book '%s' is now borrowed\n", book.GetTitle())
	}
	return nil
}

// Return is invalid for available books
func (as *AvailableState) Return(book *Book) error {
	return fmt.Errorf("cannot return a book that is already available")
}

// MarkOverdue is invalid for available books
func (as *AvailableState) MarkOverdue(book *Book) error {
	return fmt.Errorf("cannot mark an available book as overdue")
}

// GetStateName returns state name
func (as *AvailableState) GetStateName() string {
	return "Available"
}
