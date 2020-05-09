package toy_test

import (
	"errors"
	"testing"

	"github.com/andream16/toy/internal/toy"
)

func TestService(t *testing.T) {
	t.Run("should implement Manager", func(t *testing.T) {
		var _ toy.Manager = toy.Service{}
	})
}

type mockRepoNoToys struct{}

func (m mockRepoNoToys) Get() []toy.Toy { return make([]toy.Toy, 0) }
func (m mockRepoNoToys) Put(t toy.Toy)  {}
func (m mockRepoNoToys) Delete()        {}

type mockRepoEvenToys struct{}

func (m mockRepoEvenToys) Get() []toy.Toy {
	return []toy.Toy{
		{
			Name:        "gopher plushie",
			Description: "amazing golang gopher plushie",
		},
		{
			Name:        "docker plushie",
			Description: "amazing docker gopher plushie",
		},
	}
}
func (m mockRepoEvenToys) Put(t toy.Toy) {}
func (m mockRepoEvenToys) Delete()       {}

type mockRepoOddToys struct{}

func (m mockRepoOddToys) Get() []toy.Toy {
	return []toy.Toy{
		{
			Name:        "gopher plushie",
			Description: "amazing golang gopher plushie",
		},
	}
}
func (m mockRepoOddToys) Put(t toy.Toy) {}
func (m mockRepoOddToys) Delete()       {}

func TestService_Get(t *testing.T) {
	t.Run("it should return an empty list because no toys exist", func(t *testing.T) {
		svc := toy.NewService(mockRepoNoToys{})
		toys, err := svc.Get()
		if err != nil {
			t.Fatalf("expected no error but got: %s", err)
		}
		if len(toys) > 0 {
			t.Fatal("expected no toys")
		}
	})
	t.Run("it should return 2 toys", func(t *testing.T) {
		svc := toy.NewService(mockRepoEvenToys{})
		toys, err := svc.Get()
		if err != nil {
			t.Fatalf("expected no error but got: %s", err)
		}
		if len(toys) != 2 {
			t.Fatalf("expected 2 toys, got %d", len(toys))
		}
	})
	t.Run("it should return OddNumberOfToysError", func(t *testing.T) {
		svc := toy.NewService(mockRepoOddToys{})
		toys, err := svc.Get()
		if err == nil {
			t.Fatal("expected OddNumberOfToysError, got nil")
		}
		var e toy.OddNumberOfToysError
		if !errors.As(err, &e) {
			t.Fatalf("expected OddNumberOfToysError, got %s", err.Error())
		}
		if len(toys) != 0 {
			t.Fatalf("expected 0 toys, got %d", len(toys))
		}
	})
}

func TestService_Put(t *testing.T) {
	t.Run("it should return an error because the toy name is not valid", func(t *testing.T) {
		svc := toy.NewService(nil)
		err := svc.Put(toy.Toy{})
		if err == nil {
			t.Fatalf("expected InvalidToyError, got nil")
		}
		var e toy.InvalidToyError
		if !errors.As(err, &e) {
			t.Fatalf("expected InvalidToyError, got %s", err.Error())
		}
		if e.Attribute != "name" {
			t.Fatalf("expected InvalidToyError for invalid attribute name but it was %s", e.Attribute)
		}
	})
	t.Run("it should return an error because the toy description is not valid", func(t *testing.T) {
		svc := toy.NewService(nil)
		err := svc.Put(toy.Toy{Name: "john wick action figure"})
		if err == nil {
			t.Fatalf("expected InvalidToyError, got nil")
		}
		var e toy.InvalidToyError
		if !errors.As(err, &e) {
			t.Fatalf("expected InvalidToyError, got %s", err.Error())
		}
		if e.Attribute != "description" {
			t.Fatalf("expected InvalidToyError for invalid attribute description but it was %s", e.Attribute)
		}
	})
	t.Run("it should add a new toy to the collection", func(t *testing.T) {
		svc := toy.NewService(mockRepoNoToys{})
		err := svc.Put(toy.Toy{Name: "john wick action figure", Description: "neat"})
		if err != nil {
			t.Fatalf("expected error %s", err.Error())
		}
	})
}

func TestService_Delete(t *testing.T) {
	t.Run("it should call repository's delete", func(t *testing.T) {
		toy.NewService(&toy.InMemory{}).Delete()
	})
}
