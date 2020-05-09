package transporthttp_test

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/andream16/toy/internal/toy"
	transporthttp "github.com/andream16/toy/internal/transport/http"
)

type mockOddToysManager struct{}

func (m mockOddToysManager) Get(ctx context.Context) ([]toy.Toy, error) {
	return nil, toy.OddNumberOfToysError{Number: 3}
}
func (m mockOddToysManager) Put(ctx context.Context, t toy.Toy) error { return nil }
func (m mockOddToysManager) Delete(ctx context.Context)               {}

type mockEmptyToysManager struct{}

func (m mockEmptyToysManager) Get(ctx context.Context) ([]toy.Toy, error) {
	return make([]toy.Toy, 0), nil
}
func (m mockEmptyToysManager) Put(ctx context.Context, t toy.Toy) error { return nil }
func (m mockEmptyToysManager) Delete(ctx context.Context)               {}

type mockFullToysManager struct{}

func (m mockFullToysManager) Get(ctx context.Context) ([]toy.Toy, error) {
	return []toy.Toy{
		{Name: "hulk action figure", Description: "very nice"},
		{Name: "superman action figure", Description: "very cool"},
	}, nil
}
func (m mockFullToysManager) Put(ctx context.Context, t toy.Toy) error { return nil }
func (m mockFullToysManager) Delete(ctx context.Context)               {}

type mockInvalidAttributeManager struct{}

func (m mockInvalidAttributeManager) Get(ctx context.Context) ([]toy.Toy, error) {
	return make([]toy.Toy, 0), nil
}
func (m mockInvalidAttributeManager) Put(ctx context.Context, t toy.Toy) error {
	return toy.InvalidToyError{Attribute: "name"}
}
func (m mockInvalidAttributeManager) Delete(ctx context.Context) {}

func TestHandler_GetToys(t *testing.T) {
	t.Run("it should return an APIError because the number of toys is odd", func(t *testing.T) {
		handler := transporthttp.Handler{Manager: mockOddToysManager{}}
		req := httptest.NewRequest(
			http.MethodGet,
			"/",
			nil,
		)
		rr := httptest.NewRecorder()
		handler.GetToys(rr, req)

		if rr.Code != http.StatusBadRequest {
			t.Fatalf("expected status %d, got %d", http.StatusBadRequest, rr.Code)
		}

		var apiErr transporthttp.APIError

		if err := json.NewDecoder(rr.Body).Decode(&apiErr); err != nil {
			t.Fatal(err)
		}

		if apiErr.Message == "" {
			t.Fatal("expected an error message but got an empty message")
		}
	})
	t.Run("it should return an empty list of toys", func(t *testing.T) {
		handler := transporthttp.Handler{Manager: mockEmptyToysManager{}}
		req := httptest.NewRequest(
			http.MethodGet,
			"/",
			nil,
		)
		rr := httptest.NewRecorder()
		handler.GetToys(rr, req)

		if rr.Code != http.StatusOK {
			t.Fatalf("expected status %d, got %d", http.StatusOK, rr.Code)
		}

		var toys []transporthttp.Toy

		if err := json.NewDecoder(rr.Body).Decode(&toys); err != nil {
			t.Fatal(err)
		}

		if len(toys) > 0 {
			t.Fatalf("expected no toys but got %d", len(toys))
		}
	})
	t.Run("it should return two toys", func(t *testing.T) {
		handler := transporthttp.Handler{Manager: mockFullToysManager{}}
		req := httptest.NewRequest(
			http.MethodGet,
			"/",
			nil,
		)
		rr := httptest.NewRecorder()
		handler.GetToys(rr, req)

		if rr.Code != http.StatusOK {
			t.Fatalf("expected status %d, got %d", http.StatusOK, rr.Code)
		}

		var toys []transporthttp.Toy

		if err := json.NewDecoder(rr.Body).Decode(&toys); err != nil {
			t.Fatal(err)
		}

		if len(toys) != 2 {
			t.Fatalf("expected 2 toys but got %d", len(toys))
		}
	})
}

func TestHandler_PutToy(t *testing.T) {
	t.Run("it should return bad request because the request is malformed", func(t *testing.T) {
		handler := transporthttp.Handler{}

		w := httptest.NewRecorder()
		r := httptest.NewRequest(
			http.MethodPut,
			"/",
			bytes.NewBufferString(`{`),
		)

		handler.PutToy(w, r)

		if w.Code != http.StatusBadRequest {
			t.Fatalf("expected status %d, got %d", http.StatusBadRequest, w.Code)
		}
	})
	t.Run("it should return bad request the toy name is invalid", func(t *testing.T) {
		handler := transporthttp.Handler{Manager: mockInvalidAttributeManager{}}

		w := httptest.NewRecorder()
		r := httptest.NewRequest(
			http.MethodPut,
			"/",
			bytes.NewBufferString(`{"description":"cool"}`),
		)

		handler.PutToy(w, r)

		if w.Code != http.StatusBadRequest {
			t.Fatalf("expected status %d, got %d", http.StatusBadRequest, w.Code)
		}

		var apiError transporthttp.APIError
		if err := json.NewDecoder(w.Body).Decode(&apiError); err != nil {
			t.Fatal(err)
		}

		if -1 == strings.Index(apiError.Message, "name") {
			t.Fatalf("error message %q does not contain expected attribute %q", apiError.Message, "name")
		}
	})
	t.Run("it should successfully add a new toy", func(t *testing.T) {
		handler := transporthttp.Handler{Manager: mockEmptyToysManager{}}

		w := httptest.NewRecorder()
		r := httptest.NewRequest(
			http.MethodPut,
			"/",
			bytes.NewBufferString(`{"description":"cool"}`),
		)

		handler.PutToy(w, r)

		if w.Code != http.StatusCreated {
			t.Fatalf("expected status %d, got %d", http.StatusCreated, w.Code)
		}
	})
}

func TestHandler_DeleteToy(t *testing.T) {
	t.Run("it should successfully delete a toy", func(t *testing.T) {
		handler := transporthttp.Handler{Manager: mockEmptyToysManager{}}

		w := httptest.NewRecorder()
		r := httptest.NewRequest(
			http.MethodDelete,
			"/",
			nil,
		)

		handler.DeleteToy(w, r)

		if w.Code != http.StatusOK {
			t.Fatalf("expected status %d, got %d", http.StatusOK, w.Code)
		}
	})
}
