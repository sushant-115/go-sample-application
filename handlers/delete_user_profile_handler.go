package handlers

import (
	"fmt"
	model "go-simple-app/models/mongo"
	"net/http"
)

func DeleteUserProfile(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("token")
	if err != nil {
		fmt.Fprintf(w, "Unauthorized, Please login")
	}
	user := &model.User{}
	err = user.DecodeToken(cookie.Value)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Invalid Token")
	}

	err = user.DeleteProfile()
	if err != nil {
		w.WriteHeader(http.StatusExpectationFailed)
		fmt.Fprintf(w, "Failed to delete the profile.")
	}
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Successfully deleted the profile")
	return
}
