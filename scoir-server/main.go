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
	err := http.ListenAndServe(":"+config.GetPort(), corsHandler(http.DefaultServeMux))
	if err != nil {
		log.Fatal("Error starting the server: ", err)
	}
}

func corsHandler(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Set the CORS headers
		// Allow requests from localhost:3000
		w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		// Handle preflight requests (OPTIONS method)
		if r.Method == "OPTIONS" {
			// Respond with a 204 No Content status for preflight requests
			w.WriteHeader(http.StatusNoContent)
			return
		}

		// Call the actual handler
		handler.ServeHTTP(w, r)
	})
}
