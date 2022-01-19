package Posts

import "time"

type Posts struct {
	Post          string    `json:"Post"`
	CreatedAt     time.Time `json:"CreatedAt"`
	Email         string    `json:"Email"`
	IsExists      bool      `json:"IsExists"`
	Uid           string    `json:"uid"`
	LikesCount    int64     `json:"likesCount"`
	CommentsCount int64     `json:"commentsCount"`
	UserImage     string    `json:"userImage"`
}
