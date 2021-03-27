package main

import (
	"net/http"
)

func (app *application) commonHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./ui/dist/index.html")
}

func (app *application) routes() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/profile", app.commonHandler)
	mux.HandleFunc("/about", app.commonHandler)
	mux.HandleFunc("/login", app.commonHandler)
	mux.HandleFunc("/signup", app.commonHandler)

	fileServer := http.FileServer(http.Dir("./ui/dist/"))
	mux.Handle("/", http.StripPrefix("/", fileServer))

	mux.HandleFunc("/api/posts", app.PostsAPI)
	mux.HandleFunc("/api/create-post", app.CreatePostAPI)
	mux.HandleFunc("/api/auth/signup", app.SignUp)

	return mux
}