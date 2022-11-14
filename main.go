package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func main() {
	log.Println("server started")
	log.Fatal(http.ListenAndServe(":3000", handler{}))
}

type handler struct {
}

func (h handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)

	m := movie{
		ID:    1,
		Title: "scarface",
		Desc:  "Drama",
	}

	json.NewEncoder(w).Encode(m)

}

type movie struct {
	ID    int
	Title string
	Desc  string
}
