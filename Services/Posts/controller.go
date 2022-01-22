package Posts

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/PraveenKusuluri08/utils"
)

func createPost(post Posts) string {

	//TODO:add this route as the closed end to get the logged in userData
	//TODO:From the token by that we can store user data into the db

	//TODO:Add validators to the Post by that we can prevent the empty data
	app := utils.InitializeFbApp()
	post.CreatedAt = time.Now()

	fmt.Println(post)

	db, _ := app.Firestore(context.Background())

	_, set, err := db.Collection("POSTS").Add(context.Background(), post)

	fmt.Println(set)
	if err != nil {
		log.Fatal(err)
	} else {
		return "Document added successfully"
	}
	return ""
}

func CreatePost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Allow-Control-Allow-Access", "POST")
	var post Posts
	_ = json.NewDecoder(r.Body).Decode(&post)

	if msg := createPost(post); msg != "" {
		json.NewEncoder(w).Encode(msg)
	}
}

func getAllPosts() {
	app := utils.InitializeFbApp()

	client, err := app.Firestore(context.Background())

	if err != nil {
		log.Fatal(err)
	}

	data := client.Collection("POSTS").Where("isExists", "==", true)

	fmt.Println(data)
}

func GetAllPosts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	getAllPosts()

	json.NewEncoder(w).Encode("done")
}
