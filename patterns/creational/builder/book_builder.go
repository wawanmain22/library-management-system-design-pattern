package builder

import "fmt"

// BookBuilder is used to construct Book objects step by step
type BookBuilder struct {
	book *Book
}

// NewBookBuilder creates a new BookBuilder with default values
func NewBookBuilder() *BookBuilder {
	return &BookBuilder{
		book: &Book{
			ID:        generateID(),
			PageCount: 0,
			Language:  "English",
			Tags:      []string{},
		},
	}
}

// SetTitle sets the book title (required)
func (b *BookBuilder) SetTitle(title string) *BookBuilder {
	b.book.Title = title
	return b
}

// SetAuthor sets the book author (required)
func (b *BookBuilder) SetAuthor(author string) *BookBuilder {
	b.book.Author = author
	return b
}

// SetISBN sets the ISBN of the book
func (b *BookBuilder) SetISBN(isbn string) *BookBuilder {
	b.book.ISBN = isbn
	return b
}

// SetPublisher sets the publisher of the book
func (b *BookBuilder) SetPublisher(publisher string) *BookBuilder {
	b.book.Publisher = publisher
	return b
}

// SetCategory sets the category of the book
func (b *BookBuilder) SetCategory(category string) *BookBuilder {
	b.book.Category = category
	return b
}

// SetPageCount sets the number of pages
func (b *BookBuilder) SetPageCount(count int) *BookBuilder {
	b.book.PageCount = count
	return b
}

// SetLanguage sets the language of the book
func (b *BookBuilder) SetLanguage(language string) *BookBuilder {
	b.book.Language = language
	return b
}

// SetPublished sets the publication date
func (b *BookBuilder) SetPublished(published string) *BookBuilder {
	b.book.Published = published
	return b
}

// AddTag adds a tag to the book
func (b *BookBuilder) AddTag(tag string) *BookBuilder {
	b.book.Tags = append(b.book.Tags, tag)
	return b
}

// Build creates the final Book object
// Returns a new copy of the book to avoid shared reference issues
// Returns an error if required fields are missing
func (b *BookBuilder) Build() (*Book, error) {
	// Validate required fields
	if b.book.Title == "" {
		return nil, fmt.Errorf("title is required")
	}
	if b.book.Author == "" {
		return nil, fmt.Errorf("author is required")
	}

	// Return a copy, not the internal reference
	tags := make([]string, len(b.book.Tags))
	copy(tags, b.book.Tags)

	result := &Book{
		ID:        b.book.ID,
		Title:     b.book.Title,
		Author:    b.book.Author,
		ISBN:      b.book.ISBN,
		Publisher: b.book.Publisher,
		Category:  b.book.Category,
		PageCount: b.book.PageCount,
		Language:  b.book.Language,
		Published: b.book.Published,
		Tags:      tags,
	}

	return result, nil
}

var bookCount int

func init() {
	bookCount = 0
}

// generateID generates a unique ID for the book using counter
func generateID() string {
	bookCount++
	return fmt.Sprintf("BK-%d", bookCount)
}
