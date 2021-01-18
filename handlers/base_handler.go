package handlers

import (
	"encoding/json"
	"net/http"
)

func BaseHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode("Welcome to the application!")
	json.NewEncoder(w).Encode("Please login or Sign up to continue...")
	return
}
