package main

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"

	aru "github.com/superkrasotkaaru/go-dev-tsis1/pkg"
)

type Response struct {
	Wishlist []aru.Wishlist `json:"wishlist"`
}

func respondWithError(w http.ResponseWriter, code int, message string) {
	respondWithJSON(w, code, map[string]string{"error": message})
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, err := json.Marshal(payload)

	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "500 Internal Server Error")
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

func healthCheck(w http.ResponseWriter, r *http.Request) {
	respondWithJSON(w, http.StatusOK, map[string]string{"status": "ok", "info": aru.Info()})
}

func wishlist(w http.ResponseWriter, r *http.Request) {
	wishlist := aru.GetWishlist()
	respondWithJSON(w, http.StatusOK, wishlist)
}

func wish(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	wish, err := aru.GetWishlist(id)
	if err != nil {
		respondWithError(w, http.StatusNotFound, "404 Not Found")
		return
	}

	respondWithJSON(w, http.StatusOK, wish)
}