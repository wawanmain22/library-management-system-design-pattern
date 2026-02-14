package decorator

import "fmt"

// Book represents a base book
type Book struct {
	Title     string
	Author    string
	ISBN      string
	Publisher string
	Copies    int
}

// GetTitle returns the book title
func (b *Book) GetTitle() string {
	return b.Title
}

// GetAuthor returns the book author
func (b *Book) GetAuthor() string {
	return b.Author
}

// GetISBN returns the book ISBN
func (b *Book) GetISBN() string {
	return b.ISBN
}

// GetDetails returns book details
func (b *Book) GetDetails() string {
	return fmt.Sprintf("Book{Title: %s, Author: %s, ISBN: %s, Copies: %d}",
		b.Title, b.Author, b.ISBN, b.Copies)
}

// CanBorrow checks if the book can be borrowed
func (b *Book) CanBorrow() bool {
	return b.Copies > 0
}

// CanReadInLibrary checks if the book can be read in the library
func (b *Book) CanReadInLibrary() bool {
	return true
}

// Borrow reduces the number of available copies
func (b *Book) Borrow() {
	if b.Copies > 0 {
		b.Copies--
	}
}

// Return increases the number of available copies
func (b *Book) Return() {
	b.Copies++
}
