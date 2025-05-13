package main

import (
	"context"
	"example.com/testserver/db"
	"example.com/testserver/routes"
	"log"
	"net/http"
	"os"
	"time"
	"github.com/gin-gonic/gin"
)

var shutdownChan = routes.ShutdownChan

func main() {
	db.InitDB()
	server := gin.Default()

    routes.RegisterRoutes(server)

	// Create an http.Server with the Gin engine as its handler
	srv := &http.Server{
		Addr:    ":8080",
		Handler: server,
	}

	// Start the server in a goroutine so that it doesn't block.
	go func() {
		log.Println("Server starting on :8080")
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	<-shutdownChan
	log.Println("Shutdown signal received. Shutting down server gracefully...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalf("Server forced to shutdown: %v", err)
	}
	log.Println("Server exited properly")
	os.Exit(0)
}


