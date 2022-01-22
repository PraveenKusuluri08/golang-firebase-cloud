package routes

import (
	"github.com/PraveenKusuluri08/Services/Authentication"
	"github.com/PraveenKusuluri08/Services/Posts"
	"github.com/PraveenKusuluri08/Services/Users"
	"github.com/gorilla/mux"
)

func Router() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)

	authentication := router.PathPrefix("/auth").Subrouter()

	authentication.HandleFunc("/signUp", Authentication.SignUp).Methods("POST")

	createPosts := router.PathPrefix("/posts").Subrouter()

	// createPosts.Use(helpers.EndPoint)

	createPosts.HandleFunc("/createpost", Posts.CreatePost).Methods("POST")

	posts := router.PathPrefix("/posts").Subrouter()

	posts.HandleFunc("/getallposts", Posts.GetAllPosts)

	users := router.PathPrefix("/users").Subrouter()

	users.HandleFunc("/updateUser/{uid}", Users.UpdateUser).Methods("PUT")

	return router

}
