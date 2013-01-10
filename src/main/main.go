package main

import (
	"controllers/home"
	"github.com/gorilla/mux"
	"net/http"
)

func main() {
	// So we've go ourselves an express style router.
	r := mux.NewRouter()
	get := r.Methods("get").Subrouter()
	post := r.Methods("post").Subrouter()
	put := r.Methods("put").Subrouter()
	delete := r.Methods("delete").Subrouter()

	get.HandleFunc("/", home.Index)
	post.HandleFunc("/", home.Index)
	put.HandleFunc("/", home.Index)
	delete.HandleFunc("/", home.Index)

	http.Handle("/", r)
	http.ListenAndServe(":8080", nil)
}
