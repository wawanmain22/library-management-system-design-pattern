package strategy

import (
	"strings"
)

// TitleSearchStrategy searches books by title
type TitleSearchStrategy struct{}

// NewTitleSearchStrategy creates a new title search strategy
func NewTitleSearchStrategy() *TitleSearchStrategy {
	return &TitleSearchStrategy{}
}

// Search searches for books by title (case-insensitive, partial match)
func (tss *TitleSearchStrategy) Search(query string, books []Book) []Book {
	results := make([]Book, 0)
	lowerQuery := strings.ToLower(query)

	for _, book := range books {
		if strings.Contains(strings.ToLower(book.Title), lowerQuery) {
			results = append(results, book)
		}
	}
	return results
}

// GetStrategyName returns the strategy name
func (tss *TitleSearchStrategy) GetStrategyName() string {
	return "Title Search"
}
