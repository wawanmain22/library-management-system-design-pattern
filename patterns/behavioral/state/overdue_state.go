package state

import "fmt"

// OverdueState represents when a book is overdue
type OverdueState struct{}

// NewOverdueState creates a new overdue state
func NewOverdueState() *OverdueState {
	return &OverdueState{}
}

// Borrow is invalid for overdue books
func (os *OverdueState) Borrow(book *Book) error {
	return fmt.Errorf("cannot borrow an overdue book")
}

// Return allows returning book
func (os *OverdueState) Return(book *Book) error {
	if book != nil {
		book.SetState(NewAvailableState())
		fmt.Printf("Book is now available (returned from overdue)\n")
	}
	return nil
}

// MarkOverdue is invalid for already overdue books
func (os *OverdueState) MarkOverdue(book *Book) error {
	return fmt.Errorf("book is already overdue")
}

// GetStateName returns state name
func (os *OverdueState) GetStateName() string {
	return "Overdue"
}
