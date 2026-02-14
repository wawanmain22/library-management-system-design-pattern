package state

// BookState interface defines methods for all book states
type BookState interface {
	Borrow(book *Book) error
	Return(book *Book) error
	MarkOverdue(book *Book) error
	GetStateName() string
}
