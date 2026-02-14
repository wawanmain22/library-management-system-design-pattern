package builder

import "fmt"

// Book represents a book in the library
type Book struct {
	ID        string
	Title     string
	Author    string
	ISBN      string
	Publisher string
	Category  string
	PageCount int
	Language  string
	Published string
	Tags      []string
}

// String returns a string representation of the book
func (b *Book) String() string {
	return fmt.Sprintf("Book{ID: %s, Title: %s, Author: %s, ISBN: %s}", b.ID, b.Title, b.Author, b.ISBN)
}
