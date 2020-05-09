package toy

// Repository contains the repository layer.
type Repository interface {
	Get() []Toy
	Put(t Toy)
	Delete()
}

// InMemory is the in-memory concrete implementation of the repository.
type InMemory struct {
	Toys []Toy
}

// Get returns all toys.
func (im *InMemory) Get() []Toy {
	return im.Toys
}

// Put add a toy to the cache.
func (im *InMemory) Put(t Toy) {
	im.Toys = append(im.Toys, t)
}

// Delete removes the oldest toy from the cache.
func (im *InMemory) Delete() {
	if len(im.Toys) == 0 {
		return
	}
	im.Toys = im.Toys[1:]
}
