package main

import (
	"fmt"

	"library-management-system/patterns/behavioral/state"
	"library-management-system/patterns/behavioral/strategy"
	"library-management-system/patterns/creational/builder"
	"library-management-system/patterns/creational/prototype"
	"library-management-system/patterns/structural/decorator"
)

func main() {
	fmt.Println("=== LIBRARY MANAGEMENT SYSTEM - DESIGN PATTERN DEMO ===")
	fmt.Println()

	demoBuilderPattern()
	fmt.Println()

	demoPrototypePattern()
	fmt.Println()

	demoDecoratorPattern()
	fmt.Println()

	demoStatePattern()
	fmt.Println()

	demoStrategyPattern()
}

// BUILDER PATTERN DEMO
func demoBuilderPattern() {
	fmt.Println("=== BUILDER PATTERN ===")
	fmt.Println("Creating books using Builder pattern with method chaining")

	book1, _ := builder.NewBookBuilder().
		SetTitle("1984").
		SetAuthor("George Orwell").
		SetISBN("9780451524935").
		SetPublisher("Signet Classic").
		SetCategory("Fiction").
		SetPageCount(328).
		SetLanguage("English").
		SetPublished("1949").
		AddTag("Dystopian").
		AddTag("Classic").
		Build()

	fmt.Printf("Created: %s\n", book1)

	book2, _ := builder.NewBookBuilder().
		SetTitle("Clean Code").
		SetAuthor("Robert C. Martin").
		SetISBN("9780132350884").
		SetCategory("Technology").
		SetPageCount(464).
		Build()

	fmt.Printf("Created: %s\n", book2)

	book3, _ := builder.NewBookBuilder().
		SetTitle("The Great Gatsby").
		SetAuthor("F. Scott Fitzgerald").
		Build()

	fmt.Printf("Created: %s\n", book3)

	bookWithoutTitle, err := builder.NewBookBuilder().
		SetAuthor("Test Author").
		Build()
	if err != nil {
		fmt.Printf("Validation Error: %s\n", err)
	}
	fmt.Printf("Book without title: %v\n", bookWithoutTitle)
}

// PROTOTYPE PATTERN DEMO
func demoPrototypePattern() {
	fmt.Println("=== PROTOTYPE PATTERN ===")
	fmt.Println("Creating book copies using Prototype pattern")

	manager := prototype.NewPrototypeManager()

	originalBook := &prototype.Book{
		ID:        "BK-001",
		Title:     "Design Patterns",
		Author:    "Erich Gamma",
		ISBN:      "9780201633610",
		Publisher: "Addison-Wesley",
		Category:  "Technology",
		Price:     45.99,
		Stock:     10,
	}

	manager.RegisterPrototype("design-patterns", originalBook)
	fmt.Printf("Original book: %s\n", originalBook.GetDetails())

	clone1, _ := manager.GetPrototype("design-patterns")
	clonedBook1 := clone1.(*prototype.Book)
	clonedBook1.UpdateISBN("9780201633611")
	clonedBook1.UpdatePrice(50.99)
	fmt.Printf("Cloned book 1: %s\n", clonedBook1.GetDetails())

	clone2, _ := manager.GetPrototype("design-patterns")
	clonedBook2 := clone2.(*prototype.Book)
	clonedBook2.UpdateStock(5)
	fmt.Printf("Cloned book 2: %s\n", clonedBook2.GetDetails())

	fmt.Printf("Original book after clones: %s\n", originalBook.GetDetails())
}

// DECORATOR PATTERN DEMO
func demoDecoratorPattern() {
	fmt.Println("=== DECORATOR PATTERN ===")
	fmt.Println("Adding features to books using Decorator pattern")

	baseBook := &decorator.Book{
		Title:     "The Catcher in the Rye",
		Author:    "J.D. Salinger",
		ISBN:      "9780316769488",
		Publisher: "Little, Brown and Company",
		Copies:    3,
	}

	fmt.Printf("Base book: %s\n", baseBook.GetDetails())
	fmt.Printf("Can borrow: %t\n", baseBook.CanBorrow())
	fmt.Printf("Can read in library: %t\n", baseBook.CanReadInLibrary())

	reservedBook := decorator.NewReservedBookDecorator(baseBook, "student-123")
	fmt.Printf("\nReserved book: %s\n", reservedBook.GetDetails())
	fmt.Printf("Can borrow: %t\n", reservedBook.CanBorrow())
	fmt.Printf("Can read in library: %t\n", reservedBook.CanReadInLibrary())

	referenceBook := decorator.NewReferenceOnlyBookDecorator(baseBook)
	fmt.Printf("\nReference only book: %s\n", referenceBook.GetDetails())
	fmt.Printf("Can borrow: %t\n", referenceBook.CanBorrow())
	fmt.Printf("Can read in library: %t\n", referenceBook.CanReadInLibrary())
}

// STATE PATTERN DEMO
func demoStatePattern() {
	fmt.Println("=== STATE PATTERN ===")
	fmt.Println("Managing book states using State pattern")

	book := state.NewBook("The Alchemist", "9780062315007")
	book.Display()

	err := book.Borrow()
	if err != nil {
		fmt.Printf("Error: %s\n", err)
	}
	book.Display()

	err = book.Borrow()
	if err != nil {
		fmt.Printf("Error: %s\n", err)
	}

	err = book.MarkOverdue()
	if err != nil {
		fmt.Printf("Error: %s\n", err)
	}
	book.Display()

	err = book.Borrow()
	if err != nil {
		fmt.Printf("Error: %s\n", err)
	}

	err = book.Return()
	if err != nil {
		fmt.Printf("Error: %s\n", err)
	}
	book.Display()

	err = book.Return()
	if err != nil {
		fmt.Printf("Error: %s\n", err)
	}
}

// STRATEGY PATTERN DEMO
func demoStrategyPattern() {
	fmt.Println("=== STRATEGY PATTERN ===")
	fmt.Println("Searching books using different Strategy patterns")

	books := []strategy.Book{
		{Title: "Clean Code", Author: "Robert C. Martin", ISBN: "9780132350884", Category: "Technology"},
		{Title: "Refactoring", Author: "Martin Fowler", ISBN: "9780201485677", Category: "Technology"},
		{Title: "The Pragmatic Programmer", Author: "Andrew Hunt", ISBN: "9780135957059", Category: "Technology"},
		{Title: "Design Patterns", Author: "Erich Gamma", ISBN: "9780201633610", Category: "Technology"},
		{Title: "Clean Architecture", Author: "Robert C. Martin", ISBN: "9780134494166", Category: "Technology"},
	}

	catalog := strategy.NewCatalog(books)
	catalog.DisplayCatalog()

	catalog.SetStrategy(strategy.NewTitleSearchStrategy())
	results := catalog.Find("Clean")
	catalog.DisplayResults(results, "Clean")

	catalog.SetStrategy(strategy.NewAuthorSearchStrategy())
	results = catalog.Find("Robert")
	catalog.DisplayResults(results, "Robert")

	catalog.SetStrategy(strategy.NewTitleSearchStrategy())
	results = catalog.Find("Design")
	catalog.DisplayResults(results, "Design")
}
