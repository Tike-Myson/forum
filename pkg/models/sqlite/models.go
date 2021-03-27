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
	Author string `json:author`
	Created time.Time `json:created`
	ImageURL string `json:imageURL`
	Rating string `json:rating`
}

type Comment struct {
	ID int `json:`
	PostID int `json:`
	Author string `json:author`
	Content string `json:content`
	Created time.Time `json:created`
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
	Login string `json:login`
	Password string `json:password`
	Mail string `json:mail`
	Created int `json:created`
	Rating int `json:rating`
}
