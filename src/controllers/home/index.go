package home

import (
	"encoding/json"
	"net/http"
)

func Index(w http.ResponseWriter, r *http.Request) {
	info := map[string]string{}
	info["verb"] = "get"
	enc := json.NewEncoder(w)
	enc.Encode(&info)
}
