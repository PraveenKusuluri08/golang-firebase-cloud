package helpers

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"

	"github.com/PraveenKusuluri08/utils"
)

func EndPoint(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var header = r.Header.Get("Authorization")

		json.NewEncoder(w).Encode(r)
		header = strings.Trim(header, " ")

		if header == "" {
			w.WriteHeader(http.StatusForbidden)
			json.NewEncoder(w).Encode("UnAuthorised")
			return
		} else {
			app := utils.InitializeFbApp()
			client, err := app.Auth(context.Background())
			if err != nil {
				log.Fatal(err)
			} else {
				token, err := client.VerifyIDToken(context.Background(), header)
				if err != nil {
					log.Fatal(err)
				}
				expire := token.Expires
				fmt.Println(expire)

				if expire == 0 {
					json.NewEncoder(w).Encode("Please verify the token!! Token is expired")
				} else {
					tokenInfo := token.UID
					fmt.Println(tokenInfo)
					r.Body = io.NopCloser(strings.NewReader(tokenInfo))
					next.ServeHTTP(w, r)
				}
			}
		}
		next.ServeHTTP(w, r)
	})
}
