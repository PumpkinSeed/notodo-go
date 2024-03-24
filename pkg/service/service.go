package service

import (
	"net/http"
	"os"
)

func Run() {
	if err := http.ListenAndServe(addr(), routes()); err != nil {
		panic(err)
	}
}

func routes() *http.ServeMux {
	router := http.NewServeMux()

	router.HandleFunc("POST /items", createItem)
	router.HandleFunc("GET /items", getItems)
	router.HandleFunc("GET /items/{key}", getItem)

	return router
}

func addr() string {
	if address := os.Getenv("APP_ADDR"); address != "" {
		return address
	}
	return ":8080"
}
