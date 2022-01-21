package Users

import "time"

type UserUpdate struct {
	Name       string    `json:"name"`
	LastUpdate time.Time `json:"lastUpdate"`
}
