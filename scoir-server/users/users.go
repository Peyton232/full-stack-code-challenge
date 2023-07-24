package users

import (
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

type UserService struct {
	Router *mux.Router
	DB     *UserDB
}

func NewService() *UserService {
	service := &UserService{
		Router: mux.NewRouter(),
		DB:     NewDB(),
	}

	// Register routes and their respective handlers
	service.Router.HandleFunc("/users/add", service.addDog).Methods(http.MethodPost)
	service.Router.HandleFunc("/users/remove", service.removeDog).Methods(http.MethodDelete)
	service.Router.HandleFunc("/users/clear", service.clearDogs).Methods(http.MethodDelete)
	service.Router.HandleFunc("/users/all", service.getDogs).Methods(http.MethodGet)

	return service
}

// TODO make sure this works with both URL types
func getBreedFromURL(url string) string {
	// Split the URL by "/"
	parts := strings.Split(url, "/")

	// Find the index of "breeds" in the parts slice
	breedIndex := -1
	for i, part := range parts {
		if part == "breeds" {
			breedIndex = i
			break
		}
	}

	// Check if "breeds" was found and there is a word after it
	if breedIndex != -1 && breedIndex+1 < len(parts) {
		return parts[breedIndex+1]
	}

	// Return an empty string if "breeds" was not found or there's no word after it
	return ""
}
