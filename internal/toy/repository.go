package toy

type Repository interface {
	Get() []Toy
	Put(t Toy)
}

type InMemory struct {
	Toys []Toy
}

func (im *InMemory) Get() []Toy {
	return im.Toys
}

func (im *InMemory) Put(t Toy) {
	im.Toys = append(im.Toys, t)
}
