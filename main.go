package main

import (

	"github.com/gorilla/mux"
	"net/http"
	"Demo-auth/handler"
	"log"
)

func main()  {

	router := mux.NewRouter()
	router.HandleFunc("/login",handler.LoginHandler).Methods("POST")
	router.Handle("/simple",handler.JwtValidator.Handler(handler.SimpleHandler)).Methods("GET")
	log.Println("Now listening...")

	http.ListenAndServe(":8080",router)
}
