package users

import (
	"encoding/json"
	"net/http"

	"github.com/Peyton232/scoir-server/util"
)

type addDogRequest struct {
	Picture string `json:"picture"`
}

func (svc *UserService) addDog(w http.ResponseWriter, r *http.Request) {
	// Decode the JSON request body into the dogRequest struct
	var request addDogRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	breed := getBreedFromURL(request.Picture)

	user := util.GetTokenFromRequest(r)
	svc.DB.addDog(user, breed, request.Picture)

	// Respond with a success message
	response := map[string]string{
		"message": "Dog added successfully",
	}
	jsonData, err := json.Marshal(response)
	if err != nil {
		http.Error(w, "Error encoding JSON response", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonData)
}

func (svc *UserService) getDogs(w http.ResponseWriter, r *http.Request) {
	user := util.GetTokenFromRequest(r)
	dogs := svc.DB.getDogs(user)

	// Respond with a success message
	response := map[string][]Dog{
		"message": dogs,
	}
	jsonData, err := json.Marshal(response)
	if err != nil {
		http.Error(w, "Error encoding JSON response", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonData)
}

func (svc *UserService) removeDog(w http.ResponseWriter, r *http.Request) {
	// Decode the JSON request body into the dogRequest struct
	var request addDogRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	user := util.GetTokenFromRequest(r)
	svc.DB.deleteDog(user, request.Picture)

	// Respond with a success message
	response := map[string]string{
		"message": "Dog added successfully",
	}
	jsonData, err := json.Marshal(response)
	if err != nil {
		http.Error(w, "Error encoding JSON response", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonData)
}

func (svc *UserService) clearDogs(w http.ResponseWriter, r *http.Request) {
	user := util.GetTokenFromRequest(r)
	svc.DB.deleteAllDog(user)

	// Respond with a success message
	response := map[string]string{
		"message": "Dog added successfully",
	}
	jsonData, err := json.Marshal(response)
	if err != nil {
		http.Error(w, "Error encoding JSON response", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonData)
}
