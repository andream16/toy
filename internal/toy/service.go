package toy

import (
	"context"
	"fmt"
)

// Manager describes the business logic.
type Manager interface {
	Get(ctx context.Context) ([]Toy, error)
	Put(ctx context.Context, t Toy) error
	Delete(ctx context.Context)
}

// Service implements the business logic.
type Service struct {
	Repository Repository
}

// OddNumberOfToysError is returned when there's an odd number of toys.
type OddNumberOfToysError struct {
	Number int
}

func (o OddNumberOfToysError) Error() string {
	return fmt.Sprintf("odd number of toys %d", o.Number)
}

// InvalidToyError is returnes when a toy has one or more invalid attributes.
type InvalidToyError struct {
	Attribute string
}

func (i InvalidToyError) Error() string {
	return fmt.Sprintf("invalid attribute %s", i.Attribute)
}

// NewService returns a new service.
func NewService(repo Repository) Service {
	return Service{Repository: repo}
}

// Get returns cached toys if their quantity is even. OddNumberOfToysError is returned if there's an odd number of them.
func (s Service) Get(ctx context.Context) ([]Toy, error) {
	toys := s.Repository.Get(ctx)
	if len(toys)%2 != 0 {
		return nil, OddNumberOfToysError{Number: len(toys)}
	}
	return toys, nil
}

// Put adds a toy to the cache.
func (s Service) Put(ctx context.Context, t Toy) error {
	switch {
	case t.Name == "":
		return InvalidToyError{Attribute: "name"}
	case t.Description == "":
		return InvalidToyError{Attribute: "description"}
	}
	s.Repository.Put(ctx, t)
	return nil
}

// Delete deletes the oldest cached toy.
func (s Service) Delete(ctx context.Context) {
	s.Repository.Delete(ctx)
}
