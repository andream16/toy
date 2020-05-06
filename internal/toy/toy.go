package toy

type Toy struct {
	Name string
	Description string
}

type Manager interface {
	Get() ([]Toy, error)
	Put(t Toy) error
}
