package strategy

// SearchStrategy defines the interface for search algorithms
type SearchStrategy interface {
	Search(query string, books []Book) []Book
	GetStrategyName() string
}
