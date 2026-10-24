package handlers

import (
	"fmt"
	"log"
	"net/http"
)

type Music struct {
	mux *http.ServeMux
}

func NewMusic() *Music {
	router := http.NewServeMux()
	music := &Music{mux: router}
	router.HandleFunc("GET /", music.ListMusic)
	router.HandleFunc("POST /", music.CreateMusic)
	router.HandleFunc("GET /{id}", music.GetMusic)
	router.HandleFunc("PUT /{id}", music.UpdateMusic)
	router.HandleFunc("DELETE /{id}", music.DeleteMusic)

	return music
}

func (m *Music) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	m.mux.ServeHTTP(w, r)
}

func (m *Music) ListMusic(w http.ResponseWriter, r *http.Request) {
	_, err := fmt.Fprintln(w, "List Music")
	if err != nil {
		log.Fatal(err)
	}
}

func (m *Music) CreateMusic(w http.ResponseWriter, r *http.Request) {
	_, err := fmt.Fprintln(w, "CreateMusic Music")
	if err != nil {
		log.Fatal(err)
	}
}

func (m *Music) GetMusic(w http.ResponseWriter, r *http.Request) {
	_, err := fmt.Fprintln(w, "GetMusic Music")
	if err != nil {
		log.Fatal(err)
	}
}

func (m *Music) UpdateMusic(w http.ResponseWriter, r *http.Request) {
	_, err := fmt.Fprintln(w, "UpdateMusic Music")
	if err != nil {
		log.Fatal(err)
	}
}

func (m *Music) DeleteMusic(w http.ResponseWriter, r *http.Request) {
	_, err := fmt.Fprintln(w, "DeleteMusic Music")
	if err != nil {
		log.Fatal(err)
	}
}
