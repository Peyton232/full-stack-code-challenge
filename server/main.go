package main

import (
	"log"
	"net/http"

	"github.com/Peyton232/scoir-server/auth"
	"github.com/Peyton232/scoir-server/config"
	"github.com/Peyton232/scoir-server/dogs"
	"github.com/Peyton232/scoir-server/users"
	"github.com/Peyton232/scoir-server/util"
)

func main() {
	// Create instances of feature-specific services and repositories
	authService := auth.NewService()
	dogService := dogs.NewService()
	userService := users.NewService()

	// Unprotected routes
	http.HandleFunc("/login", authService.Login)

	// Protected routes - require authentication
	http.Handle("/dogs/", authService.Middleware(util.WithLogging(dogService.Router)))
	http.Handle("/", authService.Middleware(util.WithLogging(userService.Router)))

	// Start the server
	log.Printf("Server started on port %s", config.GetPort())
	err := http.ListenAndServe(":"+config.GetPort(), nil)
	if err != nil {
		log.Fatal("Error starting the server: ", err)
	}
}
