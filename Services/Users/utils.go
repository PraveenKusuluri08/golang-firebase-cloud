package Users

import (
	"context"
	"fmt"
	"log"

	"cloud.google.com/go/firestore"
	"github.com/PraveenKusuluri08/utils"
)

func IsUserExists(uid string) firestore.Query {
	app := utils.InitializeFbApp()
	client, err := app.Firestore(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	data := client.Collection("USERS").Where("uid", "==", uid).Where("IsExists", "==", true).Limit(1)
	fmt.Println("data", data)
	return data
}
