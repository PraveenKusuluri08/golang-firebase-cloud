package Posts

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/PraveenKusuluri08/utils"
	"github.com/gorilla/mux"
	"google.golang.org/api/iterator"
)

var app = utils.InitializeFbApp()

//create post
func createPost(post Posts) string {

	//TODO:add this route as the closed end to get the logged in userData
	//TODO:From the token by that we can store user data into the db

	//TODO:Add validators to the Post by that we can prevent the empty data

	post.CreatedAt = time.Now()

	fmt.Println(post)

	db, _ := app.Firestore(context.Background())
	defer db.Close()

	doc := db.Collection("POSTS-GOLANG").NewDoc()

	fmt.Println(doc.ID)
	set, err := doc.Set(context.Background(), map[string]interface{}{
		"post":          post.Post,
		"createdAt":     post.CreatedAt,
		"postId":        doc.ID,
		"likesCount":    post.LikesCount,
		"commentsCount": post.CommentsCount,
		"email":         post.Email,
		"isExists":      post.IsExists,
	})

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

///////////////////////////////////////////////////////////////////////

//get all posts from db
func getAllPosts() (map[string]interface{}, error) {

	client, err := app.Firestore(context.Background())

	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()
	data := client.Collection("POSTS-GOLANG").Where("isExists", "==", true).Documents(context.Background())
	for {
		doc, err := data.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return map[string]interface{}{
				"error": "Failed to get the documents form db",
			}, err
		}
		return doc.Data(), nil
	}

	return nil, nil
}

func GetAllPosts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	data, err := getAllPosts()

	if data != nil {

		json.NewEncoder(w).Encode(data)
	} else {
		json.NewEncoder(w).Encode(err)
	}

	json.NewEncoder(w).Encode("done")
}

////////////////////////////////////////////////////////

func doComment(postId string, comment Comment) string {

	//TODO:check first post is in the db or not
	client, err := app.Firestore(context.Background())
	var checkDocExists bool
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(postId)
	data := client.Collection("POSTS-GOLANG").Where("postId", "==", postId).Documents(context.Background())
	for {
		doc, err := data.Next()
		if err == iterator.Done {
			break
		}

		if err != nil {
			fmt.Println("Post not exists")
		}
		fmt.Println(doc)
		break
	}
	comment.CommentedAt = time.Now()
	if checkDocExists {
		commentCol := client.Collection("COMMENTS-GOLANG").NewDoc()

		_, err := commentCol.Set(context.Background(), map[string]interface{}{
			"comment":      comment.Comment,
			"commentedAt":  comment.CommentedAt,
			"email":        comment.Email,
			"postId":       postId,
			"commentLikes": comment.CommentLikes,
			"commentId":    commentCol.ID,
		})

		if err != nil {
			log.Fatal(err)
		}
	}

	return "Commented successfully"
}

func DoCommnet(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Allow-Origin-Allow-Methods", "POST")

	defer r.Body.Close()
	var comment Comment
	_ = json.NewDecoder(r.Body).Decode(&comment)

	params := mux.Vars(r)

	msg := doComment(params["postId"], comment)

	json.NewEncoder(w).Encode(msg)

}
