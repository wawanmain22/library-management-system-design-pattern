package decorator

import "fmt"

// ReferenceOnlyBookDecorator wraps a book that can only be read in the library
// Embeds BaseBookDecorator for default delegation, overrides only changed behavior
type ReferenceOnlyBookDecorator struct {
	BaseBookDecorator
}

// NewReferenceOnlyBookDecorator creates a new reference only book decorator
func NewReferenceOnlyBookDecorator(book BookComponent) *ReferenceOnlyBookDecorator {
	return &ReferenceOnlyBookDecorator{
		BaseBookDecorator: BaseBookDecorator{Component: book},
	}
}

// GetDetails returns the book details with reference only info
func (robd *ReferenceOnlyBookDecorator) GetDetails() string {
	return fmt.Sprintf("%s [Reference Only]", robd.Component.GetDetails())
}

// CanBorrow returns false for reference only books
func (robd *ReferenceOnlyBookDecorator) CanBorrow() bool {
	return false
}

// CanReadInLibrary returns true for reference only books
func (robd *ReferenceOnlyBookDecorator) CanReadInLibrary() bool {
	return true
}

// Borrow is not allowed for reference only books
func (robd *ReferenceOnlyBookDecorator) Borrow() {
	fmt.Println("Cannot borrow a reference only book")
}

// Return is not allowed for reference only books
func (robd *ReferenceOnlyBookDecorator) Return() {
	fmt.Println("Cannot return a reference only book")
}
