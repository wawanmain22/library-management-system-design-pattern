package decorator

import "fmt"

// ReservedBookDecorator wraps a book with reservation functionality
// Embeds BaseBookDecorator for default delegation, overrides only changed behavior
type ReservedBookDecorator struct {
	BaseBookDecorator
	reservedBy string
}

// NewReservedBookDecorator creates a new reserved book decorator
func NewReservedBookDecorator(book BookComponent, reservedBy string) *ReservedBookDecorator {
	return &ReservedBookDecorator{
		BaseBookDecorator: BaseBookDecorator{Component: book},
		reservedBy:        reservedBy,
	}
}

// GetDetails returns the book details with reservation info
func (rbd *ReservedBookDecorator) GetDetails() string {
	return fmt.Sprintf("%s [Reserved by: %s]", rbd.Component.GetDetails(), rbd.reservedBy)
}

// CanBorrow returns false because a reserved book cannot be borrowed by others
func (rbd *ReservedBookDecorator) CanBorrow() bool {
	return false
}

// Borrow is not allowed for reserved books
func (rbd *ReservedBookDecorator) Borrow() {
	fmt.Println("Cannot borrow a reserved book")
}

// Return is not allowed for reserved books
func (rbd *ReservedBookDecorator) Return() {
	fmt.Println("Cannot return a reserved book")
}

// GetReservedBy returns who reserved the book
func (rbd *ReservedBookDecorator) GetReservedBy() string {
	return rbd.reservedBy
}
