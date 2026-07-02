package main

import (
	"context"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/sanket2004-b/REST-API-GO/internal/config"
	"github.com/sanket2004-b/REST-API-GO/internal/http/handler/student"
)

func main() {
	cfg := config.MustLoad()
	router := http.NewServeMux()
	// router.HandleFunc("POST /api", func(w http.ResponseWriter, r *http.Request) {
	// 	slog.Info("Matched /api", "method", r.Method)
	// 	w.Write([]byte("Method: " + r.Method))
	// })
	router.HandleFunc("POST/api", student.New())
	router.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Welcome this is first go api"))
	})

	server := http.Server{
		Addr:    cfg.Address,
		Handler: router,
	}
	// fmt.Println("Server starting on", cfg.Address)
	slog.Info("Server starting on", slog.String("address", cfg.Address))
	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		err := server.ListenAndServe()
		if err != nil {
			slog.Error("server is failed...", slog.String("error", err.Error()))
		}
	}()
	// err := server.ListenAndServe()
	// // fmt.Printf("server Started")
	// if err != nil {
	// 	log.Fatal("failed server")
	// }

	<-done
	slog.Info("shutting down the server")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	err := server.Shutdown(ctx)

	if err != nil {
		slog.Error("server shutdown failed", "error", err.Error())
	}

	slog.Info("server exited properly")

}
