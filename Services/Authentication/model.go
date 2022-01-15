package Authentication

import "time"

type AuthSignUp struct {
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	Name      string    `json:"name"`
	Role      int32     `json:"role"`
	CreatedAt time.Time `json:"created_at"`
	IsExists  bool      `json:"is_exists"`
}
