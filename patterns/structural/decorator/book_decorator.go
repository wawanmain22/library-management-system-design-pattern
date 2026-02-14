package decorator

// BookComponent interface defines the Component in Decorator pattern
// Both concrete component (Book) and decorators implement this interface
type BookComponent interface {
	GetTitle() string
	GetAuthor() string
	GetISBN() string
	GetDetails() string
	CanBorrow() bool
	CanReadInLibrary() bool
	Borrow()
	Return()
}

// BaseBookDecorator provides default delegation to the wrapped component
// Concrete decorators embed this struct and override only the methods they need
type BaseBookDecorator struct {
	Component BookComponent
}

// GetTitle delegates to the wrapped component
func (d *BaseBookDecorator) GetTitle() string {
	return d.Component.GetTitle()
}

// GetAuthor delegates to the wrapped component
func (d *BaseBookDecorator) GetAuthor() string {
	return d.Component.GetAuthor()
}

// GetISBN delegates to the wrapped component
func (d *BaseBookDecorator) GetISBN() string {
	return d.Component.GetISBN()
}

// GetDetails delegates to the wrapped component
func (d *BaseBookDecorator) GetDetails() string {
	return d.Component.GetDetails()
}

// CanBorrow delegates to the wrapped component
func (d *BaseBookDecorator) CanBorrow() bool {
	return d.Component.CanBorrow()
}

// CanReadInLibrary delegates to the wrapped component
func (d *BaseBookDecorator) CanReadInLibrary() bool {
	return d.Component.CanReadInLibrary()
}

// Borrow delegates to the wrapped component
func (d *BaseBookDecorator) Borrow() {
	d.Component.Borrow()
}

// Return delegates to the wrapped component
func (d *BaseBookDecorator) Return() {
	d.Component.Return()
}
