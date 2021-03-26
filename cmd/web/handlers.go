package main

import (
	"html/template"
	"log"
	"net/http"
)

func (app *application) home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	switch r.Method {
	case "GET":
		files := []string {
			"./ui/html/homePage.html",
			"./ui/html/baseLayout.html",
			"./ui/html/footerPartial.html",
		}
		ts, err := template.ParseFiles(files...)
		if err != nil {
			log.Println(err.Error())
			http.Error(w, "Internal Server Error", 500)
			return
		}
		err = ts.Execute(w, nil)
		if err != nil {
			log.Println(err.Error())
			http.Error(w, "Internal Server Error", 500)
		}
	case "POST":
		//
	default:
		app.errorStatus(w, http.StatusMethodNotAllowed)
	}
}

func (app *application) login(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/login" {
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

func (app *application) signup(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/signup" {
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

func (app *application) logout(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/logout" {
		http.NotFound(w, r)
		return
	}
	switch r.Method {
	case "POST":
		//
	default:
		app.errorStatus(w, http.StatusMethodNotAllowed)
	}
}

func (app *application) comment(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/comment" {
		http.NotFound(w, r)
		return
	}

	switch r.Method {
	case "POST":
		//
	default:
		app.errorStatus(w, http.StatusMethodNotAllowed)
	}
}

func (app *application) errorStatus(w http.ResponseWriter, status int) {
	w.WriteHeader(status)
}

func (app *application) like(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/like" {
		http.NotFound(w, r)
		return
	}

	switch r.Method {
	case "POST":
		//
	default:
		app.errorStatus(w, http.StatusMethodNotAllowed)
	}
}

func (app *application) dislike(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/dislike" {
		http.NotFound(w, r)
		return
	}

	switch r.Method {
	case "POST":
		//
	default:
		app.errorStatus(w, http.StatusMethodNotAllowed)
	}
}

func (app *application) profile(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/profile" {
		http.NotFound(w, r)
		return
	}

	switch r.Method {
	case "POST":
		//
	default:
		app.errorStatus(w, http.StatusMethodNotAllowed)
	}
}

func (app *application) likedPosts(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/likedPosts" {
		http.NotFound(w, r)
		return
	}

	switch r.Method {
	case "POST":
		//
	default:
		app.errorStatus(w, http.StatusMethodNotAllowed)
	}
}

func (app *application) filterByCategory(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/category" {
		http.NotFound(w, r)
		return
	}

	switch r.Method {
	case "POST":
		//
	default:
		app.errorStatus(w, http.StatusMethodNotAllowed)
	}
}