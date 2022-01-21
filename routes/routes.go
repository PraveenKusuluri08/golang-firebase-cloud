package routes

import (
	"github.com/PraveenKusuluri08/Services/Authentication"
	"github.com/PraveenKusuluri08/Services/Posts"
	"github.com/PraveenKusuluri08/Services/Users"
	"github.com/PraveenKusuluri08/helpers"
	"github.com/gorilla/mux"
)

func Router() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)

	authentication := router.PathPrefix("/auth").Subrouter()

	authentication.HandleFunc("/signUp", Authentication.SignUp).Methods("POST")

	posts := router.PathPrefix("/posts").Subrouter()

	posts.Use(helpers.EndPoint)

	posts.HandleFunc("/createpost", Posts.CreatePost).Methods("POSTS")

	users := router.PathPrefix("/users").Subrouter()

	users.HandleFunc("/updateUser", Users.UpdateUser).Methods("PUT")

	return router

}
