package utils

import (
	"context"
	"log"

	firebase "firebase.google.com/go"
	"google.golang.org/api/option"
)

func InitializeFbApp() *firebase.App {
	opt := option.WithCredentialsJSON([]byte(`<Config File>`))

	app, err := firebase.NewApp(context.Background(), nil, opt)

	if err != nil {
		log.Fatal(err)
	}
	return app
}
