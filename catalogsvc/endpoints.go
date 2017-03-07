package catalogsvc

import (
	"context"

	"github.com/go-kit/kit/endpoint"
	"github.com/mvonbodun/geaux-commerce/catalogsvc/category"
)

// Endpoints collects all of the endpoints for the category service
type Endpoints struct {
	PostCategoryEndpoint endpoint.Endpoint
	GetCategoryEndpoint  endpoint.Endpoint
}

// MakeServerEndpoints returns an Endpoints struct where each endpoint invokes
// the corresponding method on the provided service. Useful in a catalogsvc
// server.
func MakeServerEndpoints(s Service) Endpoints {
	return Endpoints{
		PostCategoryEndpoint: MakePostCategoryEnpoint(s),
		GetCategoryEndpoint:  MakeGetCategoryEndpoint(s),
	}
}

// MakePostCategoryEnpoint returns an endpoint via the passed service.
func MakePostCategoryEnpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(postCategoryRequest)
		c, e := s.PostCategory(ctx, req.Category)
		return postCategoryResponse{Category: c, Err: e}, nil
	}
}

// MakeGetCategoryEndpoint returns an endpoint via the passed service.
func MakeGetCategoryEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(getCategoryRequest)
		c, e := s.GetCategory(ctx, req.ID)
		return getCategoryResponse{Category: c, Err: e}, nil
	}
}

type postCategoryRequest struct {
	Category category.Category
}

type postCategoryResponse struct {
	Category category.Category
	Err      error `json:"err,omitempty"`
}

func (r postCategoryResponse) error() error { return r.Err }

type getCategoryRequest struct {
	ID string
}

type getCategoryResponse struct {
	Category category.Category `json:"category,omitempty"`
	Err      error             `json:"err,omitempty"`
}

func (r getCategoryResponse) error() error { return r.Err }
