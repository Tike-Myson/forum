package main

import (
	"github.com/Tike-Myson/forum/cmd/database"
	"net/http"
)

func (app *application) home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	switch r.Method {
	case "GET":
		//
	case "POST":
		//
	default:
		app.errorStatus(w, http.StatusMethodNotAllowed)
	}
}

func (app *application) PostsAPI(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	//if r.URL.Path != "/api/posts" {
	//	http.NotFound(w, r)
	//	return
	//}
	//
	//switch r.Method {
	//case "GET":
		database.SendAllPosts(db,w,r)
	//default:
	//	app.errorStatus(w, http.StatusMethodNotAllowed)
	//}
}

func (app *application) CreatePostAPI(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	//if r.URL.Path != "/api/create-post/" {
	//	http.NotFound(w, r)
	//	return
	//}

	//switch r.Method {
	//case "POST":
		database.InsertPost(db,w,r)
	//default:
	//	app.errorStatus(w, http.StatusMethodNotAllowed)
	//}
}

func (app *application) SignUp(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/api/auth/signup" {
		http.NotFound(w, r)
		return
	}

	switch r.Method {
	case "POST":
		database.InsertUser(db, w, r)
	default:
		app.errorStatus(w, http.StatusMethodNotAllowed)
	}
}

func (app *application) Login(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/api/auth/login" {
		http.NotFound(w, r)
		return
	}

	switch r.Method {
	case "GET":
		//
	case "POST":
		//
	default:
		app.errorStatus(w, http.StatusMethodNotAllowed)
	}
}

func (app *application) errorStatus(w http.ResponseWriter, status int) {
	w.WriteHeader(status)
}