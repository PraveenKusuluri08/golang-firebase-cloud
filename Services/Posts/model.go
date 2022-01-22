package Posts

import "time"

type Posts struct {
	Post          string    `json:"post"`
	CreatedAt     time.Time `json:"createdAt"`
	Email         string    `json:"email"`
	IsExists      bool      `json:"isExists"`
	Uid           string    `json:"uid"`
	LikesCount    int64     `json:"likesCount"`
	CommentsCount int64     `json:"commentsCount"`
	UserImage     string    `json:"userImage"`
}
