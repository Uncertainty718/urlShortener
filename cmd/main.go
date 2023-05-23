package main

import (
	"log"
	"net/http"

	"github.com/uncertainty718/urlshortener/internal/handlers"
	storage "github.com/uncertainty718/urlshortener/internal/storage/inmem"
)

func main() {
	repo := storage.NewInmem()
	r := handlers.NewHandler(repo)

	log.Fatal(http.ListenAndServe(":8080", r))
}
