package models

import "time"

type Post struct {
	ID int
	Title string
	Content string
	Author string
	Created time.Time
	ImageURL string
}

type Comment struct {
	ID int
	PostID int
	Author string
	Content string
	Created time.Time
}

type Category struct {
	ID int
	Name string
}

type RatingPost struct {
	PostID int
	Author string
	Value int
}

type RatingComment struct {
	CommentID int
	Author int
	Value int
}

type CategoryPostLink struct {
	PostID int
	CategoryID int
}

type User struct {
	login string
	password string
	mail string
}
