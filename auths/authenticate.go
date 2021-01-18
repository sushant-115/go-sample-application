package auth

import (
	"fmt"
	"net/http"

	model "go-simple-app/models/mongo"
)

func Authenticate(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie("token")
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			fmt.Fprintf(w, "Unauthorized, Please login")
			return
		}
		user := &model.User{}
		err = user.DecodeToken(cookie.Value)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintf(w, "Invalid Token. Error: ", err)
			return
		}
		err = user.GetFullProfile()
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			fmt.Fprintf(w, "Unauthorized, Please login")
			return
		}
		next.ServeHTTP(w, r)
	})
}
