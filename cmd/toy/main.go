package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/andream16/toy/internal/toy"
	transporthttp "github.com/andream16/toy/internal/transport/http"

	"github.com/gorilla/mux"
	"golang.org/x/sync/errgroup"
)

func main() {

	const address = "0.0.0.0:8080"

	var (
		ctx, cancel = context.WithCancel(context.Background())
		handler     = transporthttp.Handler{
			Router:  mux.NewRouter(),
			Manager: toy.NewService(&toy.InMemory{}),
		}
	)

	defer cancel()

	handler.Router.HandleFunc("/", handler.GetToys).Methods(http.MethodGet)
	handler.Router.HandleFunc("/", handler.PutToy).Methods(http.MethodPut)
	handler.Router.HandleFunc("/", handler.DeleteToy).Methods(http.MethodDelete)

	srv := &http.Server{
		Addr:         address,
		WriteTimeout: time.Second * 5,
		ReadTimeout:  time.Second * 10,
		IdleTimeout:  time.Second * 60,
		Handler:      handler.Router,
	}

	g, ctx := errgroup.WithContext(ctx)

	g.Go(func() error {
		log.Println(fmt.Sprintf("serve listening on %s...", address))
		if err := srv.ListenAndServe(); err != http.ErrServerClosed {
			return err
		}
		return nil
	})

	g.Go(func() error {
		<-ctx.Done()

		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		return srv.Shutdown(ctx)
	})

	if err := g.Wait(); err != nil {
		log.Println(err)
	}

}
