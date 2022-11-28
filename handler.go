package main

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"strconv"
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
		get(w, r)
	case http.MethodPost:
		post(w, r)
	case http.MethodPut:
		put(w, r)
	case http.MethodDelete:
		remove(w, r)
	default:
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(`{ "message": "not found" }`))
	}

}

func get(w http.ResponseWriter, r *http.Request) {
	id, err := parse(r)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	value := movieDB[id]
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(value)
}

func post(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusCreated)
	var m movie
	json.NewDecoder(r.Body).Decode(&m)
	movieDB[m.ID] = m
}

func put(w http.ResponseWriter, r *http.Request) {
	id, err := parse(r)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	var m movie
	json.NewDecoder(r.Body).Decode(&m)
	movieDB[id] = m
}

func remove(w http.ResponseWriter, r *http.Request) {
	id, err := parse(r)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}
	delete(movieDB, id)
}

func parse(r *http.Request) (int, error) {
	p, ok := r.URL.Query()["id"]
	if !ok {
		return 0, errors.New("bad request")
	}

	id, err := strconv.Atoi(p[0])
	if err != nil {
		return 0, err
	}

	return id, nil

}
