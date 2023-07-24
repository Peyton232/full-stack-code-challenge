package dogs

import (
	"net/http"

	"github.com/gorilla/mux"
)

type DogAPIResponse struct {
	Message string `json:"message"`
	Status  string `json:"status"`
}

type DogService struct {
	Router *mux.Router
}

func NewService() *DogService {
	service := &DogService{
		Router: mux.NewRouter(),
	}

	// Register routes and their respective handlers
	service.Router.HandleFunc("/dogs/random", service.randomDogHandler).Methods(http.MethodGet)
	service.Router.HandleFunc("/dogs/{breed}", service.dogBreedHandler).Methods(http.MethodGet)
	service.Router.HandleFunc("/dogs/breed/list/all", service.dogBreedListHandler).Methods(http.MethodGet)

	return service
}
