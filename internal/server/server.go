package server

import (
	"fmt"
	"net/http"
)

type Server struct {
	handler http.Handler
}

func NewServer() *Server {
	handler := http.NewServeMux()

	handler.Handle("/", NewWebServer())
	handler.Handle("/api/", http.StripPrefix("/api", NewAPIServer()))

	return &Server{
		handler: handler,
	}
}

func (s *Server) Run() error {
	fmt.Println("Listening on :1512...")
	err := http.ListenAndServe(":1512", s.handler)
	if err != nil {
		return err
	}

	return nil
}
