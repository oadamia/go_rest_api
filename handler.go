package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func registerHandler() {
	http.Handle("/restapi/v1/movies", handler{})
}

type handler struct {
}

func (h handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log.Println("serve http")
	w.Header().Set("Content-Type", "application/json")

	switch r.Method {
	case http.MethodGet:
		w.WriteHeader(http.StatusOK)
		m := movie{
			ID:    1,
			Title: "scarface",
			Desc:  "Drama",
		}
		json.NewEncoder(w).Encode(m)
	case http.MethodPost:
		w.WriteHeader(http.StatusCreated)
		var m movie
		json.NewDecoder(r.Body).Decode(&m)
		log.Println(m)
	default:
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(`{ "message": "not found" }`))
	}

}
