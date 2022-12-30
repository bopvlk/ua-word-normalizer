package scraper

import (
	"encoding/json"

	"net/http"
)

func fetch(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	var badStr []string

	if r.Method == "POST" {
		w.WriteHeader(200)
		json.NewDecoder(r.Body).Decode(&badStr)
		MakingGoodWords(badStr)
		json.NewEncoder(w).Encode(badStr)
	}

}

func Server() {
	http.HandleFunc("/", fetch)
	http.ListenAndServe(":1207", nil)
}
