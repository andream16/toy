package toy

import "fmt"

type Service struct {
	Repository Repository
}

type OddNumberOfToysError struct {
	Number int
}

func (o OddNumberOfToysError) Error() string {
	return fmt.Sprintf("odd number of toys %d", o.Number)
}

type InvalidToyError struct {
	Attribute string
}

func (i InvalidToyError) Error() string {
	return fmt.Sprintf("invalid attribute %s", i.Attribute)
}

func NewService(repo Repository) Service {
	return Service{Repository: repo}
}

func (s Service) Get() ([]Toy, error) {
	toys := s.Repository.Get()
	if len(toys) % 2 != 0 {
		return nil, OddNumberOfToysError{Number: len(toys)}
	}
	return toys, nil
}

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
