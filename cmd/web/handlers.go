package main

import "net/http"

func home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	switch r.Method {
	case "GET":
		w.Write([]byte("Hello World!"))
	case "POST":
		//
	default:
		errorStatus(w, http.StatusMethodNotAllowed)
	}
}

func login(w http.ResponseWriter, r *http.Request) {
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
		errorStatus(w, http.StatusMethodNotAllowed)
	}
}

func signup(w http.ResponseWriter, r *http.Request) {
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
		errorStatus(w, http.StatusMethodNotAllowed)
	}
}

func logout(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/logout" {
		http.NotFound(w, r)
		return
	}
	switch r.Method {
	case "POST":
		//
	default:
		errorStatus(w, http.StatusMethodNotAllowed)
	}
}

func comment(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/comment" {
		http.NotFound(w, r)
		return
	}

	switch r.Method {
	case "POST":
		//
	default:
		errorStatus(w, http.StatusMethodNotAllowed)
	}
}

func errorStatus(w http.ResponseWriter, status int) {
	w.WriteHeader(status)
}

func like(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/like" {
		http.NotFound(w, r)
		return
	}

	switch r.Method {
	case "POST":
		//
	default:
		errorStatus(w, http.StatusMethodNotAllowed)
	}
}

func dislike(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/dislike" {
		http.NotFound(w, r)
		return
	}

	switch r.Method {
	case "POST":
		//
	default:
		errorStatus(w, http.StatusMethodNotAllowed)
	}
}

func profile(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/profile" {
		http.NotFound(w, r)
		return
	}

	switch r.Method {
	case "POST":
		//
	default:
		errorStatus(w, http.StatusMethodNotAllowed)
	}
}

func likedPosts(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/likedPosts" {
		http.NotFound(w, r)
		return
	}

	switch r.Method {
	case "POST":
		//
	default:
		errorStatus(w, http.StatusMethodNotAllowed)
	}
}

func filterByCategory(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/category" {
		http.NotFound(w, r)
		return
	}

	switch r.Method {
	case "POST":
		//
	default:
		errorStatus(w, http.StatusMethodNotAllowed)
	}
}