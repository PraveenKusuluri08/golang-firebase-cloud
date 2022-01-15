package routes

import (
	"github.com/PraveenKusuluri08/Services/Authentication"
	"github.com/gorilla/mux"
)

func Router() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)

	authentication := router.PathPrefix("/auth").Subrouter()

	authentication.HandleFunc("/signUp", Authentication.SignUp).Methods("POST")

	return router

}
