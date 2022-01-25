package Users

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"cloud.google.com/go/firestore"
	"github.com/PraveenKusuluri08/utils"
	"github.com/gorilla/mux"
	"google.golang.org/api/iterator"
)

var app = utils.InitializeFbApp()

func updateUser(updateUser UserUpdate, uid string) string {
	updateUser.LastUpdate = time.Now()

	client, err := app.Firestore(context.Background())

	if err != nil {
		log.Fatal(err)
	}
	data, err1 := client.Collection("USERS-GOLANG").Doc(uid).Get(context.Background())

	if err1 != nil {
		log.Fatal(err1)
	}
	data.Ref.Update(context.Background(), []firestore.Update{
		{
			Path:  "Name",
			Value: updateUser.Name,
		},
		{
			Path:  "lastUpdate",
			Value: updateUser.LastUpdate,
		},
	})

	return "User Update successfully"
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	defer r.Body.Close()

	uid := params["uid"]

	var userUpdate UserUpdate
	_ = json.NewDecoder(r.Body).Decode(&userUpdate)
	msg := updateUser(userUpdate, uid)

	json.NewEncoder(w).Encode(msg)

}

func getAuthUserData(uid string, email string) interface{} {
	var user map[string]interface{}
	client, err := app.Firestore(context.Background())

	if err != nil {
		log.Fatal(err)
	}
	data, err1 := client.Collection("USERS-GOLANG").Doc(uid).Get(context.Background())

	if err1 != nil {
		log.Fatal(err1)
	}
	fmt.Println(data.Data())
	user = data.Data()

	postsData := client.Collection("POSTS-GOLANG").Where("email", "==", email).Documents(context.Background())

	for {
		doc, err2 := postsData.Next()

		if err2 == iterator.Done {
			log.Fatal(err)
			break
		}

		if err2 != nil {
			log.Fatal(err2)
			break
		}
		user = doc.Data()

	}
	return user
}

func GetAuthUserData(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Allow-Origin-Allow-Methods", "GET")

	params := mux.Vars(r)

	user := getAuthUserData(params["uid"], params["email"])

	json.NewEncoder(w).Encode(user)
}
