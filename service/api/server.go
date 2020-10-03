package api

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
)

type Server struct {
	*http.Server
}

func NewServer() (*Server, error) {
	log.Println("configuring server...")

	api, err := New(true)
	if err != nil {
		return nil, err
	}
	srv := http.Server{
		Addr:    ":3000",
		Handler: api,
	}

	return &Server{&srv}, nil
}

// Start runs ListenAndServe on the http.Server with graceful shutdown.
func (srv *Server) Start() {
	log.Println("starting server...")
	go func() {
		if err := srv.ListenAndServe(); err != http.ErrServerClosed {
			panic(err)
		}
	}()
	log.Printf("Listening on %s\n", srv.Addr)

	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	sig := <-quit
	log.Println("Shutting down server... Reason:", sig)
	// teardown logic...

	if err := srv.Shutdown(context.Background()); err != nil {
		panic(err)
	}
	log.Println("Server gracefully stopped")
}
