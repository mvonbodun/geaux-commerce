package category

import "errors"

// Category represents a collection of child products
type Category struct {
	ID          int64  `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

var (
	ErrInconsistentIDs = errors.New("inconsistent IDs")
	ErrAlreadyExists   = errors.New("already exists")
	ErrNotFound        = errors.New("not found")
)

// Repository provides access to a category store
type Repository interface {

	// Store saves a given Category, assigning it a new ID.
	Store(c *Category) (id int64, err error)

	// Find retrieves a given Category by its ID.
	Find(id int64) (*Category, error)

	// Close closes the database, freeing up any available resources.
	Close()
}
