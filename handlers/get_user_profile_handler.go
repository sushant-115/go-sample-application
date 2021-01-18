package handlers

import (
	"fmt"
	model "go-simple-app/models/mongo"
	"net/http"
)

func GetUserProfile(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("token")
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		fmt.Fprintf(w, "Unauthorized, Please login")
	}
	user := &model.User{}
	err = user.DecodeToken(cookie.Value)
	if err != nil {
		w.WriteHeader(http.StatusExpectationFailed)
		fmt.Println(w, "Invalid Token.")
	}
	err = user.GetFullProfile()
	if err != nil {
		w.WriteHeader(http.StatusExpectationFailed)
		fmt.Fprintf(w, "Failed to get the profile")
	}
	w.WriteHeader(http.StatusOK)
	b, err := user.GetJSON()
	if err != nil {
		fmt.Println(err)
	}
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, string(b))
	return
}
