package strategy

// Book represents a book for searching
type Book struct {
	Title    string
	Author   string
	ISBN     string
	Category string
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

// GetCategory returns the book category
func (b *Book) GetCategory() string {
	return b.Category
}
