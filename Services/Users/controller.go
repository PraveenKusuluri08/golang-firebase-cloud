package Users

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

func updateUser(updateUser UserUpdate, uid string) string {
	if uid != "" {
		return "Check the user uid !! and try again"
	}
	updateUser.LastUpdate = time.Now()
	info := IsUserExists(uid)
	fmt.Println(info, updateUser, uid)
	return ""
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	uid := params["uid"]

	var userUpdate UserUpdate
	_ = json.NewDecoder(r.Body).Decode(userUpdate)
	msg := updateUser(userUpdate, uid)

	fmt.Println(msg)

	r.Body.Close()

}
