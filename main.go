package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/PraveenKusuluri08/routes"
)

// func router() *mux.Router {
// 	router := mux.NewRouter().StrictSlash(true)
// 	secure := router.PathPrefix("/login").Subrouter()
// 	secure.Use(JwtVerify)

// 	secure.HandleFunc("/api", ApiHandler).Methods("GET")

// 	someThing := router.PathPrefix("/body").Subrouter()

// 	someThing.HandleFunc("/", ApiHandler1).Methods("GET")

// 	return router
// }

func main() {
	r := routes.Router()

	fmt.Println("App is listining")
	log.Fatal(http.ListenAndServe(":5000", r))
}

// func JwtVerify(next http.Handler) http.Handler {
// 	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 		var header = r.Header.Get("Authorization")

// 		json.NewEncoder(w).Encode(r)
// 		header = strings.Trim(header, "Bearer ")

// 		if header == "" {
// 			w.WriteHeader(http.StatusForbidden)
// 			json.NewEncoder(w).Encode("Missing auth token")
// 			return
// 		} else {
// 			json.NewEncoder(w).Encode(fmt.Sprintf("Token found. Value %s", header))
// 		}
// 		next.ServeHTTP(w, r)
// 	})
// }

// func ApiHandler(w http.ResponseWriter, r *http.Request) {
// 	w.WriteHeader(http.StatusOK)
// 	w.Header().Set("Content-Type", "application/json")
// 	json.NewEncoder(w).Encode("SUCCESS!")
// 	return
// }
// func ApiHandler1(w http.ResponseWriter, r *http.Request) {
// 	w.WriteHeader(http.StatusOK)
// 	w.Header().Set("Content-Type", "application/json")
// 	json.NewEncoder(w).Encode("🤖")
// 	return
// }
