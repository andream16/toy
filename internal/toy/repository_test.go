package toy_test

import (
	"testing"

	"github.com/andream16/toy/internal/toy"
)

func TestInMemory(t *testing.T) {
	t.Run("should implement Repository", func(t *testing.T) {
		var _ toy.Repository = &toy.InMemory{}
	})
}

func TestInMemory_Get(t *testing.T) {
	t.Run("it should return one toy", func(t *testing.T) {

		const (
			name        = "batman action figure"
			description = "pretty cool"
		)

		repo := &toy.InMemory{
			Toys: []toy.Toy{
				{
					Name:        name,
					Description: description,
				},
			},
		}
		toys := repo.Get()
		if len(toys) != 1 {
			t.Fatalf("expected 1 toy, got %d", len(toys))
		}
		if toys[0].Name != name {
			t.Fatalf("expected toy name to be %s, got %s", name, toys[0].Name)
		}
		if toys[0].Description != description {
			t.Fatalf("expected toy description to be %s, got %s", description, toys[0].Description)
		}
	})
}

func TestInMemory_Put(t *testing.T) {
	t.Run("it should add another toy for a total of two toys", func(t *testing.T) {
		repo := &toy.InMemory{
			Toys: []toy.Toy{
				{
					Name:        "batman action figure",
					Description: "pretty cool",
				},
			},
		}
		repo.Put(toy.Toy{
			Name:        "spiderman action figure",
			Description: "meh",
		})
		toys := repo.Get()
		if len(toys) != 2 {
			t.Fatalf("expected 2 toys, got %d", len(toys))
		}
	})
}

func TestInMemory_Delete(t *testing.T) {
	t.Run("it should delete the oldest toy from the cache", func(t *testing.T) {
		const (
			remainingToyName        = "batman action figure"
			remainingToyDescription = "pretty cool"
		)

		repo := &toy.InMemory{}

		for _, t := range []toy.Toy{
			{
				Name:        "jotaro kujo action figure",
				Description: "amazing",
			},
			{
				Name:        remainingToyName,
				Description: remainingToyDescription,
			},
		} {
			repo.Put(t)
		}

		repo.Delete()
		toys := repo.Get()
		if len(toys) != 1 {
			t.Fatalf("expected 1 toys, got %d", len(toys))
		}
		if remainingToyName != toys[0].Name {
			t.Fatalf("expected toy name to be %q, got %q", remainingToyName, toys[0].Name)
		}
		if remainingToyDescription != toys[0].Description {
			t.Fatalf("expected toy description to be %q, got %q", remainingToyDescription, toys[0].Description)
		}
	})
	t.Run("it should safely execute without causing panic even if the cache is empty", func(t *testing.T) {
		const (
			remainingToyName        = "batman action figure"
			remainingToyDescription = "pretty cool"
		)
		repo := &toy.InMemory{}
		repo.Delete()
		toys := repo.Get()
		if len(toys) != 0 {
			t.Fatalf("expected 0 toys, got %d", len(toys))
		}
	})
}
