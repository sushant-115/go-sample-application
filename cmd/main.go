package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	auth "go-simple-app/auths"
	"go-simple-app/configs"
	mongodb "go-simple-app/db/mongo"
	"go-simple-app/handlers"

	"github.com/gorilla/mux"
)

func getMuxRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(false)
	router.HandleFunc("/", handlers.BaseHandler)
	router.HandleFunc("/login", handlers.LoginHandler)
	router.HandleFunc("/createProfile", handlers.CreateUserProfile).Methods("POST")
	secure := router.PathPrefix("/app").Subrouter()
	secure.Use(auth.Authenticate)
	secure.HandleFunc("/getProfile/", handlers.GetUserProfile).Methods("GET")
	secure.HandleFunc("/updateProfile/", handlers.UpdateUserProfile).Methods("PUT")
	secure.HandleFunc("/deleteProfile/", handlers.DeleteUserProfile).Methods("DELETE")
	return router
}

func main() {
	var wait time.Duration
	var configPath string
	flag.DurationVar(&wait, "graceful-timeout", time.Second*15, "the duration for which the server gracefully wait for existing connections to finish - e.g. 15s or 1m")
	flag.StringVar(&configPath, "config-path", "../configs/config.json", "Path for the config file")
	flag.Parse()

	config, err := configs.InitConfig(configPath)
	if err != nil {
		log.Fatal("Error occured while initializing the config. Error: ", err)
	}

	// Mongo client initialization
	mongoURI, err := config.MongoDB.GetURI()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(mongoURI)
	err = mongodb.InitClient(mongoURI)
	fmt.Println("Initialized mongo client", err)
	if err != nil {
		log.Fatal(err)
	}
	r := getMuxRouter()
	// Add your routes as needed

	srvAddr, err := config.Server.GetURI()
	if err != nil {
		log.Fatal(err)
	}
	srv := &http.Server{
		Addr:         srvAddr,
		WriteTimeout: config.Server.GetWriteTimeout(),
		ReadTimeout:  config.Server.GetReadTimeout(),
		IdleTimeout:  config.Server.GetIdleTimeout(),
		Handler:      r,
	}

	// Run the server in a goroutine so that it doesn't block.
	go func() {
		if err := srv.ListenAndServe(); err != nil {
			log.Println(err)
		}
	}()

	c := make(chan os.Signal, 1)
	// We'll accept graceful shutdowns when quit via SIGINT (Ctrl+C)
	// SIGKILL, SIGQUIT or SIGTERM (Ctrl+/) will not be caught.
	signal.Notify(c, os.Interrupt)

	// Block until we receive our signal.
	<-c

	// Create a deadline to wait for.
	ctx, cancel := context.WithTimeout(context.Background(), wait)
	defer cancel()
	// Shutdown the server
	srv.Shutdown(ctx)
	log.Println("shutting down")
	os.Exit(0)

}
