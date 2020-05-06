package toy_test

import (
	"github.com/andream16/toy/internal/toy"
	"testing"
)

func TestInMemory_Get(t *testing.T) {
	t.Run("it should return one toy", func(t *testing.T) {

		const (
			name = "batman action figure"
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
