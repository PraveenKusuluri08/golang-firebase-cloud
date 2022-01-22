package Users

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"time"

	"cloud.google.com/go/firestore"
	"github.com/PraveenKusuluri08/utils"
	"github.com/gorilla/mux"
)

func updateUser(updateUser UserUpdate, uid string) string {
	updateUser.LastUpdate = time.Now()

	app := utils.InitializeFbApp()
	client, err := app.Firestore(context.Background())

	if err != nil {
		log.Fatal(err)
	}
	data, err1 := client.Collection("USERS").Doc(uid).Get(context.Background())

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
