package main

import (
	"fmt"
	"github/olegvel/ol-app/internal/server"
	"log"
)

func main() {
	fmt.Println("OL - app for everything")

	srv := server.NewServer()
	err := srv.Run()
	if err != nil {
		log.Fatal(err)
	}

	return
}
