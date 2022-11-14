package main

import (
	"log"
	"net/http"
)

func main() {
	log.Println("server started")
	registerHandler()
	log.Fatal(http.ListenAndServe(":3000", nil))
}
