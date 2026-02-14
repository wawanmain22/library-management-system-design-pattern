package prototype

import "fmt"

// cloneCounter is a global counter to generate unique IDs for cloned books
var cloneCounter int

// Prototype interface defines the Clone method
type Prototype interface {
	Clone() Prototype
	GetDetails() string
}

// Book represents a book that can be cloned
type Book struct {
	ID        string
	Title     string
	Author    string
	ISBN      string
	Publisher string
	Category  string
	Price     float64
	Stock     int
}

// Clone creates a deep copy of the book with a new unique ID
func (b *Book) Clone() Prototype {
	cloneCounter++
	newBook := &Book{
		ID:        fmt.Sprintf("BK-C%d", cloneCounter),
		Title:     b.Title,
		Author:    b.Author,
		ISBN:      b.ISBN,
		Publisher: b.Publisher,
		Category:  b.Category,
		Price:     b.Price,
		Stock:     b.Stock,
	}
	return newBook
}

// GetDetails returns a string representation of the book
func (b *Book) GetDetails() string {
	return fmt.Sprintf("Book{ID: %s, Title: %s, Author: %s, ISBN: %s, Price: %.2f, Stock: %d}",
		b.ID, b.Title, b.Author, b.ISBN, b.Price, b.Stock)
}

// UpdateISBN updates the ISBN of the book
func (b *Book) UpdateISBN(isbn string) {
	b.ISBN = isbn
}

// UpdatePrice updates the price of the book
func (b *Book) UpdatePrice(price float64) {
	b.Price = price
}

// UpdateStock updates the stock of the book
func (b *Book) UpdateStock(stock int) {
	b.Stock = stock
}
