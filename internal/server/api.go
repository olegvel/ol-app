package server

import (
	"fmt"
	"github/olegvel/ol-app/internal/handlers"
	"log"
	"net/http"
)

func NewAPIServer() http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("/hello", HelloHandler)
	mux.Handle("/media/", http.StripPrefix("/media", handlers.NewMedia()))
	mux.Handle("/music/", http.StripPrefix("/music", handlers.NewMusic()))

	return mux
}

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	_, err := fmt.Fprintln(w, "Hello, World! "+r.RequestURI)

	if err != nil {
		log.Fatal(err)
	}
}
