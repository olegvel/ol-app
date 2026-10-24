package handlers

import (
	"fmt"
	"log"
	"net/http"
)

type Media struct {
	mux *http.ServeMux
}

func NewMedia() *Media {
	router := http.NewServeMux()
	media := &Media{mux: router}
	router.HandleFunc("GET /", media.ListMedia)
	router.HandleFunc("POST /", media.CreateMedia)
	router.HandleFunc("GET /{id}", media.GetMedia)
	router.HandleFunc("PUT /{id}", media.UpdateMedia)
	router.HandleFunc("DELETE /{id}", media.DeleteMedia)

	return media
}

func (m *Media) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	m.mux.ServeHTTP(w, r)
}

func (m *Media) ListMedia(w http.ResponseWriter, r *http.Request) {
	_, err := fmt.Fprintln(w, "List media")
	if err != nil {
		log.Fatal(err)
	}
}

func (m *Media) CreateMedia(w http.ResponseWriter, r *http.Request) {
	_, err := fmt.Fprintln(w, "CreateMedia media")
	if err != nil {
		log.Fatal(err)
	}
}

func (m *Media) GetMedia(w http.ResponseWriter, r *http.Request) {
	_, err := fmt.Fprintln(w, "GetMedia media")
	if err != nil {
		log.Fatal(err)
	}
}

func (m *Media) UpdateMedia(w http.ResponseWriter, r *http.Request) {
	_, err := fmt.Fprintln(w, "UpdateMedia media")
	if err != nil {
		log.Fatal(err)
	}
}

func (m *Media) DeleteMedia(w http.ResponseWriter, r *http.Request) {
	_, err := fmt.Fprintln(w, "DeleteMedia media")
	if err != nil {
		log.Fatal(err)
	}
}
