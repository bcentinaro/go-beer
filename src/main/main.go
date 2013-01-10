package main

import (
	"controllers/home"
	"github.com/gorilla/mux"
	"net/http"
)

func main() {
	r := mux.NewRouter()
	get := r.Methods("get").Subrouter()
	post := r.Methods("post").Subrouter()
	put := r.Methods("put").Subrouter()
	delete := r.Methods("delete").Subrouter()

	get.HandleFunc("/", home.Index)
	post.HandleFunc("/", home.Create)
	put.HandleFunc("/", home.Update)
	delete.HandleFunc("/", home.Delete)

	http.Handle("/", r)
	http.ListenAndServe(":8080", nil)
}
