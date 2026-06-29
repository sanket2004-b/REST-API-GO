package main

import (
	"fmt"
	"log"
	"net/http"

	"REST-API-GO/internal/config/config.go"
)

func main() {
	cfg := config.MustLoad()
	router := http.NewServeMux()

	server := http.Server{
		Addr:    cfg.Addr,
		Handler: router,
	}
	err := server.ListenAndServe()
	fmt.Printf("server Started")
	if err != nil {
		log.Fatal("failed server")
	}
}
