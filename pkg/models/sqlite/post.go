package models

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func CreatePostsTable(db *sql.DB) {
	postsTable, err := db.Prepare(createPostsTableSQL)
	if err != nil {
		log.Fatal(err.Error())
	}

	_, err = postsTable.Exec()
	if err != nil {
		log.Fatal(err.Error())
	}

	log.Println("Post table created in DB")
}

func insertPostIntoDB(db *sql.DB, postData Post) {
	insertPost, err := db.Prepare(insertPostSQL)
	if err != nil {
		log.Fatal(err)
	}

	_, err = insertPost.Exec(
		postData.Author.ID,
		postData.Author.Username,
		postData.Title,
		postData.Content,
		postData.CreatedAt,
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("New post inserted into DB")
}

func InsertPost(db *sql.DB, w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	var post Post
	err := json.NewDecoder(r.Body).Decode(&post)
	if err != nil {
		log.Fatal(err)
	}

	insertPostIntoDB(db, post)
	err = json.NewEncoder(w).Encode(post)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("New post sent to clientside")
}

func getAllPosts(db *sql.DB) []Post {
	var posts []Post

	row, err := db.Query(getAllPostsSQL)
	if err != nil {
		log.Fatal(err)
	}
	defer row.Close()

	var author User

	for row.Next() {
		var post = Post{
			Author: author,
		}
		err = row.Scan(
			&post.ID,
			&post.Author.ID,
			&post.Author.Username,
			&post.Title,
			&post.Content,
			&post.CreatedAt,
		)

		if err != nil {
			log.Fatal(err)
		}
		posts = append(posts, post)
	}
	return posts
}

func SendAllPosts(db *sql.DB, w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	posts := getAllPosts(db)
	json.NewEncoder(w).Encode(posts)

	fmt.Println("All posts sent to clientside")
}