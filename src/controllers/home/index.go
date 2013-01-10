package home

import (
	"fmt"
	"html"
	"net/http"
)

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, html.EscapeString("GET"))
}

func Create(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, html.EscapeString("POST"))
}

func Update(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, html.EscapeString("PUT"))
}

func Delete(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, html.EscapeString("DELETE"))
}
