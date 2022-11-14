package main

import (
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
	w.Write([]byte("Hello World"))
}
