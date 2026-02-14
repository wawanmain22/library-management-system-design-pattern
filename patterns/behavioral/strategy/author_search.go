package strategy

import (
	"strings"
)

// AuthorSearchStrategy searches books by author
type AuthorSearchStrategy struct{}

// NewAuthorSearchStrategy creates a new author search strategy
func NewAuthorSearchStrategy() *AuthorSearchStrategy {
	return &AuthorSearchStrategy{}
}

// Search searches for books by author (case-insensitive, partial match)
func (ass *AuthorSearchStrategy) Search(query string, books []Book) []Book {
	results := make([]Book, 0)
	lowerQuery := strings.ToLower(query)

	for _, book := range books {
		if strings.Contains(strings.ToLower(book.Author), lowerQuery) {
			results = append(results, book)
		}
	}
	return results
}

// GetStrategyName returns the strategy name
func (ass *AuthorSearchStrategy) GetStrategyName() string {
	return "Author Search"
}
