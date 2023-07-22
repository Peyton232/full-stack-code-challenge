package dogs

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/Peyton232/scoir-server/util"
	"github.com/gorilla/mux"
)

// Handler for "/dogs/random" route
func (svc *DogService) randomDogHandler(w http.ResponseWriter, r *http.Request) {
	resp, err := http.Get("https://dog.ceo/api/breeds/image/random")
	if err != nil {
		http.Error(w, "Invalid request", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	// Read the response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		http.Error(w, "Invalid request", http.StatusInternalServerError)
		return
	}

	// Parse the response JSON into a struct
	var apiResponse DogAPIResponse
	err = json.Unmarshal(body, &apiResponse)
	if err != nil {
		http.Error(w, "Invalid request", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(apiResponse)
}

// Handler for "/dogs/breed/list/all" route
func (svc *DogService) dogBreedListHandler(w http.ResponseWriter, r *http.Request) {
	// simply return the server''s breed list
	json.NewEncoder(w).Encode(util.Breeds)
}

// Handler for "/dogs/[breed]" route
func (svc *DogService) dogBreedHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	breed, ok := vars["breed"]
	if !ok {
		http.Error(w, "breed name was not provided", http.StatusBadRequest)
		return
	}

	resp, err := http.Get(fmt.Sprintf("https://dog.ceo/api/breed/%s/images/random", breed))
	if err != nil {
		http.Error(w, "Invalid request", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	// Read the response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		http.Error(w, "Invalid request", http.StatusInternalServerError)
		return
	}

	// Parse the response JSON into a struct
	var apiResponse DogAPIResponse
	err = json.Unmarshal(body, &apiResponse)
	if err != nil {
		http.Error(w, "Invalid request", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(apiResponse)
}
