package transporthttp

import (
	"encoding/json"
	"net/http"

	"github.com/andream16/toy/internal/toy"

	"github.com/gorilla/mux"
)

// Handler is a http handler wrapper that allows access to the business logic.
type Handler struct {
	Router  *mux.Router
	Manager toy.Manager
}

// APIError represents an API error.
type APIError struct {
	Message string `json:"message"`
}

// Toy represents a toy and it's used to serialise/deserialise requests/responses.
type Toy struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

// GetToys returns a list of cached toys if they are an even number.
// If an odd number of toys is cached, an error will be returned.
func (h Handler) GetToys(w http.ResponseWriter, r *http.Request) {
	toys, err := h.Manager.Get(r.Context())
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_ = json.NewEncoder(w).Encode(APIError{
			Message: err.Error(),
		})
		return
	}
	_ = json.NewEncoder(w).Encode(toHandlerToys(toys))
}

// PutToy adds a new toy to the cache.
func (h Handler) PutToy(w http.ResponseWriter, r *http.Request) {
	var t Toy
	if err := json.NewDecoder(r.Body).Decode(&t); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_ = json.NewEncoder(w).Encode(APIError{
			Message: "malformed request",
		})
		return
	}
	if err := h.Manager.Put(r.Context(), toDomainToy(t)); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_ = json.NewEncoder(w).Encode(APIError{
			Message: err.Error(),
		})
		return
	}
	w.WriteHeader(http.StatusCreated)
}

// DeleteToy deleted the oldest cached toy.
func (h Handler) DeleteToy(w http.ResponseWriter, r *http.Request) {
	h.Manager.Delete(r.Context())
}

func toHandlerToys(toys []toy.Toy) []Toy {
	res := make([]Toy, len(toys))
	for i, t := range toys {
		res[i] = Toy{Name: t.Name, Description: t.Description}
	}
	return res
}

func toDomainToy(t Toy) toy.Toy {
	return toy.Toy{Name: t.Name, Description: t.Description}
}
