package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/sanket2004-b/REST-API-GO/internal/config"
)

func main() {
	cfg := config.MustLoad()
	router := http.NewServeMux()

	server := http.Server{
		Addr:    cfg.Address,
		Handler: router,
	}
	err := server.ListenAndServe()
	fmt.Printf("server Started")
	if err != nil {
		log.Fatal("failed server")
	}
}
