package main

import (
	"log"
	"net/http"
	"os"

	"ezauth/internal/api"
	"ezauth/internal/auth"
	"ezauth/internal/config"
	"ezauth/internal/middleware"
)

func main() {
	// Load configuration
	cfg := config.LoadConfig()

	// Ensure data directory exists if using file storage
	if cfg.Storage == "file" {
		if _, err := os.Stat("./data"); os.IsNotExist(err) {
			os.Mkdir("./data", 0755)
		}
	}

	// Choose repository
	var repo auth.UserRepository
	var err error
	switch cfg.Storage {
	case "file":
		repo, err = auth.NewFileUserRepository(cfg.FilePath)
		if err != nil {
			log.Fatal(err)
		}
	case "memory":
		repo = auth.NewMemoryUserRepository()
	default:
		log.Fatal("unsupported storage backend")
	}

	// Create auth service
	service := auth.NewService(repo, []byte(cfg.JWTSecret), cfg.TokenExpiry)

	// Create router with JWT secret
	router := api.NewRouter(service, []byte(cfg.JWTSecret))

	// Wrap router with logging middleware
	handler := middleware.Logging(router)

	server := &http.Server{
		Addr:    ":" + cfg.Port,
		Handler: handler,
	}

	log.Println(`
 _____ _____    _____     _   _
|   __|__   |  |  _  |_ _| |_| |_
|   __|   __|  |     | | |  _|   |
|_____|_____|  |__|__|___|_| |_|_|
`)
	log.Println("EZauth server running on http://localhost:" + cfg.Port)
	log.Fatal(server.ListenAndServe())
}
