package home

import (
	"encoding/json"
	"fmt"
	"html"
	"net/http"
)

func Index(w http.ResponseWriter, r *http.Request) {
	info := map[string]string{}
	info["verb"] = "get"
	doc, err := json.Marshal(info)
	if err == nil {
	}
	fmt.Fprintf(w, string(doc))
}
