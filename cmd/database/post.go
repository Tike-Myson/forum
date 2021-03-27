package database

import (
	"database/sql"
	"encoding/json"
	"fmt"
	models "github.com/Tike-Myson/forum/pkg/models/sqlite"
	"log"
	"net/http"
)

func CreatePostsTable(db *sql.DB) {
	postsTable, err := db.Prepare(models.CreatePostsTableSQL)
	if err != nil {
		log.Fatal(err.Error())
	}

	_, err = postsTable.Exec()
	if err != nil {
		log.Fatal(err.Error())
	}

	log.Println("Post table created in DB")
}

func insertPostIntoDB(db *sql.DB, postData models.Post) {
	insertPost, err := db.Prepare(models.InsertPostSQL)
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

	var post models.Post
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

func GetAllPosts(db *sql.DB) []models.Post {
	var posts []models.Post

	row, err := db.Query(models.GetAllPostsSQL)
	if err != nil {
		log.Fatal(err)
	}
	defer row.Close()

	var author models.User

	for row.Next() {
		var post = models.Post{
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

	posts := GetAllPosts(db)
	json.NewEncoder(w).Encode(posts)

	fmt.Println("All posts sent to clientside")
}