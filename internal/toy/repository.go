package toy

import "context"

// Repository contains the repository layer.
type Repository interface {
	Get(ctx context.Context) []Toy
	Put(ctx context.Context, t Toy)
	Delete(ctx context.Context)
}

// InMemory is the in-memory concrete implementation of the repository.
type InMemory struct {
	Toys []Toy
}

// Get returns all toys.
func (im *InMemory) Get(ctx context.Context) []Toy {
	return im.Toys
}

// Put add a toy to the cache.
func (im *InMemory) Put(ctx context.Context, t Toy) {
	im.Toys = append(im.Toys, t)
}

// Delete removes the oldest toy from the cache.
func (im *InMemory) Delete(ctx context.Context) {
	if len(im.Toys) == 0 {
		return
	}
	im.Toys = im.Toys[1:]
}
