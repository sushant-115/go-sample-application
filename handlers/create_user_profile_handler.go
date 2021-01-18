package handlers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	model "go-simple-app/models/mongo"
)

func CreateUserProfile(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	requestBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Error occured")
	}
	user := &model.User{}
	err = json.Unmarshal(requestBody, user)
	if err != nil {
		fmt.Fprintf(w, "Bad request")
		return
	}

	if !user.IsValidCreate() {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Required fields are missing in the request")
		return
	}

	err = user.Create()
	if err != nil {
		w.WriteHeader(http.StatusExpectationFailed)
		fmt.Printf("User creation failed. Error: %v", err)
	}
	fmt.Println("User created", *user)
	w.WriteHeader(http.StatusOK)
	b, err := user.GetJSON()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Fprintf(w, string(b))
	return
}
