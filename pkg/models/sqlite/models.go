package models

import (
	"errors"
	"time"
)

var ErrNoRecord = errors.New("models: no matching record found")

type Post struct {
	ID int `json:"id,string,omitempty"`
	Title string `json:title`
	Content string `json:content`
	Author User `json:author`
	CreatedAt time.Time `json:created_at`
	ImageURL string `json:image_url`
	Rating int `json:rating`
	Comments []Comment `json:comments`
}

type Comment struct {
	ID int `json:id`
	PostID int `json:post_id`
	Author string `json:author`
	Content string `json:content`
	CreatedAt time.Time `json:created_at`
	Rating int `json:rating`
}

type Category struct {
	ID int `json:id, string, omitempty`
	Name string `json:name`
}

type RatingPost struct {
	PostID int `json:id, string, omitempty`
	Author string `json:author`
	Value int `json:value`
}

type RatingComment struct {
	CommentID int `json:`
	Author int `json:`
	Value int `json:`
}

type CategoryPostLink struct {
	PostID int `json:id, string, omitempty`
	CategoryID int `json:id, string, omitempty`
}

type User struct {
	ID int `json:id, string, omitempty`
	Username string `json:username`
	Password string `json:password`
	Email string `json:email`
	CreatedAt time.Time `json:created_at`
	Rating int `json:rating`
}
