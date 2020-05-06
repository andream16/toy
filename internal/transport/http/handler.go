package http

import (
	"github.com/gorilla/mux"
)

type Handler struct {
	Router *mux.Router
}
