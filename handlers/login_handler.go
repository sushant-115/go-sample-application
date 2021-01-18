package handlers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	model "go-simple-app/models/mongo"
)

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	requestBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Error occured")
	}
	user := &model.User{}
	err = json.Unmarshal(requestBody, user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Bad request")
		return
	}

	if !user.IsValidLogin() {
		w.WriteHeader(http.StatusUnauthorized)
		fmt.Fprintf(w, "Credential missing or invalid")
		return
	}

	cookie := http.Cookie{
		Name:    "token",
		Value:   user.GetToken(),
		Expires: time.Now().Add(1 * time.Hour),
	}

	err = user.GetFullProfile()
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		fmt.Fprintf(w, "Login failed. Please check the credentials")
		return
	}
	http.SetCookie(w, &cookie)
	fmt.Fprintf(w, "Login successfully")
	return
}
