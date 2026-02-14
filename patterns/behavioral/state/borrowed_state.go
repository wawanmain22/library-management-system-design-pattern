package state

import "fmt"

// BorrowedState represents when a book is borrowed
type BorrowedState struct{}

// NewBorrowedState creates a new borrowed state
func NewBorrowedState() *BorrowedState {
	return &BorrowedState{}
}

// Borrow is invalid for borrowed books
func (bs *BorrowedState) Borrow(book *Book) error {
	return fmt.Errorf("cannot borrow a book that is already borrowed")
}

// Return allows returning book
func (bs *BorrowedState) Return(book *Book) error {
	if book != nil {
		book.SetState(NewAvailableState())
		fmt.Printf("Book is now available\n")
	}
	return nil
}

// MarkOverdue marks book as overdue
func (bs *BorrowedState) MarkOverdue(book *Book) error {
	if book != nil {
		book.SetState(NewOverdueState())
		fmt.Printf("Book is now overdue\n")
	}
	return nil
}

// GetStateName returns state name
func (bs *BorrowedState) GetStateName() string {
	return "Borrowed"
}
