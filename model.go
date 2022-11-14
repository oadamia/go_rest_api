package main

var movieDB = make(map[int]movie)

type movie struct {
	ID    int
	Title string
	Desc  string
}
