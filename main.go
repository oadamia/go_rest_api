package main

import (
	"log"
	"net/http"
)

func main() {
	log.Println("server started")
	log.Fatal(http.ListenAndServe(":3000", nil))
}
