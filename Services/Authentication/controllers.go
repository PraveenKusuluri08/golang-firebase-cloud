package Authentication

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"firebase.google.com/go/auth"
	"github.com/PraveenKusuluri08/helpers"
	"github.com/PraveenKusuluri08/utils"
)

func signUp(signUpUser AuthSignUp) string {
	app := utils.InitializeFbApp()

	client, err := app.Auth(context.Background())

	if err != nil {
		log.Fatal(err)
	}
	params := (&auth.UserToCreate{}).Email(signUpUser.Email).Password(signUpUser.Password)
	done, err1 := client.CreateUser(context.Background(), params)

	if err1 != nil {
		fmt.Println(err1)
		return "Failed to create user"
	}
	user := done.CustomClaims[signUpUser.Email]
	fmt.Println(user)
	db, _ := app.Firestore(context.Background())

	//hash Password
	hash, _ := helpers.PasswordHasher(signUpUser.Password)

	signUpUser.Password = hash

	//update the created at data
	signUpUser.CreatedAt = time.Now()

	//setting default role for users
	signUpUser.Role = 1

	//setting all users to the exists property true
	signUpUser.IsExists = true

	//TODO:Add the imageUrl and set the image url to the db
	//TODO:Default load the image with cover.jpg

	dbSet, err2 := db.Collection("USERS").Doc(done.UID).Set(context.TODO(), signUpUser)

	if err2 != nil {
		return "Failed to initialize with db"
	}
	fmt.Println(dbSet)
	return "User created Successfully"
}

func SignUp(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Allow-Control-Allow-Methods", "POST")

	var signUpUser AuthSignUp

	_ = json.NewDecoder(r.Body).Decode(&signUpUser)
	message := signUp(signUpUser)

	json.NewEncoder(w).Encode(message)
}
