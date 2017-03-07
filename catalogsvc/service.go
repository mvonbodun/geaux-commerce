package catalogsvc

import (
	"context"
	"errors"
	"log"
	"strconv"

	"github.com/mvonbodun/geaux-commerce/catalogsvc/category"
)

// ErrInvalidArgument holds the error message
var ErrInvalidArgument = errors.New("invalid argument")

// Service provides the API to the catalogsvc
type Service interface {
	PostCategory(ctx context.Context, c category.Category) (category.Category, error)
	GetCategory(ctx context.Context, id string) (category.Category, error)
}

type service struct {
	categories category.Repository
}

// NewService creates a catalogsvc with necessary dependencies.
func NewService(categories category.Repository) Service {
	return &service{
		categories: categories,
	}
}

// PostCategory creates a new category
func (s *service) PostCategory(ctx context.Context, c category.Category) (category.Category, error) {
	if c.Name == "" {
		return c, errors.New("Name cannot be blank")
	}

	id, err := s.categories.Store(&c)
	if err != nil {
		return c, err
	}

	c.ID = id
	return c, nil

}

// GetCategory gets a new Category
func (s *service) GetCategory(ctx context.Context, id string) (category.Category, error) {
	if id == "" {
		return category.Category{}, errors.New("id cannot be blank")
	}

	i, err := strconv.ParseInt(id, 10, 64)
	log.Printf("id in int64 form is: %v", i)
	if err != nil {
		return category.Category{}, err
	}

	c, err := s.categories.Find(i)
	if err != nil {
		log.Printf("Category was not found. err: %v", err)
		return category.Category{}, err
	}

	cat := *c

	return cat, nil

}
