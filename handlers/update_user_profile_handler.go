package handlers

import (
	"encoding/json"
	"fmt"
	model "go-simple-app/models/mongo"
	"io/ioutil"
	"net/http"
)

// UpdateUserProfile handler
func UpdateUserProfile(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("token")
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		fmt.Fprintf(w, "Unauthorized, Please login")
	}
	user := &model.User{}
	err = user.DecodeToken(cookie.Value)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Invalid Token")
	}
	requestBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusExpectationFailed)
		fmt.Fprintf(w, "Error occured")
	}
	updatedUser := &model.User{}
	err = json.Unmarshal(requestBody, updatedUser)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Bad request")
		return
	}

	err = user.UpdateProfile(updatedUser)
	if err != nil {
		fmt.Fprintf(w, "Failed to update the profile")
	}
	return
}
