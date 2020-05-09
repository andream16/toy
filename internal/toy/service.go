package toy

import "fmt"

// Manager describes the business logic.
type Manager interface {
	Get() ([]Toy, error)
	Put(t Toy) error
	Delete()
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
func (s Service) Get() ([]Toy, error) {
	toys := s.Repository.Get()
	if len(toys)%2 != 0 {
		return nil, OddNumberOfToysError{Number: len(toys)}
	}
	return toys, nil
}

// Put add a toy to the cache.
func (s Service) Put(t Toy) error {
	switch {
	case t.Name == "":
		return InvalidToyError{Attribute: "name"}
	case t.Description == "":
		return InvalidToyError{Attribute: "description"}
	}
	s.Repository.Put(t)
	return nil
}

// Delete deletes the oldest cached toy.
func (s Service) Delete() {
	s.Repository.Delete()
}
