package auth

import (
	"encoding/json"
	"net/http"
)

// AuthRequest represents the request data for authentication (e.g., username and password).
type AuthRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// AuthResponse represents the response data for authentication (e.g., access token).
type AuthResponse struct {
	Token string `json:"token"`
}

func (svc *AuthService) Login(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Parse the request body to get the authentication credentials (username and password).
	var request AuthRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	// Validate the authentication credentials.
	if !svc.isValidCredentials(request.Username, request.Password) {
		http.Error(w, "Invalid credentials", http.StatusUnauthorized)
		return
	}

	// Assuming you have a function to generate an access token upon successful authentication.
	token := svc.generateAccessToken(request.Username)

	// Create the response and send it back to the client.
	response := AuthResponse{Token: token}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

// isValidCredentials checks if the provided username and password are non-empty
func (svc *AuthService) isValidCredentials(username, password string) bool {
	return username != "" && password != ""
}

// generateAccessToken generates an access token for the authenticated user.
func (svc *AuthService) generateAccessToken(username string) string {
	// set the access token as the user's username
	// this is a security flaw
	return username
}
