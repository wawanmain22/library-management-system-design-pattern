package strategy

import "fmt"

// Catalog uses search strategy to find books
type Catalog struct {
	books    []Book
	strategy SearchStrategy
}

// NewCatalog creates a new catalog with default TitleSearchStrategy
func NewCatalog(books []Book) *Catalog {
	return &Catalog{
		books:    books,
		strategy: NewTitleSearchStrategy(),
	}
}

// SetStrategy changes the search strategy
func (c *Catalog) SetStrategy(strategy SearchStrategy) {
	c.strategy = strategy
}

// Find searches for books using the current strategy
func (c *Catalog) Find(query string) []Book {
	return c.strategy.Search(query, c.books)
}

// AddBook adds a book to the catalog
func (c *Catalog) AddBook(book Book) {
	c.books = append(c.books, book)
}

// GetAllBooks returns all books in the catalog
func (c *Catalog) GetAllBooks() []Book {
	return c.books
}

// GetBookCount returns the total number of books
func (c *Catalog) GetBookCount() int {
	return len(c.books)
}

// DisplayCatalog displays all books in the catalog
func (c *Catalog) DisplayCatalog() {
	fmt.Printf("Catalog contains %d books:\n", len(c.books))
	for i, book := range c.books {
		fmt.Printf("  %d. %s by %s (ISBN: %s)\n", i+1, book.Title, book.Author, book.ISBN)
	}
}

// DisplayResults displays search results
func (c *Catalog) DisplayResults(results []Book, query string) {
	if c.strategy != nil {
		fmt.Printf("\nSearch Results (%s) for '%s':\n", c.strategy.GetStrategyName(), query)
	} else {
		fmt.Printf("\nSearch Results for '%s':\n", query)
	}

	if len(results) == 0 {
		fmt.Println("  No results found")
		return
	}

	for i, book := range results {
		fmt.Printf("  %d. %s by %s (ISBN: %s)\n", i+1, book.Title, book.Author, book.ISBN)
	}
	fmt.Printf("  Found %d result(s)\n", len(results))
}
